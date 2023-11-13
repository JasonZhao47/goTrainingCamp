package dao

import (
	"context"
	"github.com/jasonzhao47/goTrainingCamp/hw5/webook/internal/domain"
	"gorm.io/gorm"
)

//go:generate mockgen -source=./async_sms.go -package=daomocks -destination=mocks/async_sms.mock.go AsyncSmsDAO
type AsyncSmsDAO interface {
	Insert(ctx context.Context, s AsyncSms) error
	GetFailed(ctx context.Context) ([]domain.AsyncSms, error)
}

const (
	FailedTask = iota + 1
	SuccessTask
)

type GORMAsyncSmsDAO struct {
	db *gorm.DB
}

func NewGORMAsyncSmsDAO(db *gorm.DB) AsyncSmsDAO {
	return &GORMAsyncSmsDAO{
		db: db,
	}
}

func (g *GORMAsyncSmsDAO) Insert(ctx context.Context, s AsyncSms) error {
	return g.db.Create(&s).Error
}

func (g *GORMAsyncSmsDAO) GetFailed(ctx context.Context) ([]domain.AsyncSms, error) {
	var sms []AsyncSms
	var res []domain.AsyncSms
	err := g.db.WithContext(ctx).Model(&AsyncSms{}).Where("status = ?", FailedTask).Take(&sms).Error
	if err != nil {
		return []domain.AsyncSms{}, err
	}
	for i := range sms {
		res = append(res, sms[i].toDomain())
	}
	return res, err
}

func (sms AsyncSms) toDomain() domain.AsyncSms {
	return domain.AsyncSms{
		Id:       sms.Id,
		TplId:    sms.TplId,
		Args:     sms.Args,
		Numbers:  sms.Numbers,
		RetryMax: sms.RetryMax,
	}
}

type AsyncSms struct {
	Id       int64
	TplId    string
	Args     []string
	Numbers  []string
	RetryMax int
}
