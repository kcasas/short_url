package urlconv

import "time"

type Storage interface {
	InsertLongURL(longurl string, exp time.Duration) (id int64, err error)
	GetLongURL(shorturl string) (longurl string, err error)
}
