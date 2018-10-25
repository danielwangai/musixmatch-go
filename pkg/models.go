package pkg

// Track fields
type TrackDetail struct {
	TrackId               int    `json:"track_id"`
	TrackMbId             string `json:"track_mbid"`
	TrackLength           int    `json:"track_length"`
	LyricsId              int    `json:"lyrics_id"`
	Instrumental           int    `json:"instrumental"`
	SubtitleId            int    `json:"subtitle_id"`
	TrackName             string `json:"track_name"`
	TrackRating           string `json:"track_rating"`
	AlbumName             string `json:"album_name"`
	AlbumId               int    `json:"album_id"`
	ArtistId              int    `json:"artist_id"`
	AlbumCoverArt100X100 string `json:"album_coverart_100x100"`
	ArtistMbId            string `json:"artist_mbid"`
	ArtistName            string `json:"artist_name"`
	UpdatedTime           string `json:"updated_time"`
}

type Header struct {
	StatusCode  int     `json:"status_code"`
	ExecuteTime float32 `json:"execute_time"`
	Available    int     `json:"available"`
}

type Track struct {
	Track TrackDetail `json:"track"`
}

//the album
type TrackList struct {
	TrackList []Track `json:"track_list"`
}

type Message struct {
	Header Header    `json:"header"`
	Body   TrackList `json:"body"`
}

// album response
type TrackResponse struct {
	Message Message `json:"message"`
}
