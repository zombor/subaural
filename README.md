# Subaural

Subaural is a Subsonic api-compatible web server for serving music and podcasts. It's designed to be run with a Subsonic api compatible client such as D-Sub.

## Project Status

This project is still in heavy development. Not all Subsonic endpoints are implemented and some features for working endpoints might be sub-optimal.

There's currently no persistent data store support. This means your music directory are scanned on the fly and podcasts are loaded on-demand. None of the persistence features of the subsonic api are functional, such as bookmarks.

## Running / Configuration

Subarual uses environment variables for configuration:

- `SUBAURAL_MUSIC_PATH`: Full path to your music directory. Only one is supported right now
- `SUBAURAL_PODCAST_URLS`: Optional space separated value of RSS urls
- `PORT`: Optional http port. Defaults to `8080`.

Run it with `go run ./cmd/subaural/main.go`
