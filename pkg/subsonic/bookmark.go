package subsonic

import "encoding/xml"

type GetPlayQueue struct {
	XMLName xml.Name `xml:"subsonic-response"`
	SubsonicResponse
	PlayQueue PlayQueue `xml:"playQueue"`
}

type PlayQueue struct {
	Current   string `xml:"current,attr"`
	Position  string `xml:"position,attr"`
	Username  string `xml:"username,attr"`
	Changed   string `xml:"changed,attr"`
	ChangedBy string `xml:"changedBy,attr"`
	Entries   []Song
}

type Bookmarks struct {
	Bookmarks []Bookmark `xml:"bookmark"`
}

type Bookmark struct {
	Position string `xml:"position,attr"`
	Username string `xml:"username,attr"`
	Comment  string `xml:"comment,attr"`
	Created  string `xml:"created,attr"`
	Changed  string `xml:"changed,attr"`
	Entries  []Song `xml:"entry"`
}
