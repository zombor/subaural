package media

type Service interface {
	ReadMedia(string, int) ([]byte, error)
	FindCoverArt(string) ([]byte, error)
}

type service struct {
	readFile     func(string, int) ([]byte, error)
	findCoverArt func(string) ([]byte, error)
}

func NewService(
	readFile func(string, int) ([]byte, error),
	findCoverArt func(string) ([]byte, error),
) *service {
	return &service{readFile, findCoverArt}
}

func (s service) ReadMedia(id string, rate int) ([]byte, error) {
	return s.readFile(id, rate)
}

func (s service) FindCoverArt(id string) ([]byte, error) {
	return s.findCoverArt(id)
}
