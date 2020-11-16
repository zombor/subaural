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
