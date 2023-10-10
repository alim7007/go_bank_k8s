package worker

import (
	"context"

	db "github.com/alim7007/go_bank_k8s/db/sqlc"
	"github.com/hibiken/asynq"
)

type TaskProcessor interface {
	ProcesssTaskSencVerifyEmail(ctx context.Context, task *asynq.Task) error
	Start() error
}

type RedisTaskProcessor struct {
	server *asynq.Server
	store  db.Store
}

func NewRedisTaskProcessor(redisOpt asynq.RedisClientOpt, store db.Store) TaskProcessor {
	server := asynq.NewServer(
		redisOpt,
		asynq.Config{},
	)
	return &RedisTaskProcessor{
		server: server,
		store:  store,
	}
}

func (processor *RedisTaskProcessor) Start() error {
	mux := asynq.NewServeMux()
	mux.HandleFunc(TaskSendVerifyEmail, processor.ProcesssTaskSencVerifyEmail)
	return processor.server.Start(mux)
}