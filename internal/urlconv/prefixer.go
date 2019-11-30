package urlconv

import (
	"math/rand"
	"time"
)

type Prefixer interface {
	CreateIDPrefix() (int, error)
}

type Randomizer struct {
	min int64
	max int64
}

func NewRandomizer(min int64, max int64) *Randomizer {
	return &Randomizer{min, max}
}

func (r *Randomizer) CreateIDPrefix() (int, error) {
	rand.Seed(time.Now().UnixNano())

	return rand.Intn(int(r.max-r.min+1)) + int(r.min), nil
}
