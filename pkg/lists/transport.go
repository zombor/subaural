package lists

import (
	"context"
	"encoding/json"
	"encoding/xml"
	"net/http"

	kitlog "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"

	"gitlab.com/jeremybush/gosonic/pkg/subsonic"
)

func GetAlbumList(logger kitlog.Logger) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		kithttp.ServerErrorEncoder(encodeError),
	}

	getUserHandler := kithttp.NewServer(
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

	r.Handle("/rest/getAlbumList.view", getUserHandler).Methods("GET")

	return r
}

func GetRandomSongs(logger kitlog.Logger) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		kithttp.ServerErrorEncoder(encodeError),
	}

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

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	switch err {
	/*
		case cargo.ErrUnknown:
			w.WriteHeader(http.StatusNotFound)
		case ErrInvalidArgument:
			w.WriteHeader(http.StatusBadRequest)
	*/
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.error() != nil {
		encodeError(ctx, e.error(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/xml; charset=utf-8")
	return xml.NewEncoder(w).Encode(response)
}

type errorer interface {
	error() error
}
