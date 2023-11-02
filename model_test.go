package enum

import (
	"errors"
	"testing"
)

func TestEncode(t *testing.T) {
	tests := []struct {
		input string
		err   error
	}{
		{"_", nil},
		{"a", nil},
		{"bbb", nil},
		{"1", nil},
		{"-", invalidCharacter('-')},
		{"aaaaaaaaaaaa", ErrInvalidLength},
		{"aaaaaaaaaa", nil},
		{"aaaaaaaaa_", nil},
		{"aaaaa aaaa", nil},
		{"AAAA aaa_", nil},
		{"", ErrEmptyString},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got, err := Encode(tt.input)
			if !errors.Is(err, tt.err) {
				t.Errorf("Encode() error = %v, wantErr %v", err, tt.err)
			}

			if err == nil && got.String() != tt.input {
				t.Errorf("Encode() got = %v, want %v", got, tt.input)
			}
		})
	}
}
