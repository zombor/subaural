package browsing

import (
	"context"
	"errors"
	"net/http"

	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func GetMusicFoldersHandler(
	bs Service,
	encodeResponse func(ctx context.Context, w http.ResponseWriter, response interface{}) error,
	opts []kithttp.ServerOption,
) http.Handler {
	getMusicFoldersHandler := kithttp.NewServer(
		makeGetMusicFoldersEndpoint(bs),
		func(_ context.Context, r *http.Request) (interface{}, error) { return nil, nil },
		encodeResponse,
		opts...,
	)

	r := mux.NewRouter()

	r.Handle("/rest/getMusicFolders.view", getMusicFoldersHandler).Methods("GET")

	return r
}

func GetIndexesHandler(
	bs Service,
	encodeResponse func(ctx context.Context, w http.ResponseWriter, response interface{}) error,
	opts []kithttp.ServerOption,
) http.Handler {
	getIndexesHandler := kithttp.NewServer(
		makeGetIndexesEndpoint(bs),
		func(_ context.Context, r *http.Request) (interface{}, error) { return nil, nil },
		encodeResponse,
		opts...,
	)

	r := mux.NewRouter()

	r.Handle("/rest/getIndexes.view", getIndexesHandler).Methods("GET")

	return r
}

func GetMusicDirectory(
	bs Service,
	encodeResponse func(ctx context.Context, w http.ResponseWriter, response interface{}) error,
	opts []kithttp.ServerOption,
) http.Handler {
	getMusicDirectory := kithttp.NewServer(
		makeGetMusicDirectoryEndpoint(bs),
		decodeGetMusicDirectoryRequest,
		encodeResponse,
		opts...,
	)

	r := mux.NewRouter()

	r.Handle("/rest/getMusicDirectory.view", getMusicDirectory).Methods("GET")

	return r
}

func decodeGetMusicDirectoryRequest(_ context.Context, r *http.Request) (interface{}, error) {
	query := r.URL.Query()
	id, ok := query["id"]
	if !ok {
		return nil, errors.New("missing id")
	}
	return getMusicDirectoryRequest{ID: id[0]}, nil
}
