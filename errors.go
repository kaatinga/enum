package enum

import "errors"

type invalidInput byte

const ErrInvalidLength invalidInput = iota

func (i invalidInput) Error() string {
	switch i {
	case ErrInvalidLength:
		return "input string is too long, the maximum is 10 characters"
	default:
		return "unknown error"
	}
}

func (i invalidInput) Is(target error) bool {
	var ii invalidInput
	ok := errors.As(target, &ii)
	if !ok {
		return false
	}

	return i == ii
}

type invalidCharacter rune

const (
	ErrInvalidCharacter invalidCharacter = iota
)

func (err invalidCharacter) Error() string {
	return "input string contains an incorrect character '" + string(err) + "'"
}

func (err invalidCharacter) Is(target error) bool {
	var ic invalidCharacter
	return errors.As(target, &ic)
}
