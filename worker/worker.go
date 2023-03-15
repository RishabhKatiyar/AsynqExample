package main

import (
	"log"

	"github.com/RishabhKatiyar/AsynqExample/tasks"
	"github.com/hibiken/asynq"
)

func main() {
	redisConnection := asynq.RedisClientOpt{
		Addr: "localhost:6379",
	}

	worker := asynq.NewServer(redisConnection, asynq.Config{
		// Specify how many concurrent workers to use.
		Concurrency: 4,
		// Specify multiple queues with different priority.
		Queues: map[string]int{
			"critical": 6, // processed 60% of the time
			"default":  3, // processed 30% of the time
			"low":      1, // processed 10% of the time
		},
	})

	mux := asynq.NewServeMux()

	mux.HandleFunc(
		tasks.TypeWelcomeEmail,       // task type
		tasks.HandleWelcomeEmailTask, // handler function
	)

	if err := worker.Run(mux); err != nil {
		log.Fatal(err)
	}

}
