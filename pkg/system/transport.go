package system

import (
	"context"
	"net/http"

	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"

	"github.com/zombor/subaural/pkg/subsonic"
)

func GetPingHandler(
	encodeResponse func(ctx context.Context, w http.ResponseWriter, response interface{}) error,
	opts []kithttp.ServerOption,
) http.Handler {
	getPingHandler := kithttp.NewServer(
		func(ctx context.Context, request interface{}) (interface{}, error) {
			return subsonic.SubsonicResponse{Status: "ok", Version: "1.16.1"}, nil
		},
		func(_ context.Context, r *http.Request) (interface{}, error) { return nil, nil },
		encodeResponse,
		opts...,
	)

	r := mux.NewRouter()

	r.Handle("/rest/ping.view", getPingHandler).Methods("GET")

	return r
}

func GetLicenseHandler(
	encodeResponse func(ctx context.Context, w http.ResponseWriter, response interface{}) error,
	opts []kithttp.ServerOption,
) http.Handler {
	getLicenseHandler := kithttp.NewServer(
		func(ctx context.Context, request interface{}) (interface{}, error) {
			return subsonic.GetLicenseResponse{
				SubsonicResponse: subsonic.SubsonicResponse{
					Status: "ok", Version: "1.16.1",
				},
				License: subsonic.License{
					Valid:          true,
					Email:          "foo@bar.com",
					LicenseExpires: "3019-09-03T14:46:43",
				},
			}, nil
		},
		func(_ context.Context, r *http.Request) (interface{}, error) { return nil, nil },
		encodeResponse,
		opts...,
	)

	r := mux.NewRouter()

	r.Handle("/rest/getLicense.view", getLicenseHandler).Methods("GET")

	return r
}
