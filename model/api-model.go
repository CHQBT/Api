package model

// User
type UserLogin struct {
	Login        string `json:"login"`
	PasswordHash string `json:"password"`
}

// Twit
type TwitResponse struct {
	ID           string          `json:"id"`
	UserID       string          `json:"user_id"`
	PublisherFio string          `json:"publisher_fio"`
	Type         string          `json:"type"`
	Texts        string          `json:"texts"`
	Title        string          `json:"title"`
	ReadersCount int32           `json:"readers_count"`
	CreatedAt    string          `json:"created_at"`
	Videos       []*VideoInfo    `json:"videos"`
	Photos       []*PhotoInfo    `json:"photos"`
	Musics       []*MusicInfo    `json:"musics"`
	Locations    []*LocationInfo `json:"locations"`
	Urls         []*UrlInfo      `json:"urls"`
}

// VideoInfo struct
type VideoInfo struct {
	VideoID  string `json:"video_id"`
	VideoURL string `json:"video_url"`
}

// PhotoInfo struct
type PhotoInfo struct {
	PhotoID  string `json:"photo_id"`
	PhotoURL string `json:"photo_url"`
}

// MusicInfo struct
type MusicInfo struct {
	MusicID  string `json:"music_id"`
	MusicURL string `json:"music_url"`
}

// LocationInfo struct
type LocationInfo struct {
	LocationID string `json:"location_id"`
	Latitude   string `json:"latitude"`
	Longitude  string `json:"longitude"`
}

// UrlInfo struct
type UrlInfo struct {
	UrlID string `json:"url_id"`
	Url   string `json:"url"`
}

type CreateTwitRequestApi struct {
	PublisherFIO string `json:"publisher_fio"`
	Type         string `json:"type"`
	Texts        string `json:"texts"`
	Title        string `json:"title"`
}
