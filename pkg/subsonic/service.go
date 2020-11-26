package subsonic

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/go-flac/flacvorbis"
	flac "github.com/go-flac/go-flac"
)

type Service interface {
	ParseFlac(string, string) (bool, FlacMeta, error)
	ReadFile(string, int) ([]byte, error)
	FindCoverArt(string) ([]byte, error)
	ReadDir(string) ([]os.FileInfo, error)
}

type service struct {
	musicPath string
}

func NewService(musicPath string) service {
	return service{musicPath}
}

func (s service) ParseFlac(parent, fileName string) (bool, FlacMeta, error) {
	var (
		decoded []byte

		meta FlacMeta

		err error
	)

	decoded, err = base64.RawStdEncoding.DecodeString(parent)

	if err != nil {
		return false, meta, err
	}

	fileName = fmt.Sprintf("%s/%s/%s", s.musicPath, decoded, fileName)

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

func (s service) ReadFile(path string, rate int) ([]byte, error) {
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
		cmd := exec.Command("ffmpeg", "-i", fmt.Sprintf("%s/%s", s.musicPath, decoded), "-f", "mp3", "-b:a", fmt.Sprintf("%dk", rate), "pipe:1")
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

	return ioutil.ReadFile(fmt.Sprintf("%s/%s", s.musicPath, decoded))
}

func (s service) FindCoverArt(path string) ([]byte, error) {
	var (
		decoded []byte

		err error
	)

	decoded, err = base64.RawStdEncoding.DecodeString(path)
	if err != nil {
		return nil, err
	}

	pathParts := strings.Split(string(decoded), "/")

	for len(pathParts) > 0 {
		path = strings.Join(pathParts, "/")

		if ok, data, err := findCoverArt(fmt.Sprintf("%s/%s/cover.png", s.musicPath, path)); ok {
			return data, err
		}
		if ok, data, err := findCoverArt(fmt.Sprintf("%s/%s/cover.jpg", s.musicPath, path)); ok {
			return data, err
		}
		if ok, data, err := findCoverArt(fmt.Sprintf("%s/%s/Front.jpg", s.musicPath, path)); ok {
			return data, err
		}
		if ok, data, err := findCoverArt(fmt.Sprintf("%s/%s/folder.jpg", s.musicPath, path)); ok {
			return data, err
		}
		if ok, data, err := findCoverArt(fmt.Sprintf("%s/%s/Folder.jpg", s.musicPath, path)); ok {
			return data, err
		}

		pathParts = pathParts[:len(pathParts)-1]
	}

	return nil, nil
}

func (s service) ReadDir(dir string) ([]os.FileInfo, error) {
	var (
		decoded []byte

		err error
	)

	decoded, err = base64.RawStdEncoding.DecodeString(dir)
	if err != nil {
		return nil, err
	}

	return ioutil.ReadDir(fmt.Sprintf("%s/%s", s.musicPath, decoded))
}

func findCoverArt(path string) (bool, []byte, error) {
	_, err := os.Stat(path)

	if err != nil {
		return false, nil, nil
	}

	data, err := ioutil.ReadFile(path)

	return true, data, err
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
