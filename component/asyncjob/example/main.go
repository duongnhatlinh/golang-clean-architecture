package main

import (
	"context"
	"fmt"
	"food_delivery/component/asyncjob"
)

func main() {
	j1 := func(ctx context.Context) error {
		fmt.Printf("I am testing1")
		return nil
	}
	job1 := asyncjob.NewJob(j1)

	j2 := func(ctx context.Context) error {
		fmt.Printf("I am testing2")
		return nil
	}
	job2 := asyncjob.NewJob(j2)

	g := asyncjob.NewGroup(true, job1, job2)

	g.Run(context.Background())

}
