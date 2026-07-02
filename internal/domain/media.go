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
