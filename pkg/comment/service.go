package comment

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

func (s Service) Find(ctx context.Context, query *IndexQuery) (*common.PaginateResponse, error) {
	return nil, nil
}

func (s Service) FindById(ctx context.Context, id int) (*ContributionRes, error) {
	return nil, nil
}

func (s Service) Create(ctx context.Context, body *ContributionCreateReq) (*ContributionRes, error) {
	return nil, nil
}

func (s Service) Update(ctx context.Context, id int, body *ContributionUpdateReq) (*ContributionRes, error) {
	return nil, nil
}

func (s Service) Delete(ctx context.Context, id int) error {
	return nil
}
