package media

import "gitlab.com/jeremybush/gosonic/pkg/subsonic"

type Service interface {
	ReadMedia(string, int) ([]byte, error)
	FindCoverArt(string) ([]byte, error)
}

type service struct{}

func NewService() *service {
	return &service{}
}

func (s service) ReadMedia(id string, rate int) ([]byte, error) {
	return subsonic.ReadFile(id, rate)
}

func (s service) FindCoverArt(id string) ([]byte, error) {
	return subsonic.FindCoverArt(id)
}
