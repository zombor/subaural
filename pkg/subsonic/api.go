package subsonic

import (
	"encoding/base64"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/go-flac/flacvorbis"
	flac "github.com/go-flac/go-flac"
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

func ReadDir(s string) ([]os.FileInfo, error) {
	var (
		decoded []byte

		err error
	)

	decoded, err = base64.RawStdEncoding.DecodeString(s)
	if err != nil {
		return nil, err
	}

	return ioutil.ReadDir(fmt.Sprintf("/mnt/media/music/%s", decoded))
}

type FlacMeta struct {
	Artist string
	Title  string
	Album  string
	Date   int
	Track  int
	Genre  string
}

func parseFlacMeta(c *flacvorbis.MetaDataBlockVorbisComment) FlacMeta {
	var meta FlacMeta

	for i, _ := range c.Comments {
		parts := strings.Split(c.Comments[i], "=")

		if parts[0] == "ARTIST" {
			meta.Artist = parts[1]
		}
		if parts[0] == "TITLE" {
			meta.Title = parts[1]
		}
		if parts[0] == "ALBUM" {
			meta.Album = parts[1]
		}
		if parts[0] == "DATE" {
			meta.Date, _ = strconv.Atoi(parts[1])
		}
		if parts[0] == "TRACKNUMBER" {
			meta.Track, _ = strconv.Atoi(parts[1])
		}
		if parts[0] == "GENRE" {
			meta.Genre = parts[1]
		}
	}

	return meta
}

func ParseFlac(parent, fileName string) (bool, FlacMeta, error) {
	var (
		decoded []byte

		meta FlacMeta

		err error
	)

	decoded, err = base64.RawStdEncoding.DecodeString(parent)

	if err != nil {
		return false, meta, err
	}

	fileName = fmt.Sprintf("/mnt/media/music/%s/%s", decoded, fileName)

	f, err := flac.ParseFile(fileName)
	if err != nil {
		return false, meta, err
	}

	for i, _ := range f.Meta {
		var cmt *flacvorbis.MetaDataBlockVorbisComment

		if f.Meta[i].Type == flac.VorbisComment {
			cmt, err = flacvorbis.ParseFromMetaDataBlock(*f.Meta[i])
			if err != nil {
				return false, meta, err
			}

			meta = parseFlacMeta(cmt)
		}
	}

	return true, meta, nil
}

func ReadFile(path string, rate int) ([]byte, error) {
	var (
		decoded []byte
		//data    []byte

		err error
	)

	decoded, err = base64.RawStdEncoding.DecodeString(path)
	if err != nil {
		return nil, err
	}

	if rate > 0 {
		cmd := exec.Command("ffmpeg", "-i", fmt.Sprintf("/mnt/media/music/%s", decoded), "-f", "mp3", "-b:a", fmt.Sprintf("%dk", rate), "pipe:1")
		stdout, err := cmd.StdoutPipe()
		if err != nil {
			return nil, err
		}
		err = cmd.Start()
		if err != nil {
			return nil, err
		}

		data, err := ioutil.ReadAll(stdout)
		if err != nil {
			return nil, err
		}

		return data, cmd.Wait()
	}

	return ioutil.ReadFile(fmt.Sprintf("/mnt/media/music/%s", decoded))
}

func FindCoverArt(path string) ([]byte, error) {
	var (
		decoded []byte

		err error
	)

	decoded, err = base64.RawStdEncoding.DecodeString(path)
	if err != nil {
		return nil, err
	}

	if ok, data, err := findCoverArt(fmt.Sprintf("/mnt/media/music/%s/cover.png", decoded)); ok {
		return data, err
	}
	if ok, data, err := findCoverArt(fmt.Sprintf("/mnt/media/music/%s/cover.jpg", decoded)); ok {
		return data, err
	}
	if ok, data, err := findCoverArt(fmt.Sprintf("/mnt/media/music/%s/Front.jpg", decoded)); ok {
		return data, err
	}
	if ok, data, err := findCoverArt(fmt.Sprintf("/mnt/media/music/%s/Folder.jpg", decoded)); ok {
		return data, err
	}

	return nil, nil
}

func findCoverArt(path string) (bool, []byte, error) {
	_, err := os.Stat(path)

	if err != nil {
		return false, nil, nil
	}

	data, err := ioutil.ReadFile(path)

	return true, data, err
}
