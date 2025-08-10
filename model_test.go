package enum

import (
	"encoding/json"
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

func TestEnum_JSON_MarshalUnmarshal(t *testing.T) {
	type S struct {
		E Enum `json:"e"`
	}

	original := S{E: MustEncode("Hello")}
	data, err := json.Marshal(original)
	if err != nil {
		t.Fatalf("json.Marshal failed: %v", err)
	}

	// At this point, data contains: {"e":"Hello"}
	// The Enum is marshaled as a string, not a number.

	var decoded S
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("json.Unmarshal failed: %v", err)
	}

	if decoded.E.String() != "Hello" {
		t.Errorf("expected %q, got %q", "Hello", decoded.E.String())
	}
}
