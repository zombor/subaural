package podcast

import (
	"context"
	"net/http"

	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func GetPodcasts(
	podcastUrls []string,
	encodeResponse func(ctx context.Context, w http.ResponseWriter, response interface{}) error,
	opts []kithttp.ServerOption,
) http.Handler {
	getPodcastsHandler := kithttp.NewServer(
		makeGetPodcastsEndpoint(podcastUrls),
		decodeGetPodcastsRequest,
		encodeResponse,
		opts...,
	)

	r := mux.NewRouter()

	r.Handle("/rest/getPodcasts.view", getPodcastsHandler).Methods("GET")

	return r
}

func decodeGetPodcastsRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var (
		id              string
		includeEpisodes bool = true
	)

	query := r.URL.Query()
	idq, ok := query["id"]
	if ok {
		id = idq[0]
	}
	ie, ok := query["includeEpisodes"]
	if ok {
		includeEpisodes = ie[0] == "true"
	}

	return getPodcastsRequest{ID: id, IncludeEpisodes: includeEpisodes}, nil
}

type getPodcastsRequest struct {
	ID              string
	IncludeEpisodes bool
}
