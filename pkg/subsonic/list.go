package subsonic

import "encoding/xml"

type GetAlbumList struct {
	XMLName xml.Name `xml:"subsonic-response"`
	SubsonicResponse
	AlbumList AlbumLists `xml:"musicFolders"`
}

type AlbumLists struct {
	AlbumList []Album `xml:"albumList"`
}

type Album struct {
	XMLName       xml.Name `xml:"album"`
	ID            string   `xml:"id,attr"`
	Parent        string   `xml:"parent,attr"`
	Title         string   `xml:"title,attr"`
	Artist        string   `xml:"artist,attr"`
	IsDir         bool     `xml:"isDir,attr"`
	CoverArt      int      `xml:"coverArt,attr"`
	UserRating    string   `xml:"userRating,attr"`
	AverageRating string   `xml:"averageRating,attr"`
}

type GetRandomSongs struct {
	XMLName xml.Name `xml:"subsonic-response"`
	SubsonicResponse
	RandomSongs RandomSongs `xml:"randomSongs"`
}

type RandomSongs struct {
	RandomSongs []Song `xml:"randomSongs"`
}

type Song struct {
	XMLName     xml.Name `xml:"song"`
	ID          string   `xml:"id,attr"`
	Parent      string   `xml:"parent,attr"`
	Title       string   `xml:"title,attr"`
	IsDir       bool     `xml:"isDir,attr"`
	Album       string   `xml:"album,attr"`
	Artist      string   `xml:"artist,attr"`
	Track       string   `xml:"track,attr"`
	Year        string   `xml:"year,attr"`
	Genre       string   `xml:"genre,attr"`
	CoverArt    string   `xml:"coverArt,attr"`
	Size        string   `xml:"size,attr"`
	ContentType string   `xml:"contentType,attr"`
	Suffix      string   `xml:"suffix,attr"`
	Duration    string   `xml:"duration,attr"`
	Bitrate     string   `xml:"bitrate,attr"`
	Path        string   `xml:"path,attr"`
}
