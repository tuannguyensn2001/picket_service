package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseTime(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		layout := "HH:MM:SS DD/MM/YYYY"
		val := "10:42:23 23/11/2022"

		output, err := ParseTime(layout, val)
		assert.NotNil(t, output)
		assert.Nil(t, err)
		assert.Equal(t, 10, output.Hour())
		assert.Equal(t, 42, output.Minute())
		assert.Equal(t, 23, output.Day())

	})
}
