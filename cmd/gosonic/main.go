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

	"gitlab.com/jeremybush/gosonic/pkg/bookmarks"
	"gitlab.com/jeremybush/gosonic/pkg/browsing"
	"gitlab.com/jeremybush/gosonic/pkg/lists"
	"gitlab.com/jeremybush/gosonic/pkg/media"
	"gitlab.com/jeremybush/gosonic/pkg/system"
	"gitlab.com/jeremybush/gosonic/pkg/user"
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

	httpLogger := log.With(logger, "component", "http")

	var bs browsing.Service
	{
		bs = browsing.NewService()
	}

	mux := http.NewServeMux()

	mux.Handle("/rest/ping.view", system.GetPingHandler(httpLogger))
	mux.Handle("/rest/getLicense.view", system.GetLicenseHandler(httpLogger))
	mux.Handle("/rest/getMusicFolders.view", browsing.GetMusicFoldersHandler(bs, httpLogger))
	mux.Handle("/rest/getMusicDirectory.view", browsing.GetMusicDirectory(bs, httpLogger))
	mux.Handle("/rest/getIndexes.view", browsing.GetIndexesHandler(bs, httpLogger))
	mux.Handle("/rest/getUser.view", user.GetUserHandler(httpLogger))
	mux.Handle("/rest/getAlbumList.view", lists.GetAlbumList(httpLogger))
	mux.Handle("/rest/getRandomSongs.view", lists.GetRandomSongs(httpLogger))
	mux.Handle("/rest/getPlayQueue.view", bookmarks.GetPlayQueue(httpLogger))
	mux.Handle("/rest/getBookmarks.view", bookmarks.GetBookmarks(httpLogger))
	mux.Handle("/rest/getAvatar.view", media.GetAvatar(httpLogger))
	mux.Handle("/rest/stream.view", media.Stream(httpLogger))
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
