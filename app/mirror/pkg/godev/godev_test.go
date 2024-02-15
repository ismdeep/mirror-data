package godev

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetDownloadLinks(t *testing.T) {
	links, err := GetDownloadLinks()
	assert.NoError(t, err)
	t.Logf("links size: %v", len(links))
}
