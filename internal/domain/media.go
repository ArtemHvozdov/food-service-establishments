package domain

import "time"

type Article struct {
	Name string
	URL  string
	Date time.Time
}

type YoutubeVideo struct {
	Name string
	URL  string
	Date time.Time
}

type Photo struct {
    Path      string
    SourceURL string
    AddedAt   time.Time
}

func (p Photo) IsZero() bool { return p.Path == "" }