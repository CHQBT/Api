package model

type CreateTwitRequest struct {
	UserID string `json:"user_id"`
	Texts  string `json:"texts"`
}

type Twit struct {
	ID           string `json:"id"`
	UserID       string `json:"user_id"`
	Texts        string `json:"texts"`
	ReadersCount int    `json:"readers_count"`
}
