package contribution

import (
	"context"
	"github.com/casbin/casbin/v2"
	"mcm-api/config"
	"mcm-api/pkg/apperror"
	"mcm-api/pkg/common"
	"mcm-api/pkg/contributesession"
	"mcm-api/pkg/queue"
)

type Service struct {
	cfg                      *config.Config
	repository               *repository
	enforcer                 *casbin.Enforcer
	queue                    queue.Queue
	contributeSessionService *contributesession.Service
}

func InitializeService(
	cfg *config.Config,
	repository *repository,
	enforcer *casbin.Enforcer,
	queue queue.Queue,
	cs *contributesession.Service,
) *Service {
	return &Service{
		queue:                    queue,
		cfg:                      cfg,
		repository:               repository,
		enforcer:                 enforcer,
		contributeSessionService: cs,
	}
}

func (s Service) Find(ctx context.Context, query *IndexQuery) (*common.PaginateResponse, error) {
	session, err := s.contributeSessionService.GetCurrentSession(ctx)
	if err != nil && !apperror.Is(err, apperror.ErrNotFound) {
		return nil, err
	}
	if session == nil {
		return common.NewEmptyPaginateResponse(), nil
	}
	// TODO
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
