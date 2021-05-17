package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetDate(t *testing.T) {
	date := GetDate()
	assert.NotEqual(t, 0, len(date))
}
