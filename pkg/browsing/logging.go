package browsing

import (
	"time"

	"github.com/go-kit/kit/log"

	"gitlab.com/jeremybush/gosonic/pkg/subsonic"
)

type loggingService struct {
	logger log.Logger
	Service
}

func NewLoggingService(logger log.Logger, s Service) Service {
	return &loggingService{logger, s}
}

func (s loggingService) GetMusicFolders() (_ subsonic.MusicFolders, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "GetMusicFolders",
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())

	return s.Service.GetMusicFolders()
}

func (s loggingService) GetIndexes() (_ subsonic.Indexes, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "GetIndexes",
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())

	return s.Service.GetIndexes()
}

func (s loggingService) GetMusicDirectory(id string) (_ subsonic.MusicDirectory, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "GetMusicDirectory",
			"id", id,
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())

	return s.Service.GetMusicDirectory(id)
}
