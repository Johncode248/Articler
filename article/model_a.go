package article

import "time"

type Article struct {
	Id            string
	title         string
	description   string
	summary       string
	timePublished time.Time
}
