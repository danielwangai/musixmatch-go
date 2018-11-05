package pkg

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

// GetTopTracks fetches a list of top trakcs of a given country
// Reference https://developer.musixmatch.com/documentation/api-reference/track-chart-get
// Requires Authentication
// currently works with format=json
// TODO - XML
func GetTopTracks(country, chartName string, page, pageSize, hasLyrics int) (*TrackMusicResponse, error) {
	if country == "" {
		// default country is US
		country = "us"
	}
	if chartName == "" || !Contains(chartName, []string{"top", "hot", "mxmweekly", "mxmweekly_new"}) {
		return nil, errors.New("Invalid Chart name. Can only be - top, hot, mxmweekly, mxmweekly_new")
	}
	if page <= 0 || pageSize <= 0 {
		// set default values
		page = 1
		pageSize = 1
	}
	if hasLyrics != 1 || hasLyrics != 0 {
		// set default to 1
		hasLyrics = 1
	}
	topTrackUrl := BaseURL + "chart.tracks.get"
	u, err := url.Parse(topTrackUrl)
	if err != nil {
		return nil, err
	}
	q, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		return nil, err
	}
	q.Add("chart_name", chartName)
	q.Add("page", strconv.Itoa(page))
	q.Add("page_size", strconv.Itoa(pageSize))
	q.Add("country", country)
	q.Add("f_has_lyrics", strconv.Itoa(hasLyrics))
	q.Add("apikey", "aa4afda40b0f535b659af399decb21e4")
	u.RawQuery = q.Encode()
	resp, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}
	var trackMusicResp TrackMusicResponse
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(data, &trackMusicResp)
	return &trackMusicResp, nil
}
