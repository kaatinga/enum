package enum

import (
	"github.com/stretchr/testify/assert"
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
		{ErrInvalidLength, ErrEmptyString, false},
		{ErrInvalidLength, ErrInvalidCharacter, false},
		{ErrInvalidLength, nil, false},
		{ErrEmptyString, ErrInvalidLength, false},
		{ErrEmptyString, ErrEmptyString, true},
		{ErrEmptyString, ErrInvalidCharacter, false},
		{ErrEmptyString, nil, false},
	}

	tests2 := []testCase[invalidCharacter]{
		{invalidCharacter(33), ErrInvalidCharacter, true},
		{invalidCharacter(33), ErrEmptyString, false},
		{invalidCharacter(33), nil, false},
	}

	for _, tt := range tests {
		target := "/<nil>"
		if tt.target != nil {
			target = "/" + tt.target.Error()
		}
		t.Run(tt.err.Error()+target, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.err.Is(tt.target), "Is(%v)", tt.target)
		})
	}

	for _, tt := range tests2 {
		target := "/<nil>"
		if tt.target != nil {
			target = "/" + tt.target.Error()
		}
		t.Run(tt.err.Error()+"/"+target, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.err.Is(tt.target), "Is(%v)", tt.target)
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
			assert.Equalf(t, tt.want, tt.err.Error(), "Error()")
		})
	}
}
