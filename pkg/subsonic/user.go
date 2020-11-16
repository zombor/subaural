package subsonic

import "encoding/xml"

type GetUserResponse struct {
	XMLName xml.Name `xml:"subsonic-response"`
	SubsonicResponse
	User User `xml:"user"`
}

type User struct {
	Username          string `xml:"username,attr"`
	Email             string `xml:"email,attr"`
	ScrobblingEnabled bool   `xml:"scrobblingEnabled,attr"`
	AdminRole         bool   `xml:"adminRole,attr"`
	SettingsRole      bool   `xml:"settingsRole,attr"`
	DownloadRole      bool   `xml:"downloadRole,attr"`
	UploadRole        bool   `xml:"uploadRole,attr"`
	PlaylistRole      bool   `xml:"playlistRole,attr"`
	CoverArtRole      bool   `xml:"coverArtRole,attr"`
	CommentRole       bool   `xml:"commentRole,attr"`
	PodcastRole       bool   `xml:"podcastRole,attr"`
	StreamRole        bool   `xml:"streamRole,attr"`
	JukeboxRole       bool   `xml:"jukeboxRole,attr"`
	ShareRole         bool   `xml:"shareRole,attr"`

	Folders []int `xml:"folder"`
}
