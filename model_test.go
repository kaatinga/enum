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
		{"bbbe", nil},
		{"1", nil},
		{"-", invalidCharacter('-')},
		{"aa1aaaaaaaaa", ErrInvalidLength},
		{"", nil},
		{"aaaaaaaaba", nil},
		{"aaaaaaaaa_", nil},
		{"aaaaa aaaa", nil},
		{"AAAA aaa_", nil},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got, err := Encode(tt.input)
			if !errors.Is(err, tt.err) {
				t.Errorf("Encode() error = %v, wantErr %v", err, tt.err)
				return
			}
			gotString := got.String()
			if err == nil && gotString != tt.input {
				t.Errorf("Encode() got = %s, want %v", gotString, tt.input)
			}
		})
	}
}
