package lists

import (
	"context"
	"net/http"

	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"

	"gitlab.com/jeremybush/gosonic/pkg/subsonic"
)

// Not sure how often this is really used?
func GetAlbumList(
	encodeResponse func(ctx context.Context, w http.ResponseWriter, response interface{}) error,
	opts []kithttp.ServerOption,
) http.Handler {
	getAlbumList := kithttp.NewServer(
		func(ctx context.Context, request interface{}) (interface{}, error) {
			return subsonic.GetAlbumList{
				SubsonicResponse: subsonic.SubsonicResponse{
					Status: "ok", Version: "1.16.1",
				},
				AlbumList: subsonic.AlbumLists{
					AlbumList: []subsonic.Album{
						{
							ID:     "11",
							Parent: "1",
							Title:  "Test Album",
							Artist: "Test Artist",
						},
					},
				},
			}, nil
		},
		func(_ context.Context, r *http.Request) (interface{}, error) { return nil, nil },
		encodeResponse,
		opts...,
	)

	r := mux.NewRouter()

	r.Handle("/rest/getAlbumList.view", getAlbumList).Methods("GET")

	return r
}

func GetRandomSongs(
	encodeResponse func(ctx context.Context, w http.ResponseWriter, response interface{}) error,
	opts []kithttp.ServerOption,
) http.Handler {
	getUserHandler := kithttp.NewServer(
		func(ctx context.Context, request interface{}) (interface{}, error) {
			return subsonic.GetRandomSongs{
				SubsonicResponse: subsonic.SubsonicResponse{
					Status: "ok", Version: "1.16.1",
				},
				RandomSongs: subsonic.RandomSongs{
					RandomSongs: []subsonic.Song{
						{
							ID:     "11",
							Parent: "1",
							Title:  "Test Album",
							Artist: "Test Artist",
						},
					},
				},
			}, nil
		},
		func(_ context.Context, r *http.Request) (interface{}, error) { return nil, nil },
		encodeResponse,
		opts...,
	)

	r := mux.NewRouter()

	r.Handle("/rest/getRandomSongs.view", getUserHandler).Methods("GET")

	return r
}
