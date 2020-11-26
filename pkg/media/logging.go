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

func (s loggingService) ReadMedia(id string, rate int) (_ []byte, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"domain", "media",
			"method", "ReadMedia",
			"id", id,
			"rate", rate,
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())

	return s.Service.ReadMedia(id, rate)
}

func (s loggingService) FindCoverArt(id string) (_ []byte, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"domain", "media",
			"method", "FindCoverArt",
			"id", id,
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())

	return s.Service.FindCoverArt(id)
}
