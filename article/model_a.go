package article

import "time"

type Article struct {
	Id            string
	Title         string
	Description   string
	Summary       string
	TimePublished time.Time
	Author        string
}
