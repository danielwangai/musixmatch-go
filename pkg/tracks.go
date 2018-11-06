package pkg

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

// models
// chart.tracks.get
type TrackMusicGenreDetails struct {
	Genre
	// MusicGenreVanity string `json:"music_genre_vanity"`
}

type TrackMusicGenre struct {
	MusicGenre TrackMusicGenreDetails `json:"music_genre"`
}

type TrackMusicGenreList struct {
	MusicGenreList []TrackMusicGenre `json:"music_genre_list"`
}

type TrackMusicDetail struct {
	TrackId                 int                 `json:"track_id"`
	TrackMbID               string              `json:"track_mbid"`
	TrackIsrc               int                 `json:"track_isrc"`
	TrackSpotifyID          string              `json:"track_spotify_id"`
	TrackSoundcloudID       string              `json:"track_soundcloud_id"`
	TrackXboxMusicID        string              `json:"track_xboxmusic_id"`
	TrackName               string              `json:"track_name"`
	TrackNameTransitionList []string            `json:"track_name_translation_list"` // TODO - verify data type
	TrackRating             string              `json:"track_rating"`
	TrackLength             int                 `json:"track_length"`
	CommonTrackID           int                 `json:"commontrack_id"`
	Instrumental            int                 `json:"instrumental"`
	Explicit                int                 `json:"explicit"`
	HasLyrics               int                 `json:"has_lyrics"`
	HasLyricsCrowd          int                 `json:"has_lyrics_crowd"`
	HasSubtitles            int                 `json:"has_subtitles"`
	HasRichSync             int                 `json:"has_richsync"`
	NumFavourites           int                 `json:"num_favourite"`
	LyricsID                int                 `json:"lyrics_id"`
	SubtitleID              int                 `json:"subtitle_id"`
	AlbumID                 int                 `json:"album_id"`
	AlbumName               string              `json:"album_name"`
	ArtistID                int                 `json:"artist_id"`
	ArtistMbID              string              `json:"artist_mbid"`
	ArtistName              string              `json:"artist_name"`
	AlbumCoverArt100X100    string              `json:"album_coverart_100x100"`
	AlbumCoverArt350X350    string              `json:"album_coverart_350x350"`
	AlbumCoverArt500X500    string              `json:"album_coverart_500x500"`
	AlbumCoverArt800X800    string              `json:"album_coverart_800x800"`
	TrackShareURL           string              `json:"track_share_url"`
	TrackEditURL            string              `json:"track_edit_url"`
	CommonTrackVanityID     string              `json:"commontrack_vanity_id"`
	Restricted              int                 `json:"restricted"`
	FirstReleaseDate        string              `json:"first_release_date"`
	UpdateTime              string              `json:"updated_time"`
	PrimaryGenre            TrackMusicGenreList `json:"primary_genre"`
	SecondaryGenre          TrackMusicGenreList `json:"secondary_genre"`
}
type TrackMusic struct {
	Track TrackMusicDetail `json:"track"`
}

type TrackMusicList struct {
	TrackList []TrackMusic `json:"track_list"`
}

type TrackMusicMessage struct {
	Header Header         `json:"header"`
	Body   TrackMusicList `json:"body"`
}

type TrackMusicResponse struct {
	Message TrackMusicMessage `json:"message"`
}

// end chart.tracks.get
// track.search
type TrackSearchHeader struct {
	Header
	Available int `json:"available"`
}

type TrackSearchMessage struct {
	Header TrackSearchHeader `json:"header"`
	Body   TrackMusicList    `json:"body"`
}

type TrackSearchResponse struct {
	Message TrackSearchMessage `json:"message"`
}

// end track.search

// GetTopTracks fetches a list of top trakcs of a given country
// Reference https://developer.musixmatch.com/documentation/api-reference/track-chart-get
// Requires Authentication
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
	q.Add("apikey", APIKEY)
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

// GetTopTracks fetches a list of top trakcs of a given country
// Reference https://developer.musixmatch.com/documentation/api-reference/track-chart-get
// Requires Authentication
// currently works with format=json
// TODO - XML
// sample URL - http://api.musixmatch.com/ws/1.1/track.search?q_artist=justin bieber&page_size=3&page=1&s_track_rating=desc&apikey=APIKEY
func SearchTracks(artist, order string, page, pageSize int) (*TrackSearchResponse, error) {
	if artist == "" {
		return nil, errors.New("Provide artist name to search.")
	}
	if order == "" || !Contains(order, []string{"desc", "asc"}) {
		// set to default
		order = "desc"
	}
	if page <= 0 || pageSize <= 0 {
		// set default values
		page = 1
		pageSize = 1
	}
	searchUrl := BaseURL + "track.search"
	u, err := url.Parse(searchUrl)
	if err != nil {
		return nil, err
	}
	q, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		return nil, err
	}
	q.Add("q_artist", artist)
	q.Add("s_track_rating", order)
	q.Add("page", strconv.Itoa(page))
	q.Add("page_size", strconv.Itoa(pageSize))
	q.Add("apikey", APIKEY)
	u.RawQuery = q.Encode()
	resp, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}
	var trackSearch TrackSearchResponse
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(data, &trackSearch)
	return &trackSearch, nil
}
