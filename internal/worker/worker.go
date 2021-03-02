package worker

import (
	"context"
	"errors"
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

func newWorker(config *config.Config, queue queue.Queue) *worker {
	return &worker{
		cfg:   config,
		queue: queue,
	}
}

func (w worker) Start() {
	log.Logger.Info("starting worker")
	ctx, cancelFunc := context.WithCancel(context.Background())
	go func() {
		signalChannel := make(chan os.Signal, 1)
		signal.Notify(signalChannel, os.Interrupt)
		s := <-signalChannel
		log.Logger.Info("receive signal", zap.String("signal", s.String()))
		cancelFunc()
		log.Logger.Info("grateful shutdown...")
	}()
poolQueueLoop:
	for {
		select {
		case <-ctx.Done():
			break poolQueueLoop
		default:
			message, err := w.queue.Pop(ctx)
			if err != nil {
				if !errors.Is(err, context.Canceled) {
					log.Logger.Error("pop queue error", zap.Error(err))
				}
				break poolQueueLoop
			}

			if message == nil {
				log.Logger.Debug("receive empty message")
				continue
			}
			log.Logger.Info("receive message", zap.Any("message", message))
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
