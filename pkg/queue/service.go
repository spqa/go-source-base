package queue

import (
	"context"
	"github.com/go-redis/redis/v8"
	"mcm-api/config"
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
	Add(message *Message) error
	Pop() (*Message, error)
}

type RedisQueue struct {
	cfg         *config.Config
	redisClient *redis.Client
	ctx         context.Context
}

func InitializeRedisQueue(ctx context.Context, cfg *config.Config, client *redis.Client) Queue {
	return &RedisQueue{
		ctx:         ctx,
		cfg:         cfg,
		redisClient: client,
	}
}

func (r *RedisQueue) Add(message *Message) error {
	panic("implement me")
}

func (r *RedisQueue) Pop() (*Message, error) {
	panic("implement me")
}
