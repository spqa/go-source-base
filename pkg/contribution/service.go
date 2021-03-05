package contribution

import (
	"context"
	"mcm-api/config"
	"mcm-api/pkg/apperror"
	"mcm-api/pkg/common"
	"mcm-api/pkg/contributesession"
	"mcm-api/pkg/queue"
)

type Service struct {
	cfg                      *config.Config
	repository               *repository
	queue                    queue.Queue
	contributeSessionService *contributesession.Service
}

func InitializeService(
	cfg *config.Config,
	repository *repository,
	queue queue.Queue,
	cs *contributesession.Service,
) *Service {
	return &Service{
		queue:                    queue,
		cfg:                      cfg,
		repository:               repository,
		contributeSessionService: cs,
	}
}

func (s Service) Find(ctx context.Context, query *IndexQuery) (*common.PaginateResponse, error) {
	user, err := common.GetLoggedInUser(ctx)
	if err != nil {
		return nil, err
	}
	var result []*Entity
	var count int64
	switch user.Role {
	case common.MarketingManager:
		result, count, err = s.repository.FindAndCount(ctx, &IndexQuery{
			PaginateQuery:         query.PaginateQuery,
			FacultyId:             query.FacultyId,
			StudentId:             query.StudentId,
			ContributionSessionId: query.ContributionSessionId,
			Status:                Accepted,
		})
		if err != nil {
			return nil, err
		}
		break
	case common.MarketingCoordinator:
		result, count, err = s.repository.FindAndCount(ctx, &IndexQuery{
			PaginateQuery:         query.PaginateQuery,
			FacultyId:             user.FacultyId,
			StudentId:             query.StudentId,
			ContributionSessionId: query.ContributionSessionId,
			Status:                query.Status,
		})
		if err != nil {
			return nil, err
		}
		break
	case common.Student:
		result, count, err = s.repository.FindAndCount(ctx, &IndexQuery{
			PaginateQuery:         query.PaginateQuery,
			FacultyId:             user.FacultyId,
			StudentId:             &user.Id,
			ContributionSessionId: query.ContributionSessionId,
			Status:                query.Status,
		})
		if err != nil {
			return nil, err
		}
	default:
		return nil, apperror.New(apperror.ErrForbidden, "", nil)
	}

	return common.NewPaginateResponse(result, count, query.Page, query.GetLimit()), nil
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
