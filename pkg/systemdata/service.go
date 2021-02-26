package systemdata

import (
	"context"
	"github.com/casbin/casbin/v2"
	"mcm-api/config"
	"mcm-api/pkg/common"
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

func (s Service) Find(ctx context.Context) (*common.PaginateResponse, error) {
	return nil, nil
}

func (s Service) Update(ctx context.Context, id string, body *DataUpdateReq) error {
	return nil
}
