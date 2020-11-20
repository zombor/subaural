package subsonic

import "encoding/xml"

type GetIndexes struct {
	XMLName xml.Name `xml:"subsonic-response"`
	SubsonicResponse
	Indexes Indexes `xml:"indexes"`
}

type Indexes struct {
	Indexes []Index `xml:"index"`
}

type Index struct {
	Name    string   `xml:"name,attr"`
	Artists []Artist `xml:"artists"`
}

type Artist struct {
	XMLName xml.Name `xml:"artist"`
	ID      string   `xml:"id,attr"`
	Name    string   `xml:"name,attr"`
}

type MusicDirectory struct {
	XMLName xml.Name `xml:"directory"`
	ID      string   `xml:"id,attr"`
	Parent  string   `xml:"parent,attr"`
	Name    string   `xml:"name,attr"`
	Starred string   `xml:"starred,attr"`

	Children []DirectoryChild
}

type DirectoryChild struct {
	XMLName xml.Name `xml:"child"`

	ID       string `xml:"id,attr"`
	Parent   string `xml:"parent,attr"`
	Title    string `xml:"title,attr"`
	Artist   string `xml:"artist,attr"`
	IsDir    bool   `xml:"isDir,attr"`
	CoverArt string `xml:"coverArt,attr"`

	Album       string `xml:"album,attr,omitempty"`
	Track       int    `xml:"track,attr,omitempty"`
	Year        int    `xml:"year,attr,omitempty"`
	Genre       string `xml:"genre,attr,omitempty"`
	Size        int    `xml:"size,attr,omitempty"`
	ContentType string `xml:"contentType,attr,omitempty"`
	Suffix      string `xml:"suffix,attr,omitempty"`
	Duration    int    `xml:"duration,attr,omitempty"`
	BitRate     int    `xml:"bitRate,attr,omitempty"`
	Path        string `xml:"path,attr,omitempty"`
}
