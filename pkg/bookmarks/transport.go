package bookmarks

import (
	"context"
	"net/http"

	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"

	"gitlab.com/jeremybush/gosonic/pkg/subsonic"
)

// All this is faked for now
func GetPlayQueue(
	encodeResponse func(ctx context.Context, w http.ResponseWriter, response interface{}) error,
	opts []kithttp.ServerOption,
) http.Handler {
	getPlayQueue := kithttp.NewServer(
		func(ctx context.Context, request interface{}) (interface{}, error) {
			return subsonic.GetPlayQueue{
				SubsonicResponse: subsonic.SubsonicResponse{
					Status: "ok", Version: "1.16.1",
				},
				PlayQueue: subsonic.PlayQueue{
					Entries: []subsonic.Song{
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

	r.Handle("/rest/getPlayQueue.view", getPlayQueue).Methods("GET")

	return r
}

func GetBookmarks(
	encodeResponse func(ctx context.Context, w http.ResponseWriter, response interface{}) error,
	opts []kithttp.ServerOption,
) http.Handler {
	getBookmarks := kithttp.NewServer(
		func(ctx context.Context, request interface{}) (interface{}, error) {
			return getBookmarksResponse{
				SubsonicResponse: subsonic.SubsonicResponse{
					Status: "ok", Version: "1.16.1",
				},
				Bookmarks: subsonic.Bookmarks{
					Bookmarks: []subsonic.Bookmark{
						{
							Entries: []subsonic.Song{
								{
									ID:     "11",
									Parent: "1",
									Title:  "Test Album",
									Artist: "Test Artist",
								},
							},
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

	r.Handle("/rest/getBookmarks.view", getBookmarks).Methods("GET")

	return r
}
