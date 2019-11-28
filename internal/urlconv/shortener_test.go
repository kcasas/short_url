package urlconv_test

import (
	"fmt"
	"testing"

	"github.com/kcasas/short_url/internal/mocks"
	"github.com/kcasas/short_url/internal/urlconv"
	"github.com/stretchr/testify/mock"
)

func TestShortener(t *testing.T) {
	storage := new(mocks.Storage)
	storage.On("Store", mock.Anything, mock.Anything).Return(int64(1), nil)
	shortener, _ := urlconv.New(storage)
	shorten, _ := shortener.Shorten("gago", -1)

	fmt.Println(shorten)
}
