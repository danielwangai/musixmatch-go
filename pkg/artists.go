package pkg

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

// GetTopArtists fetches a list of top artists of a given country
// Reference https://developer.musixmatch.com/documentation/api-reference/artist-chart-get
// Requires Authentication
// currently works with format=json
// TODO - XML
func GetTopArtists(country, format string, page, pageSize int) (*ArtistResponse, error) {
	// country - enter valid country code. Default is US
	if country == "" {
		country = "US"
	}
	if format == "" {
		// set to default format
		format = "json"
	}
	if format != "json" {
		return nil, errors.New("Invalid format. Currently accepting json.")
	}
	if page <= 0 || pageSize <= 0 {
		return nil, errors.New("Page and page size values must be positive integers.")
	}
	topArtistUrl := BaseURL + "chart.artists.get"
	u, err := url.Parse(topArtistUrl)
	if err != nil {
		return nil, err
	}
	q, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		return nil, err
	}
	q.Add("apikey", APIKEY)
	q.Add("page", strconv.Itoa(page))
	q.Add("page_size", strconv.Itoa(pageSize))
	q.Add("format", format)
	u.RawQuery = q.Encode()
	resp, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}
	var artists ArtistResponse
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(data, &artists)
	return &artists, nil
}
