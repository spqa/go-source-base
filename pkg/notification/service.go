package notification

import (
	"bytes"
	_ "embed"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
	"go.uber.org/zap"
	"html/template"
	"mcm-api/config"
	"mcm-api/pkg/log"
)

type EmailTemplate string

const (
	NewContributionTemplate EmailTemplate = "new_contribution"
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

func (s Service) sendEmail(tmpl EmailTemplate, payload interface{}) error {
	body, subject, err := generateBodyAndSubject(tmpl, payload)
	if err != nil {
		return err
	}
	output, err := s.ses.SendEmail(&ses.SendEmailInput{
		Destination: &ses.Destination{
			ToAddresses: nil,
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Html: &ses.Content{
					Charset: aws.String("UTF-8"),
					Data:    aws.String(body),
				},
				Text: &ses.Content{
					Charset: aws.String("UTF-8"),
					Data:    aws.String(body),
				},
			},
			Subject: &ses.Content{
				Charset: aws.String("UTF-8"),
				Data:    aws.String(subject),
			},
		},
		Source: aws.String(s.cfg.SesSenderEmail),
	})
	if err != nil {
		return err
	}
	log.Logger.Debug("send email result", zap.Any("output", output))
	return nil
}

var parsedTemplate *template.Template

//go:embed templates/new_contribution.tmpl
var newContributionTemplate string

func init() {
	parsedTemplate = template.Must(template.New(string(NewContributionTemplate)).Parse(newContributionTemplate))
}

func generateBodyAndSubject(tmpl EmailTemplate, payload interface{}) (string, string, error) {
	switch tmpl {
	case NewContributionTemplate:
		if v, ok := payload.(TemplateNewContributionPayLoad); ok {
			buf := new(bytes.Buffer)
			err := parsedTemplate.ExecuteTemplate(buf, string(NewContributionTemplate), v)
			if err != nil {
				return "", "", err
			}
			return buf.String(), fmt.Sprintf("New contribution from %s", v.StudentName), nil
		}
		return "", "", errors.New("wrong type of payload")
	default:
		return "", "", fmt.Errorf("unknown template %v", tmpl)
	}
}

func (s Service) SendNewContributionEmail(payload *TemplateNewContributionPayLoad) error {
	return s.sendEmail(NewContributionTemplate, payload)
}
