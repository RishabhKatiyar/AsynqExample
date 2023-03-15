package main

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis"
	"github.com/hibiken/asynq"
)

// HandleWelcomeEmailTask handler for welcome email task.
func HandleWelcomeEmailTask(c context.Context, t *asynq.Task) error {
	// Get user ID from given task.
	//id, err := t.Payload.GetInt("user_id")
	var ds map[string]interface{}
	json.Unmarshal(t.Payload(), &ds)
	//fmt.Println("In Handler..")
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	defer client.Close()
	client.Incr("checkVal")
	time.Sleep(time.Second * 2)
	return nil
}
