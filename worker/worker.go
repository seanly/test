package main

import (
	"context"
	"log"

	"github.com/hibiken/asynq"
	"golang.org/x/sync/errgroup"
)

var (
	g            errgroup.Group
	redisConnOpt = asynq.RedisClientOpt{Addr: "redis:6379"}
)

func handler(ctx context.Context, t *asynq.Task) error {
	return nil
}

func handlePreEnqueue(task *asynq.Task, opts []asynq.Option) {

}

func handlePostEnqueue(task *asynq.TaskInfo, err error) {

}

func runScheduler() error {
	scheduler := asynq.NewScheduler(
		redisConnOpt,
		&asynq.SchedulerOpts{
			PreEnqueueFunc:  handlePreEnqueue,
			PostEnqueueFunc: handlePostEnqueue,
		},
	)

	return scheduler.Run()
}

func runSrv() error {

	srv := asynq.NewServer(
		redisConnOpt,
		asynq.Config{Concurrency: 10},
	)
	return srv.Run(asynq.HandlerFunc(handler))
}

func main() {

	g.Go(func() error {
		return runSrv()
	})

	g.Go(func() error {
		return runScheduler()
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}
