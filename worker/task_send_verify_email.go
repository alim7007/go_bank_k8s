package worker

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/alim7007/go_bank_k8s/mail"
	"github.com/alim7007/go_bank_k8s/util"
	"github.com/hibiken/asynq"
	"github.com/rs/zerolog/log"
)

const TaskSendVerifyEmail = "task:send_verify_email"

type PayloadSendVerifyEmail struct {
	Username string `json:"username"`
}

func (distributor *RedisTaskDistributor) DistributeTaskSendVerifyEmail(
	ctx context.Context,
	payload *PayloadSendVerifyEmail,
	opts ...asynq.Option,
) error {
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal task payload %w", err)
	}
	task := asynq.NewTask(TaskSendVerifyEmail, jsonPayload, opts...)
	info, err := distributor.client.EnqueueContext(ctx, task)
	if err != nil {
		return fmt.Errorf("failed to enqueue task: %w", err)
	}
	log.Info().Str("type", task.Type()).Bytes("payload", task.Payload()).Str("queue", info.Queue).Int("max_retry", info.MaxRetry).Msg("enqueued task")
	return nil
}

func (processor *RedisTaskProcessor) ProcesssTaskSencVerifyEmail(ctx context.Context, task *asynq.Task) error {
	var payload PayloadSendVerifyEmail
	if err := json.Unmarshal(task.Payload(), &payload); err != nil {
		return fmt.Errorf("failed to unmarshal payload %w", asynq.SkipRetry)
	}

	config, err := util.LoadConfig("..")

	if err != nil {
		return err
	}
	sender := mail.NewGmailSender(config.EmailSenderName, config.EmailSenderAddress, config.EmailSenderPassword)

	subject := "A test email"
	content := `
	<h1>Hello world</h1>
	<p>This is a test message from <a href="http://techschool.guru">Olim Bank</a></p>
	`
	to := []string{"toni7007olim@gmail.com"}
	attachFiles := []string{"text.txt"}

	err = sender.SendEmail(subject, content, to, nil, nil, attachFiles)
	if err != nil {
		return err
	}

	user, err := processor.store.GetUser(ctx, payload.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("user doesn't exist %w", asynq.SkipRetry)
		}
		return fmt.Errorf("failed to get user %w", err)
	}
	log.Info().Str("type", task.Type()).Bytes("payload", task.Payload()).Str("email", user.Email).Msg("processed task")
	return nil
}
