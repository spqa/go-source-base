package notification

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
	"go.uber.org/zap"
	"mcm-api/config"
	"mcm-api/pkg/log"
)

type Service struct {
	cfg *config.Config
	ses *ses.SES
}

func InitializeService(cfg *config.Config) *Service {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2")},
	)
	if err != nil {
		log.Logger.Panic("failed to init notification service", zap.Error(err))
	}
	svc := ses.New(sess)
	return &Service{
		cfg: nil,
		ses: svc,
	}
}

func (s Service) BatchSendEmail() error {
	return nil
}
