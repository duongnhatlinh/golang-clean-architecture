package subscriber

import (
	"context"
	"food_delivery/common"
	"food_delivery/component/appctx"
	"food_delivery/component/asyncjob"
	"food_delivery/pubsub"
	"log"
)

type consumerJob struct {
	Title string
	Hld   func(ctx context.Context, message *pubsub.Message) error
}

type consumerEngine struct {
	appCtx appctx.AppContext
}

func NewEngine(appContext appctx.AppContext) *consumerEngine {
	return &consumerEngine{appCtx: appContext}
}

func (engine *consumerEngine) Start() error {
	//ps := engine.appCtx.GetPubsub()

	//engine.startSubTopic(common.ChanNoteCreated, asyncjob.NewGroup(
	//	false,
	//	asyncjob.NewJob(SendNotificationAfterCreateNote(engine.appCtx, context.Background(), nil))),
	//)
	//

	//engine.startSubTopic(
	//	common.TopicNoteCreated,
	//	true,
	//	DeleteImageRecordAfterCreateNote(engine.appCtx),
	//	SendEmailAfterCreateNote(engine.appCtx),
	//	EmitRealtimeAfterCreateNote(engine.appCtx, rtEngine),
	//)

	//engine.startSubTopic(
	//	common.TopicNoteCreated,
	//	false,
	//	DeleteImageRecordAfterCreateNote(engine.appCtx),
	//	SendEmailAfterCreateNote(engine.appCtx),
	//)
	// Many sub on a topic

	engine.startSubTopic(
		common.TopicUserLikeRestaurant,
		true,
		RunIncreaseLikeCountAfterUserLikeRestaurant(engine.appCtx),
	)

	engine.startSubTopic(
		common.TopicUserDislikeRestaurant,
		true,
		RunDecreaseLikeCountAfterUserDislikeRestaurant(engine.appCtx),
	)

	engine.startSubTopic(
		common.TopicUserLikeFood,
		true,
		RunIncreaseLikeCountAfterUserLikeFood(engine.appCtx),
	)

	engine.startSubTopic(
		common.TopicUserDislikeFood,
		true,
		RunDecreaseLikeCountAfterUserDislikeFood(engine.appCtx),
	)
	return nil
}

type GroupJob interface {
	Run(ctx context.Context) error
}

func (engine *consumerEngine) startSubTopic(topic pubsub.Topic, isConcurrent bool, consumerJobs ...consumerJob) error {
	c, _ := engine.appCtx.GetPubSub().Subscribe(context.Background(), topic)

	for _, item := range consumerJobs {
		log.Println("Setup consumer for:", item.Title)
	}

	getJobHandler := func(job *consumerJob, message *pubsub.Message) asyncjob.JobHandler {
		return func(ctx context.Context) error {
			log.Println("running job for ", job.Title, ". Value: ", message.Data())
			return job.Hld(ctx, message)
		}
	}

	go func() {
		for {
			msg := <-c

			jobHdlArr := make([]asyncjob.Job, len(consumerJobs))

			for i := range consumerJobs {
				jobHdl := getJobHandler(&consumerJobs[i], msg)
				jobHdlArr[i] = asyncjob.NewJob(jobHdl)
			}

			group := asyncjob.NewGroup(isConcurrent, jobHdlArr...)

			if err := group.Run(context.Background()); err != nil {
				log.Println(err)
			}
		}
	}()

	return nil
}
