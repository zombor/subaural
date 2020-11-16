package subsonic

import "encoding/xml"

type SubsonicResponse struct {
	XMLName xml.Name `xml:"subsonic-response"`
	Status  string   `xml:"status,attr"`
	Version string   `xml:"version,attr"`
}

func SubsonicHeaders() SubsonicResponse {
	return SubsonicResponse{Status: "ok", Version: "1.16.0"}
}

type MusicFolders struct {
	MusicFolders []MusicFolder `xml:"musicFolders"`
}

type MusicFolder struct {
	XMLName xml.Name `xml:"musicFolder"`
	ID      string   `xml:"id,attr"`
	Name    string   `xml:"name,attr"`
}

type GetLicenseResponse struct {
	XMLName xml.Name `xml:"subsonic-response"`
	SubsonicResponse
	License License `xml:"license"`
}

type License struct {
	Valid          bool   `xml:"valid,attr"`
	Email          string `xml:"email,attr"`
	LicenseExpires string `xml:"licenseExpires,attr"`
}
