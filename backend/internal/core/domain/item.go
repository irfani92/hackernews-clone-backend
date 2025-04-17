package domain

import "time"

type ItemType string

const (
	TypeStory   ItemType = "story"
	TypeComment ItemType = "comment"
	TypeJob     ItemType = "job"
	TypePoll    ItemType = "poll"
	TypePollOpt ItemType = "pollopt"
)

type Item struct {
	ID          int      `json:"id"`
	Type        ItemType `json:"type"`
	By          string   `json:"by"`
	Time        int64    `json:"time"`
	Text        *string  `json:"text"`
	URL         *string  `json:"url"`
	Title       *string  `json:"title"`
	Parent      *int     `json:"parent"`
	Poll        *int     `json:"poll"`
	Kids        []int    `json:"kids"`
	Dead        bool     `json:"dead"`
	Deleted     bool     `json:"deleted"`
	Score       int      `json:"score"`
	Parts       []int    `json:"parts"`
	Descendants int      `json:"descendants"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
