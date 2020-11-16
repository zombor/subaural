package bookmarks

import (
	"gitlab.com/jeremybush/gosonic/pkg/subsonic"
)

type getBookmarksResponse struct {
	subsonic.SubsonicResponse
	Bookmarks subsonic.Bookmarks `xml:"bookmarks"`
	Err       error
}
