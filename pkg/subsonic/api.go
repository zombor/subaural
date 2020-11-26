package subsonic

import (
	"encoding/base64"
	"encoding/xml"
	"strings"
)

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

func PathID(s string) string {
	return base64.RawStdEncoding.EncodeToString([]byte(s))
}

func ParentID(s string) (string, string, error) {
	decoded, err := base64.RawStdEncoding.DecodeString(s)

	if err != nil {
		return "", "", err
	}

	parts := strings.Split(string(decoded), "/")

	// top level
	if len(parts) == 1 {
		return string(decoded), "", nil
	}

	return strings.Join(parts[:len(parts)-1], "/"), parts[len(parts)-1], nil
}

type FlacMeta struct {
	Artist string
	Title  string
	Album  string
	Date   int
	Track  int
	Genre  string
}
