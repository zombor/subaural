package browsing

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"gitlab.com/jeremybush/gosonic/pkg/subsonic"
)

type Service interface {
	GetMusicFolders() (subsonic.MusicFolders, error)
	GetIndexes() (subsonic.Indexes, error)
	GetMusicDirectory(id string) (subsonic.MusicDirectory, error)
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
					ID: subsonic.PathID(f.Name()), Name: f.Name(),
				})
			} else {
				artists[letter] = []subsonic.Artist{{
					ID: subsonic.PathID(f.Name()), Name: f.Name(),
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

func (s service) GetMusicDirectory(id string) (subsonic.MusicDirectory, error) {
	var (
		files []os.FileInfo

		parent, child string
		children      []subsonic.DirectoryChild

		err error
	)

	parent, child, err = subsonic.ParentID(id)
	if err != nil {
		return subsonic.MusicDirectory{}, err
	}

	files, err = subsonic.ReadDir(id)
	if err != nil {
		return subsonic.MusicDirectory{}, err
	}

	for _, f := range files {
		if f.IsDir() {
			children = append(children, subsonic.DirectoryChild{
				ID:       subsonic.PathID(fmt.Sprintf("%s/%s", parent, f.Name())),
				CoverArt: subsonic.PathID(fmt.Sprintf("%s/%s", parent, f.Name())),
				Parent:   subsonic.PathID(parent),
				Title:    f.Name(),
				IsDir:    true,
			})
		} else if ok, meta, err := subsonic.ParseFlac(id, f.Name()); err == nil && ok {
			children = append(children, subsonic.DirectoryChild{
				ID:       subsonic.PathID(fmt.Sprintf("%s/%s/%s", parent, child, f.Name())),
				CoverArt: subsonic.PathID(fmt.Sprintf("%s/%s/%s", parent, child, f.Name())),
				Parent:   subsonic.PathID(parent),
				Title:    meta.Title,
				IsDir:    false,

				Album: meta.Album,
				Track: meta.Track,
				Year:  meta.Date,
				Genre: meta.Genre,
			})
		}
	}

	return subsonic.MusicDirectory{
		ID:     id,
		Name:   child,
		Parent: subsonic.PathID(parent),

		Children: children,
	}, nil
}
