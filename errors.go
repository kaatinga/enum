package enum

import "errors"

type invalidInput byte

const (
	ErrEmptyString invalidInput = iota
	ErrInvalidLength
)

func (err invalidInput) Error() string {
	switch err {
	case ErrEmptyString:
		return "input string is empty"
	case ErrInvalidLength:
		return "input string is too long, the maximum is 10 characters"
	default:
		return "unknown error"
	}
}

func (err invalidInput) Is(target error) bool {
	var ii invalidInput
	ok := errors.As(target, &ii)
	if !ok {
		return false
	}

	return err == ii
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
