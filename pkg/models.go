package pkg

// Track fields
type TrackDetail struct {
	Track_Id               int    `json:"track_id"`
	Track_Mbid             string `json:"track_mbid"`
	Track_Length           int    `json:"track_length"`
	Lyrics_Id              int    `json:"lyrics_id"`
	Instrumental           int    `json:"instrumental"`
	Subtitle_Id            int    `json:"subtitle_id"`
	Track_Name             string `json:"track_name"`
	Track_Rating           string `json:"track_rating"`
	Album_Name             string `json:"album_name"`
	Album_Id               int    `json:"album_id"`
	Artist_Id              int    `json:"artist_id"`
	Album_Coverart_100X100 string `json:"album_coverart_100x100"`
	Artist_Mbid            string `json:"artist_mbid"`
	Artist_Name            string `json:"artist_name"`
	Updated_Time           string `json:"updated_time"`
}

type Header struct {
	Status_Code  int     `json:"status_code"`
	Execute_Time float32 `json:"execute_time"`
	Available    int     `json:"available"`
}

type Track struct {
	Track TrackDetail `json:"track"`
}

type TrackList struct {
	TrackList []Track `json:"track_list"`
}

type Message struct {
	Header Header    `json:"header"`
	Body   TrackList `json:"body"`
}

type TrackResponse struct {
	Message Message `json:"message"`
}
