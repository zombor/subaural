package podcast

import (
	"context"
	"strings"
	"time"

	"github.com/go-kit/kit/endpoint"
	rss "github.com/ungerik/go-rss"

	"github.com/zombor/subaural/pkg/subsonic"
)

func makeGetPodcastsEndpoint(urls []string) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var (
			req      getPodcastsRequest = request.(getPodcastsRequest)
			channels []subsonic.Channel

			err error
		)

		for i, _ := range urls {
			resp, err := rss.Read(urls[i], false)
			if err != nil {
				return nil, err
			}

			ch, err := rss.Regular(resp)
			if err != nil {
				return nil, err
			}

			if req.ID != "" && req.ID != subsonic.PathID(urls[i]) {
				continue
			}

			var channel subsonic.Channel = subsonic.Channel{
				ID:          subsonic.PathID(urls[i]),
				URL:         urls[i],
				Title:       ch.Title,
				Description: ch.Description,
				Status:      "completed",
				//OriginalImageURL: ch.Image.URL,
			}

			if req.IncludeEpisodes {
				channel.Episodes = make([]subsonic.Episode, len(ch.Item))

				for j, _ := range ch.Item {
					pubDate, err := time.Parse("Mon, 2 Jan 2006 15:04:05 MST", strings.TrimSpace(string(ch.Item[j].PubDate)))
					if err != nil {
						return nil, err
					}

					if len(ch.Item[j].Enclosure) == 0 {
						continue
						//return nil, fmt.Errorf("podcast %s index %d is missing a url enclosure", urls[i], j)
					}

					channel.Episodes[j] = subsonic.Episode{
						ID:          subsonic.PathID(string(ch.Item[j].PubDate)),
						StreamID:    subsonic.PathID(ch.Item[j].Enclosure[0].URL),
						CoverArt:    "1",
						ChannelID:   subsonic.PathID(urls[i]),
						Title:       ch.Item[j].Title,
						Description: ch.Item[j].Description,
						Status:      "completed",
						Genre:       "Podcast",
						PublishDate: pubDate.Format(time.RFC3339),
						Year:        "2020",
						ContentType: "audio/mpeg",
						Suffix:      "mp3",
						//BitRate:     128,
						//Size:        100,
						//Duration:    100,
						Path: ch.Item[j].Enclosure[0].URL,
					}
				}
			}

			channels = append(channels, channel)
		}

		return getPodcastsResponse{SubsonicResponse: subsonic.SubsonicHeaders(), Podcasts: subsonic.Podcasts{Channels: channels}, Err: err}, nil
	}
}

type getPodcastsResponse struct {
	subsonic.SubsonicResponse
	Podcasts subsonic.Podcasts `xml:"podcasts"`
	Err      error
}
