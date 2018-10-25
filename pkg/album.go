package pkg

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

// get tracks from an album
func GetAlbumTracks(albumId, page, limit int) (*TrackResponse, error) {
	albumTracksUrl := BaseURL + "album.tracks.get"
	u, err := url.Parse(albumTracksUrl)
	if err != nil {
		return nil, err
	}
	q, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		return nil, err
	}
	q.Add("apikey", APIKEY)
	// convert integers to string
	q.Add("album_id", strconv.Itoa(albumId))
	q.Add("page", strconv.Itoa(page))
	q.Add("page_size", strconv.Itoa(limit))
	// add params to url
	u.RawQuery = q.Encode()
	//fetch artists
	resp, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var tracks TrackResponse
	json.Unmarshal(data, &tracks)
	return &tracks, nil
}
