package pkg

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
	ExecuteTime float32 `json:"execute_time"`
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
	Body AlbumBody `json:"body"`
}

type AlbumResponse struct {
	Message AlbumMessage `json:"message"`
}
