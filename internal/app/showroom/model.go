package showroom

type ShowroomResponses struct {
	OnLives []struct {
		GenreName string `json:"genre_name,omitempty"`
		Lives     []struct {
			RoomUrlKey       string `json:"room_url_key,omitempty"`
			Telop            string `json:"telop,omitempty"`
			FollowerNum      int    `json:"follower_num,omitempty"`
			StartedAt        int    `json:"started_at,omitempty"`
			ImageSquare      string `json:"image_square,omitempty"`
			ViewNum          int    `json:"view_num,omitempty"`
			MainName         string `json:"main_name,omitempty"`
			PremiumRoomType  int    `json:"premium_room_type,omitempty"`
			RoomID           int    `json:"room_id,omitempty"`
			StreamingUrlList []struct {
				Url string `json:"url,omitempty"`
			} `json:"streaming_url_list,omitempty"`
		} `json:"lives,omitempty"`
	} `json:"onlives,omitempty"`
}

type ShowroomStreamingUrlResponses struct {
	StreamingUrlList []StreamingUrl `json:"streaming_url_list,omitempty"`
}

type StreamingUrl struct {
	Label string `json:"label,omitempty"`
	Url   string `json:"url,omitempty"`
}
