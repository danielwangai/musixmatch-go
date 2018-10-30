package pkg

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

// GetTrackingUrl fetches tracking url details
// Reference - https://developer.musixmatch.com/documentation/api-reference/tracking-url-get
// Required Authentication
// currently works with format=json
// TODO - XML
func GetTrackingURL(domain, format string) (*TrackingURLResponse, error) {
	trackingUrl := BaseURL + "tracking.url.get"
	u, err := url.Parse(trackingUrl)
	if err != nil {
		return nil, err
	}
	q, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		return nil, err
	}
	q.Add("apikey", APIKEY)
	q.Add("domain", domain)
	q.Add("format", format)
	u.RawQuery = q.Encode()
	resp, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}
	var tUrl TrackingURLResponse
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(data, &tUrl)
	return &tUrl, nil
}
