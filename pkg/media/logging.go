package media

import (
	"time"

	"github.com/go-kit/kit/log"
)

type loggingService struct {
	logger log.Logger
	Service
}

func NewLoggingService(logger log.Logger, s Service) Service {
	return &loggingService{logger, s}
}

func (s loggingService) ReadMedia(id string) (_ []byte, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "ReadMedia",
			"id", id,
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())

	return s.Service.ReadMedia(id)
}

func (s loggingService) FindCoverArt(id string) (_ []byte, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "FindCoverArt",
			"id", id,
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())

	return s.Service.FindCoverArt(id)
}
