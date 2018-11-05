package pkg

// Artist
type ArtistDetail struct {
	ArtistID        int      `json:"artist_id"`
	ArtistMbID      string   `json:"artist_mbid"`
	ArtistName      string   `json:"artist_name"`
	ArtistAliasList []string `json:"artist_alias_list"`
	ArtistRating    int      `json:"artist_rating"`
	UpdatedTime     string   `json:"updated_time"`
}

type Artist struct {
	Artist ArtistDetail `json:"artist"`
}

type ArtistList struct {
	ArtistList []Artist `json:"artist_list"`
}

type ArtistHeader struct {
	StatusCode  int     `json:"status_code"`
	ExecuteTime float64 `json:"execute_time"`
}

type ArtistMessage struct {
	Header ArtistHeader `json:"header"`
	Body   ArtistList   `json:"body"`
}

type ArtistResponse struct {
	Message ArtistMessage `json:"message"`
}

// chart.tracks.get

type TrackMusicGenreDetails struct {
	MusicGenreID           int    `json:"music_genre_id"`
	MusicGenreParentID     int    `json:"music_genre_parent_id"`
	MusicGenreName         string `json:"music_genre_name"`
	MusicGenreNameExtended string `json:"music_genre_name_extended"`
	MusicGenreVanity       string `json:"music_genre_vanity"`
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

type TrackMusicHeader struct {
	StatusCode  int     `json:"status_code"`
	ExecuteTime float64 `json:"execute_time"`
}

type TrackMusicMessage struct {
	Header TrackMusicHeader `json:"header"`
	Body   TrackMusicList   `json:"body"`
}

type TrackMusicResponse struct {
	Message TrackMusicMessage `json:"message"`
}

// end of chart.tracks.get

// Track fields
type TrackDetail struct {
	TrackId              int    `json:"track_id"`
	TrackMbId            string `json:"track_mbid"`
	TrackLength          int    `json:"track_length"`
	LyricsId             int    `json:"lyrics_id"`
	Instrumental         int    `json:"instrumental"`
	SubtitleId           int    `json:"subtitle_id"`
	TrackName            string `json:"track_name"`
	TrackRating          string `json:"track_rating"`
	AlbumName            string `json:"album_name"`
	AlbumId              int    `json:"album_id"`
	ArtistId             int    `json:"artist_id"`
	AlbumCoverArt100X100 string `json:"album_coverart_100x100"`
	ArtistMbId           string `json:"artist_mbid"`
	ArtistName           string `json:"artist_name"`
	UpdatedTime          string `json:"updated_time"`
}

type Header struct {
	StatusCode  int     `json:"status_code"`
	ExecuteTime float64 `json:"execute_time"`
	Available   int     `json:"available"`
}

type Track struct {
	Track TrackDetail `json:"track"`
}

//the album
type TrackList struct {
	TrackList []Track `json:"track_list"`
}

type TrackMessage struct {
	Header Header    `json:"header"`
	Body   TrackList `json:"body"`
}

// album response
type TrackResponse struct {
	Message TrackMessage `json:"message"`
}

type Genre struct {
	MusicGenreId           int    `json:"music_genre_id"`
	MusicGenreParentId     int    `json:"music_genre_parent_id"`
	MusicGenreName         string `json:"music_genre_name"`
	MusicGenreNameExtended string `json:"music_genre_name_extended"`
}

type MusicGenre struct {
	MusicGenre Genre `json:"music_genre"`
}

type MusicGenreList struct {
	MusicGenreList []MusicGenre `json:"music_genre_list"`
}

type AlbumDetails struct {
	AlbumId              int            `json:"album_id"`
	AlbumMbId            string         `json:"album_mbid"`
	AlbumName            string         `json:"album_name"`
	AlbumRating          int            `json:"album_rating"`
	AlbumTrackCount      int            `json:"album_track_count"`
	AlbumReleaseDate     string         `json:"album_release_date"`
	AlbumReleaseType     string         `json:"album_release_type"`
	ArtistId             int            `json:"artist_id"`
	ArtistName           string         `json:"artist_name"`
	PrimaryGenres        MusicGenreList `json:"primary_genres"`
	SecondaryGenres      MusicGenreList `json:"secondary_genres"`
	AlbumPLine           string         `json:"album_pline"`
	AlbumCopyright       string         `json:"album_copyright"`
	AlbumLabel           string         `json:"album_label"`
	UpdatedTime          string         `json:"updated_time"`
	AlbumCoverArt100X100 string         `json:"album_coverart_100x100"`
}

type AlbumBody struct {
	Album AlbumDetails `json:"album"`
}

type AlbumHeader struct {
	StatusCode  int     `json:"status_code"`
	ExecuteTime float64 `json:"execute_time"`
}

type AlbumMessage struct {
	Header AlbumHeader `json:"header"`
	Body   AlbumBody   `json:"body"`
}

type AlbumResponse struct {
	Message AlbumMessage `json:"message"`
}

// tracking.url.get
type TrackingURL struct {
	URL string `json:"url"`
}

type TrackingURLMessage struct {
	Header AlbumHeader `json:"header"`
	Body   TrackingURL `json:"body"`
}

type TrackingURLResponse struct {
	Message TrackingURLMessage `json:"message"`
}
