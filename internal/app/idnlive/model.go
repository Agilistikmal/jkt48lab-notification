package idnlive

type IDNResponses struct {
	Data *struct {
		GetLivestreams []struct {
			Slug           string `json:"slug,omitempty"`
			Title          string `json:"title,omitempty"`
			ImageUrl       string `json:"image_url,omitempty"`
			ViewCount      int    `json:"view_count,omitempty"`
			PlaybackUrl    string `json:"playback_url,omitempty"`
			RoomIdentifier string `json:"room_identifier,omitempty"`
			Status         string `json:"status,omitempty"`
			LiveAt         string `json:"live_at,omitempty"`
			EndAt          string `json:"end_at,omitempty"`
			ScheduledAt    string `json:"scheduled_at,omitempty"`
			Category       struct {
				Name string `json:"name,omitempty"`
				Slug string `json:"slug,omitempty"`
			} `json:"category,omitempty"`
			Creator struct {
				Username       string `json:"username,omitempty"`
				Name           string `json:"name,omitempty"`
				Avatar         string `json:"avatar,omitempty"`
				BioDescription string `json:"bio_description,omitempty"`
				FollowerCount  int    `json:"follower_count,omitempty"`
			} `json:"creator,omitempty"`
		} `json:"getLivestreams,omitempty"`
	} `json:"data,omitempty"`
	Errors []struct {
		Message string   `json:"message,omitempty"`
		Path    []string `json:"path,omitempty"`
	} `json:"errors,omitempty"`
}
