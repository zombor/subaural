package subsonic

import (
	"os"
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

func (s loggingService) ParseFlac(parent, fileName string) (_ bool, _ FlacMeta, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"domain", "subsonic",
			"method", "ParseFlac",
			"parent", parent,
			"fileName", fileName,
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())

	return s.Service.ParseFlac(parent, fileName)
}

func (s loggingService) ReadFile(path string, rate int) (_ []byte, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"domain", "subsonic",
			"method", "ReadFile",
			"path", path,
			"rate", rate,
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())

	return s.Service.ReadFile(path, rate)
}

func (s loggingService) FindCoverArt(path string) (_ []byte, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"domain", "subsonic",
			"method", "FindCoverArt",
			"path", path,
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())

	return s.Service.FindCoverArt(path)
}

func (s loggingService) ReadDir(dir string) (_ []os.FileInfo, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"domain", "subsonic",
			"method", "ReadDir",
			"dir", dir,
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())

	return s.Service.ReadDir(dir)
}
