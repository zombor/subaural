package media

import (
	"context"

	"github.com/go-kit/kit/endpoint"

	"gitlab.com/jeremybush/gosonic/pkg/subsonic"
)

func makeStreamEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var (
			data []byte

			err error
		)

		req := request.(streamRequest)

		data, err = subsonic.ReadFile(req.ID)

		return streamResponse{Data: data, Err: err}, nil
	}
}

func makeGetCoverArtEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var (
			data []byte

			err error
		)

		req := request.(getCoverArtRequest)

		data, err = subsonic.FindCoverArt(req.ID)

		return data, err
	}
}
