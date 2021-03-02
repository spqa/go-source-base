package queue

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"mcm-api/config"
	"mcm-api/pkg/log"
	"time"
)

type TopicType string

const (
	ContributionCreated TopicType = "contribution-created"
	ArticleUploaded     TopicType = "article-uploaded"
)

type Message struct {
	Topic TopicType   `json:"topic"`
	Data  interface{} `json:"data"`
}

type Queue interface {
	Add(ctx context.Context, message *Message) error
	Pop(ctx context.Context) (*Message, error)
}

type RedisQueue struct {
	cfg   *config.Config
	redis *redis.Client
}

func InitializeRedisQueue(cfg *config.Config, client *redis.Client) Queue {
	return &RedisQueue{
		cfg:   cfg,
		redis: client,
	}
}

func (r *RedisQueue) Add(ctx context.Context, message *Message) error {
	bytes, err := json.Marshal(message)
	if err != nil {
		return err
	}
	push := r.redis.LPush(ctx, r.cfg.RedisQueueName, bytes)
	if push.Err() != nil {
		return push.Err()
	}
	return nil
}

func (r *RedisQueue) Pop(ctx context.Context) (*Message, error) {
	p := r.redis.BLPop(ctx, time.Second*30, r.cfg.RedisQueueName)
	if errors.Is(p.Err(), redis.Nil) {
		return nil, nil
	}
	if p.Err() != nil {
		return nil, p.Err()
	}
	m := new(Message)
	err := json.Unmarshal([]byte(p.Val()[0]), m)
	if err != nil {
		log.Logger.Error("malformed message", zap.Error(err))
		return nil, nil
	}
	return m, nil
}
