package model

import "time"

// Twit
type CreateTwitRequest struct {
	UserID       string `json:"user_id"`
	PublisherFIO string `json:"publisher_fio"`
	Type         string `json:"type"`
	Texts        string `json:"texts"`
	Title        string `json:"title"`
}

type Twit struct {
	ID           string    `json:"id"`
	UserID       string    `json:"user_id"`
	PublisherFIO string    `json:"publisher_fio"`
	Type         string    `json:"type"`
	Texts        string    `json:"texts"`
	Title        string    `json:"title"`
	ReadersCount int       `json:"readers_count"`
	CreatedAt    time.Time `json:"created_at"`
}

// Video
type CreateVideoRequest struct {
	TwitID string `json:"twit_id"`
	Video  string `json:"video"`
}

type Video struct {
	ID     string `json:"id"`
	TwitID string `json:"twit_id"`
	Video  string `json:"video"`
}

// Photo
type CreatePhotoRequest struct {
	TwitID string `json:"twit_id"`
	Photo  string `json:"photo"`
}

type Photo struct {
	ID     string `json:"id"`
	TwitID string `json:"twit_id"`
	Photo  string `json:"photo"`
}

// Music
type CreateMusicRequest struct {
	TwitID string `json:"twit_id"`
	MP3    string `json:"mp3"`
}

type Music struct {
	ID     string `json:"id"`
	TwitID string `json:"twit_id"`
	MP3    string `json:"mp3"`
}

// Location
type CreateLocationRequest struct {
	TwitID string `json:"twit_id"`
	Lat    string `json:"lat"`
	Lon    string `json:"lon"`
}

type Location struct {
	ID     string `json:"id"`
	TwitID string `json:"twit_id"`
	Lat    string `json:"lat"`
	Lon    string `json:"lon"`
}

// URL
type CreateURLRequest struct {
	TwitID string `json:"twit_id"`
	URL    string `json:"url"`
}

type URL struct {
	ID     string `json:"id"`
	TwitID string `json:"twit_id"`
	URL    string `json:"url"`
}

// User
type User struct {
	ID           string `json:"id"`
	Login        string `json:"login"`
	PasswordHash string `json:"-"`
	Role         string `json:"role"`
}
