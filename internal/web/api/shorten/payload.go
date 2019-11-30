package shorten

import (
	"fmt"
	"net/url"
)

type RequestPayload struct {
	URL        string `json:"url"`
	Expiration int64  `json:"expiration"`
}

func (payload *RequestPayload) validate() error {
	_, err := url.ParseRequestURI(payload.URL)
	if err != nil {
		return fmt.Errorf("invalid url: %w", err)
	}

	return nil
}
