package main

import (
	"log"
	"time"

	"github.com/RishabhKatiyar/AsynqExample/tasks"
	"github.com/hibiken/asynq"
)

func main() {
	redisConnection := asynq.RedisClientOpt{
		Addr: "localhost:6379",
	}

	loc, err := time.LoadLocation("UTC")
	if err != nil {
		panic(err)
	}

	scheduler := asynq.NewScheduler(
		redisConnection,
		&asynq.SchedulerOpts{
			Location: loc,
		},
	)

	welcomeEmailTask := tasks.NewWelcomeEmailTask(1)

	//entryID, err := scheduler.Register("@every 1s", welcomeEmailTask)
	entryID, err := scheduler.Register("@every 1s", welcomeEmailTask, asynq.Queue("critical"))
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("registered an entry: %q\n", entryID)

	// Run scheduler
	if err := scheduler.Run(); err != nil {
		log.Fatal(err)
	}
}
