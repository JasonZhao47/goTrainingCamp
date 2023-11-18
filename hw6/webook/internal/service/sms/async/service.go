package async

import (
	"context"
	"github.com/jasonzhao47/goTrainingCamp/hw5/webook/internal/domain"
	"github.com/jasonzhao47/goTrainingCamp/hw5/webook/internal/repository"
	"github.com/jasonzhao47/goTrainingCamp/hw5/webook/internal/service/sms"
	"sync"
	"time"
)

type SMSService struct {
	svc          sms.Service
	repo         repository.AsyncSmsRepository
	mu           sync.Mutex
	errorRecords []errorRecord
	// 判断崩溃的时间窗口（比如最近1分钟）
	timeWindow time.Duration
	// 错误率阈值
	errorRateThreshold float64
}

// errorRecord 定义了错误记录的结构
type errorRecord struct {
	timestamp time.Time
	success   bool
}

func NewSMSService() *SMSService {
	return &SMSService{}
}

func (s *SMSService) StartAsync(ctx context.Context) {
	go func() {
		reqs, err := s.repo.FindFailedRequest(ctx)
		if err != nil {
			panic(err)
		}
		for _, req := range reqs {
			s.svc.Send(ctx, req.TplId, req.Args, req.Numbers...)
		}
	}()
}

func (s *SMSService) Send(ctx context.Context, biz string, args []string, numbers ...string) error {
	err := s.svc.Send(ctx, biz, args, numbers...)
	if err != nil {
		if s.hasCrashed() {
			err := s.repo.Add(ctx, domain.AsyncSms{
				TplId:    biz,
				Args:     args,
				Numbers:  numbers,
				RetryMax: 4,
			})
			return err
		}
	}
	return nil
}

func (s *SMSService) hasCrashed() bool {
	// 判断第三方是否已经崩溃
	// 给定一个时间片，每次在这个时间片里面计算错误率
	s.mu.Lock()
	defer s.mu.Unlock()

	// 获取当前时间
	now := time.Now()
	// 定义错误计数器
	var errorCount, totalCount int

	// 遍历记录，计算时间窗口内的错误率
	for _, record := range s.errorRecords {
		if now.Sub(record.timestamp) <= s.timeWindow {
			totalCount++
			if !record.success {
				errorCount++
			}
		}
	}

	// 计算错误率
	var errorRate float64
	if totalCount > 0 {
		errorRate = float64(errorCount) / float64(totalCount)
	}
	// 判断是否超过阈值
	return errorRate >= s.errorRateThreshold
}
