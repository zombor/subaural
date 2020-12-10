package bookmarks

import (
	"github.com/zombor/subaural/pkg/subsonic"
)

type getBookmarksResponse struct {
	subsonic.SubsonicResponse
	Bookmarks subsonic.Bookmarks `xml:"bookmarks"`
	Err       error
}
