package subsonic

import "encoding/xml"

type Podcasts struct {
	XMLName  xml.Name `xml:"podcasts"`
	Channels []Channel
}

type Channel struct {
	XMLName xml.Name `xml:"channel"`

	ID          string    `xml:"id,attr"`
	URL         string    `xml:"url,attr"`
	Title       string    `xml:"title,attr"`
	Description string    `xml:"description,attr"`
	Status      string    `xml:"status,attr"`
	CoverArt    string    `xml:"coverArt,attr"`
	Episodes    []Episode `xml:"episode"`
}

type Episode struct {
	XMLName     xml.Name `xml:"episode"`
	ID          string   `xml:"id,attr"`
	StreamID    string   `xml:"streamId,attr"`
	ChannelID   string   `xml:"channelId,attr"`
	Title       string   `xml:"title,attr"`
	Description string   `xml:"description,attr"`
	Status      string   `xml:"status,attr"`
	PublishDate string   `xml:"publishDate,attr"`
	Year        string   `xml:"year,attr"`
	Genre       string   `xml:"genre,attr"`
	CoverArt    string   `xml:"coverArt,attr"`
	Size        int      `xml:"size,attr"`
	ContentType string   `xml:"contentType,attr"`
	Suffix      string   `xml:"suffix,attr"`
	Duration    int      `xml:"duration,attr"`
	BitRate     int      `xml:"bitRate,attr"`
	Path        string   `xml:"path,attr"`
}
