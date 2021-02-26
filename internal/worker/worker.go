package worker

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"mcm-api/config"
	"mcm-api/pkg/log"
	"mcm-api/pkg/queue"
	"os"
	"os/signal"
)

type worker struct {
	cfg   *config.Config
	queue queue.Queue
}

func InitializeWorker(config *config.Config, queue queue.Queue) *worker {
	return &worker{
		cfg:   config,
		queue: queue,
	}
}

func (w worker) Start() {
	ctx, cancelFunc := context.WithCancel(context.Background())
	go func() {
		signalChannel := make(chan os.Signal, 1)
		signal.Notify(signalChannel, os.Interrupt)
		s := <-signalChannel
		log.Logger.Info("Receive signal", zap.String("signal", s.String()))
		cancelFunc()
		log.Logger.Info("Grateful shutdown...")
	}()
poolQueueLoop:
	for {
		select {
		case <-ctx.Done():
			break poolQueueLoop
		default:
			message, err := w.queue.Pop()
			if err != nil {
				log.Logger.Error("pop queue error", zap.Error(err))
				cancelFunc()
			}
			err = handleMessage(message)
			if err != nil {
				log.Logger.Error("process message error", zap.Error(err))
			}
		}
	}
}

func handleMessage(message *queue.Message) error {
	switch message.Topic {
	case queue.ContributionCreated:
		return contributionCreatedHandler(message)
	case queue.ArticleUploaded:
		return articleUploadedHandler(message)
	default:
		return fmt.Errorf("unknown topic %v", message.Topic)
	}
}

func contributionCreatedHandler(message *queue.Message) error {
	return nil
}

func articleUploadedHandler(message *queue.Message) error {
	return nil
}
