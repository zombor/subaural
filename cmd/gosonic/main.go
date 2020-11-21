package main

import (
	"flag"
	"fmt"
	"net/http"
	"net/http/httputil"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"
	kithttp "github.com/go-kit/kit/transport/http"

	"gitlab.com/jeremybush/gosonic/pkg/bookmarks"
	"gitlab.com/jeremybush/gosonic/pkg/browsing"
	"gitlab.com/jeremybush/gosonic/pkg/lists"
	"gitlab.com/jeremybush/gosonic/pkg/media"
	"gitlab.com/jeremybush/gosonic/pkg/system"
	"gitlab.com/jeremybush/gosonic/pkg/user"
	"gitlab.com/jeremybush/gosonic/pkg/xml"
)

func main() {
	var addr string
	{
		addr = os.Getenv("PORT")
		if addr == "" {
			addr = "8080"
		}
	}

	var (
		httpAddr = flag.String("http.addr", ":"+addr, "HTTP listen address")
	)

	var logger log.Logger
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)

	// httpLogger := log.With(logger, "component", "http")

	var bs browsing.Service
	{
		bs = browsing.NewService()
		bs = browsing.NewLoggingService(logger, bs)
	}

	var ms media.Service
	{
		ms = media.NewService()
		ms = media.NewLoggingService(logger, ms)
	}

	mux := http.NewServeMux()

	opts := []kithttp.ServerOption{
		kithttp.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		kithttp.ServerErrorEncoder(xml.EncodeError),
	}

	mux.Handle("/rest/ping.view", system.GetPingHandler(xml.EncodeResponse, opts))
	mux.Handle("/rest/getLicense.view", system.GetLicenseHandler(xml.EncodeResponse, opts))
	mux.Handle("/rest/getMusicFolders.view", browsing.GetMusicFoldersHandler(bs, xml.EncodeResponse, opts))
	mux.Handle("/rest/getMusicDirectory.view", browsing.GetMusicDirectory(bs, xml.EncodeResponse, opts))
	mux.Handle("/rest/getIndexes.view", browsing.GetIndexesHandler(bs, xml.EncodeResponse, opts))
	mux.Handle("/rest/getUser.view", user.GetUserHandler(xml.EncodeResponse, opts))
	mux.Handle("/rest/getAlbumList.view", lists.GetAlbumList(xml.EncodeResponse, opts))
	mux.Handle("/rest/getRandomSongs.view", lists.GetRandomSongs(xml.EncodeResponse, opts))
	mux.Handle("/rest/getPlayQueue.view", bookmarks.GetPlayQueue(xml.EncodeResponse, opts))
	mux.Handle("/rest/getBookmarks.view", bookmarks.GetBookmarks(xml.EncodeResponse, opts))
	mux.Handle("/rest/getAvatar.view", media.GetAvatar(xml.EncodeImageResponse, opts))
	mux.Handle("/rest/stream.view", media.Stream(ms.ReadMedia, xml.EncodeStreamResponse, opts))
	mux.Handle("/rest/getCoverArt.view", media.GetCoverArt(ms.FindCoverArt, xml.EncodeImageResponse, opts))
	mux.Handle(
		"/",
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			dump, err := httputil.DumpRequest(r, true)
			if err != nil {
				http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
				return
			}

			fmt.Printf("%q\n", dump)
		}),
	)

	http.Handle("/", mux)

	errs := make(chan error, 2)
	go func() {
		logger.Log("transport", "http", "address", *httpAddr, "msg", "listening")
		errs <- http.ListenAndServe(*httpAddr, nil)
	}()
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	logger.Log("terminated", <-errs)
}
