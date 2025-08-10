package enum

import (
	"testing"
)

type testCase[E invalidInput | invalidCharacter] struct {
	err    E
	target error
	want   bool
}

func Test_Is(t *testing.T) {
	tests := []testCase[invalidInput]{
		{ErrInvalidLength, ErrInvalidLength, true},
		{ErrInvalidLength, ErrInvalidCharacter, false},
		{ErrInvalidLength, nil, false},
	}

	tests2 := []testCase[invalidCharacter]{
		{invalidCharacter(33), ErrInvalidCharacter, true},
		{invalidCharacter(33), nil, false},
	}

	for _, tt := range tests {
		target := "/<nil>"
		if tt.target != nil {
			target = "/" + tt.target.Error()
		}
		t.Run(tt.err.Error()+target, func(t *testing.T) {
			if tt.want != tt.err.Is(tt.target) {
				t.Errorf("Is() = %v, want %v", tt.err.Is(tt.target), tt.want)
			}
		})
	}

	for _, tt := range tests2 {
		target := "/<nil>"
		if tt.target != nil {
			target = "/" + tt.target.Error()
		}
		t.Run(tt.err.Error()+"/"+target, func(t *testing.T) {
			if tt.want != tt.err.Is(tt.target) {
				t.Errorf("Is() = %v, want %v", tt.err.Is(tt.target), tt.want)
			}
		})
	}
}

func Test_invalidCharacter_Error(t *testing.T) {
	tests := []struct {
		err  invalidCharacter
		want string
	}{
		{invalidCharacter(33), "input string contains an incorrect character '!'"},
		{invalidCharacter(0), "input string contains an incorrect character '\x00'"},
		{invalidCharacter(127), "input string contains an incorrect character '\x7f'"},
	}
	for _, tt := range tests {
		t.Run(tt.err.Error(), func(t *testing.T) {
			if tt.want != tt.err.Error() {
				t.Errorf("Error() = %v, want %v", tt.err.Error(), tt.want)
			}
		})
	}
}
