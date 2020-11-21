package user

import (
	"context"
	"net/http"

	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"

	"gitlab.com/jeremybush/gosonic/pkg/subsonic"
)

func GetUserHandler(
	encodeResponse func(ctx context.Context, w http.ResponseWriter, response interface{}) error,
	opts []kithttp.ServerOption,
) http.Handler {
	getUserHandler := kithttp.NewServer(
		func(ctx context.Context, request interface{}) (interface{}, error) {
			return subsonic.GetUserResponse{
				SubsonicResponse: subsonic.SubsonicResponse{
					Status: "ok", Version: "1.1.1",
				},
				User: subsonic.User{
					Folders: []int{1},
				},
			}, nil
		},
		func(_ context.Context, r *http.Request) (interface{}, error) { return nil, nil },
		encodeResponse,
		opts...,
	)

	r := mux.NewRouter()

	r.Handle("/rest/getUser.view", getUserHandler).Methods("GET")

	return r
}
