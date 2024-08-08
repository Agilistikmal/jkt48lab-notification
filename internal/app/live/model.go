package live

import "time"

type Live struct {
	Slug         string     `json:"slug,omitempty"`
	Title        string     `json:"title,omitempty"`
	ImageUrl     string     `json:"image_url,omitempty"`
	Views        int        `json:"views,omitempty"`
	Platform     string     `json:"platform,omitempty"`
	OriginalUrl  string     `json:"original_url,omitempty"`
	StreamingUrl string     `json:"streaming_url,omitempty"`
	RoomID       string     `json:"room_id,omitempty"`
	StartedAt    *time.Time `json:"started_at,omitempty"`
	EndAt        *time.Time `json:"end_at,omitempty"`
	Member       Member     `json:"member,omitempty"`
}

type Member struct {
	Username  string `json:"username,omitempty"`
	Name      string `json:"name,omitempty"`
	ImageUrl  string `json:"image_url,omitempty"`
	Followers int    `json:"followers,omitempty"`
}
