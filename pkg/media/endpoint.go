package media

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

func makeStreamEndpoint(readMedia func(string) ([]byte, error)) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var (
			data []byte

			err error
		)

		req := request.(streamRequest)

		data, err = readMedia(req.ID)

		return data, err
	}
}

func makeGetCoverArtEndpoint(findCoverArt func(string) ([]byte, error)) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var (
			data []byte

			err error
		)

		req := request.(getCoverArtRequest)

		data, err = findCoverArt(req.ID)

		return data, err
	}
}
