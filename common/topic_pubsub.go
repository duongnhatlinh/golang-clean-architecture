package common

import "food_delivery/pubsub"

const (
	TopicUserLikeRestaurant    pubsub.Topic = "TopicUserLikeRestaurant"
	TopicUserDislikeRestaurant pubsub.Topic = "TopicUserDislikeRestaurant"
	TopicUserLikeFood          pubsub.Topic = "TopicUserLikeFood"
	TopicUserDislikeFood       pubsub.Topic = "TopicUserDislikeFood"
)
