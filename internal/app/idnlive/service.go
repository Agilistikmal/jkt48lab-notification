package idnlive

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
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

func (s *Service) GetIDNLives() ([]*live.Live, error) {
	var lives []*live.Live

	page := 1
	for {
		query, _ := json.Marshal(map[string]any{
			"query": fmt.Sprintf(`
				query GetLivestreams {
					getLivestreams(page: %v) {
						slug
						title
						image_url
						view_count
						playback_url
						status
						live_at
						gift_icon_url
						creator {
								username
								name
								follower_count
						}
					}
				}
				`, page),
		})

		gReq, err := http.NewRequest("POST", "https://api.idn.app/graphql", bytes.NewBuffer(query))
		if err != nil {
			return nil, err
		}
		gReq.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(gReq)
		if err != nil {
			log.Println(err.Error())
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		var idnResponses IDNResponses
		err = json.Unmarshal(body, &idnResponses)
		if err != nil {
			return nil, err
		}

		if idnResponses.Data == nil && idnResponses.Errors != nil {
			return nil, fmt.Errorf("idn response error")
		}

		if len(idnResponses.Data.GetLivestreams) == 0 {
			break
		}

		for _, l := range idnResponses.Data.GetLivestreams {
			if l.Status != "live" {
				continue
			}
			if !strings.Contains(strings.ToUpper(l.Creator.Username), strings.ToUpper(os.Getenv("PREFIX"))) {
				continue
			}
			startedAt, _ := time.Parse("2006-01-02T15:04:05+07:00", l.LiveAt)
			endAt, _ := time.Parse("2006-01-02T15:04:05+07:00", l.EndAt)
			lives = append(lives, &live.Live{
				Slug:         l.Slug,
				Platform:     "IDN",
				Title:        l.Title,
				OriginalUrl:  fmt.Sprintf("https://www.idn.app/%v/live/%v", l.Creator.Username, l.Slug),
				RoomID:       l.Slug,
				StreamingUrl: l.PlaybackUrl,
				Views:        l.ViewCount,
				ImageUrl:     l.ImageUrl,
				StartedAt:    &startedAt,
				EndAt:        &endAt,
				Member: live.Member{
					Username:  l.Creator.Username,
					Name:      l.Creator.Name,
					ImageUrl:  l.Creator.Avatar,
					Followers: l.Creator.FollowerCount,
				},
			})
		}
		page++
	}
	return lives, nil
}
