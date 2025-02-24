package blog

import "time"

type BlogPost struct {
	Title                string
	Excerpt              string
	Date                 time.Time
	ReadingTimeInSeconds int
}
