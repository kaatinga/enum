package enum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncode(t *testing.T) {
	tests := []struct {
		input string
		err   assert.ErrorAssertionFunc
	}{
		{"_", assert.NoError},
		{"a", assert.NoError},
		{"bbb", assert.NoError},
		{"1", assert.NoError},
		{"-", assert.Error},
		{"aaaaaaaaaaaa", assert.Error},
		{"aaaaaaaaaa", assert.NoError},
		{"aaaaaaaaa_", assert.NoError},
		{"aaaaa aaaa", assert.NoError},
		{"AAAA aaa_", assert.NoError},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got, err := Encode(tt.input)
			tt.err(t, err)

			if err == nil && got.String() != tt.input {
				t.Errorf("Encode() got = %v, want %v", got, tt.input)
			}
		})
	}
}
