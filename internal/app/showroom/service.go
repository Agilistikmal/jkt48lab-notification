package showroom

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/agilistikmal/jkt48lab-notification/internal/app/live"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) GetLives() ([]*live.Live, error) {
	var lives []*live.Live
	resp, err := http.Get(os.Getenv("SHOWROOM_BASE_URL"))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var showroomResponses ShowroomResponses
	err = json.Unmarshal(body, &showroomResponses)
	if err != nil {
		return nil, err
	}

	for _, liveGenres := range showroomResponses.OnLives {
		if liveGenres.GenreName == "idol" {
			for _, l := range liveGenres.Lives {
				if !strings.Contains(strings.ToUpper(l.RoomUrlKey), strings.ToUpper(os.Getenv("PREFIX"))) {
					continue
				}
				streamingUrl, err := s.GetStreamingUrl(l.RoomID)
				if err != nil {
					return nil, err
				}
				startedAt := time.Unix(int64(l.StartedAt), 0)
				lives = append(lives, &live.Live{
					Slug:         l.RoomUrlKey,
					Title:        l.Telop,
					ImageUrl:     l.ImageSquare,
					Views:        l.ViewNum,
					Platform:     "SHOWROOM",
					OriginalUrl:  fmt.Sprintf("https://showroom-live.com/r/%v", l.RoomUrlKey),
					StreamingUrl: streamingUrl.Url,
					StartedAt:    &startedAt,
					EndAt:        nil,
					Member: live.Member{
						Username:  l.RoomUrlKey,
						Name:      l.MainName,
						ImageUrl:  l.ImageSquare,
						Followers: l.FollowerNum,
					},
					RoomID: fmt.Sprintf("%v", l.RoomID),
				})
			}
			break
		}
	}
	return lives, nil
}

func (s *Service) GetStreamingUrl(roomID int) (*StreamingUrl, error) {
	resp, err := http.Get(fmt.Sprintf("%v?room_id=%v", os.Getenv("SHOWROOM_STREAMING_URL_BASE_URL"), roomID))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var streamingUrlResponses ShowroomStreamingUrlResponses
	err = json.Unmarshal(body, &streamingUrlResponses)
	if err != nil {
		return nil, err
	}

	var url *StreamingUrl
	for _, streamingUrl := range streamingUrlResponses.StreamingUrlList {
		if streamingUrl.Label == "original quality" {
			url = &streamingUrl
			break
		}
	}

	if url == nil {
		return nil, fmt.Errorf("error get streaming url")
	}

	return url, nil
}
