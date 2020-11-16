package browsing

import (
	"io/ioutil"
	"os"
	"strings"

	"gitlab.com/jeremybush/gosonic/pkg/subsonic"
)

type Service interface {
	GetMusicFolders() (subsonic.MusicFolders, error)
	GetIndexes() (subsonic.Indexes, error)
}

type service struct{}

func NewService() *service {
	return &service{}
}

func (s service) GetMusicFolders() (subsonic.MusicFolders, error) {
	return subsonic.MusicFolders{
		MusicFolders: []subsonic.MusicFolder{
			{ID: "1", Name: "Music"},
		},
	}, nil
}

func (s service) GetIndexes() (subsonic.Indexes, error) {
	var (
		files []os.FileInfo

		artists map[string][]subsonic.Artist

		err error
	)

	artists = make(map[string][]subsonic.Artist)

	files, err = ioutil.ReadDir("/mnt/media/music")
	if err != nil {
		return subsonic.Indexes{}, err
	}

	for _, f := range files {
		if f.IsDir() {
			letter := string(strings.ToUpper(f.Name())[0])

			if _, ok := artists[letter]; ok {
				artists[letter] = append(artists[letter], subsonic.Artist{
					ID: "", Name: f.Name(),
				})
			} else {
				artists[letter] = []subsonic.Artist{{
					ID: "", Name: f.Name(),
				}}
			}
		}
	}

	return subsonic.Indexes{
		Indexes: func() []subsonic.Index {
			var indexes []subsonic.Index

			for name, artists := range artists {
				indexes = append(indexes, subsonic.Index{
					Name:    name,
					Artists: artists,
				})
			}

			return indexes
		}(),
	}, nil
}
