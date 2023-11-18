package repository

import (
	"context"
	"github.com/jasonzhao47/goTrainingCamp/hw5/webook/internal/domain"
	"github.com/jasonzhao47/goTrainingCamp/hw5/webook/internal/repository/dao"
)

//go:generate mockgen -source=./async_sms_repository.go -package=repomocks -destination=mocks/async_sms_repository.mock.go AsyncSmsRepository
type AsyncSmsRepository interface {
	Add(ctx context.Context, s domain.AsyncSms) error
	FindFailedRequest(ctx context.Context) ([]domain.AsyncSms, error)
}

type asyncSmsRepository struct {
	dao dao.AsyncSmsDAO
}

func NewAsyncSMSRepository(dao dao.AsyncSmsDAO) AsyncSmsRepository {
	return &asyncSmsRepository{
		dao: dao,
	}
}

func (a *asyncSmsRepository) Add(ctx context.Context, s domain.AsyncSms) error {
	return a.dao.Insert(ctx, dao.AsyncSms{
		Id:       s.Id,
		TplId:    s.TplId,
		Args:     s.Args,
		Numbers:  s.Numbers,
		RetryMax: s.RetryMax,
	})
}

func (a *asyncSmsRepository) FindFailedRequest(ctx context.Context) ([]domain.AsyncSms, error) {
	return a.dao.GetFailed(ctx)
}
