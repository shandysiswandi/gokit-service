package uuid

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValidUUID(t *testing.T) {
	testTable := []struct {
		title string
		input string
		want  bool
	}{
		{title: "", input: "", want: false},
	}

	for _, tc := range testTable {
		t.Run(tc.title, func(t *testing.T) {
			act := IsValidUUID(tc.input)
			assert.Equal(t, tc.want, act)
		})
	}
}
