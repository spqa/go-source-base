package article

import (
	"context"
	"github.com/casbin/casbin/v2"
	"mcm-api/config"
)

type Service struct {
	cfg        *config.Config
	repository *repository
	enforcer   *casbin.Enforcer
}

func InitializeService(
	cfg *config.Config,
	repository *repository,
	enforcer *casbin.Enforcer,
) *Service {
	return &Service{
		cfg:        cfg,
		repository: repository,
		enforcer:   enforcer,
	}
}

func (s Service) FindById(ctx context.Context, id int) (*ArticleRes, error) {
	return nil, nil
}
