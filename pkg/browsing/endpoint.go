package browsing

import (
	"context"

	"github.com/go-kit/kit/endpoint"

	"gitlab.com/jeremybush/gosonic/pkg/subsonic"
)

func makeGetMusicFoldersEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		//req := request.(bookCargoRequest)
		f, err := s.GetMusicFolders()
		return getMusicFoldersResponse{SubsonicResponse: subsonic.SubsonicHeaders(), MusicFolders: f, Err: err}, nil
	}
}

type getMusicFoldersResponse struct {
	subsonic.SubsonicResponse
	MusicFolders subsonic.MusicFolders `xml:"musicFolders"`
	Err          error
}

func (r getMusicFoldersResponse) error() error { return r.Err }

func makeGetIndexesEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		//req := request.(bookCargoRequest)
		i, err := s.GetIndexes()
		return getIndexesResponse{SubsonicResponse: subsonic.SubsonicHeaders(), Indexes: i, Err: err}, nil
	}
}

type getIndexesResponse struct {
	subsonic.SubsonicResponse
	Indexes subsonic.Indexes `xml:"indexes"`
	Err     error
}

func (r getIndexesResponse) error() error { return r.Err }

type getMusicDirectoryRequest struct {
	ID string
}

type getMusicDirectoryResponse struct {
	subsonic.SubsonicResponse
	Directory subsonic.MusicDirectory `xml:"directory"`
	Err       error
}

func makeGetMusicDirectoryEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getMusicDirectoryRequest)

		d, err := s.GetMusicDirectory(req.ID)
		if err != nil {
			println(err.Error())
		}

		return getMusicDirectoryResponse{SubsonicResponse: subsonic.SubsonicHeaders(), Directory: d, Err: err}, nil
	}
}
