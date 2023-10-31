package enum

import "fmt"

var chars = []int32{
	'0',
	'1',
	'2',
	'3',
	'4',
	'5',
	'6',
	'7',
	'8',
	'9',
	'A',
	'B',
	'C',
	'D',
	'E',
	'F',
	'G',
	'H',
	'I',
	'J',
	'K',
	'L',
	'M',
	'N',
	'O',
	'P',
	'Q',
	'R',
	'S',
	'T',
	'U',
	'V',
	'W',
	'X',
	'Y',
	'Z',
	'_',
	'a',
	'b',
	'c',
	'd',
	'e',
	'f',
	'g',
	'h',
	'i',
	'j',
	'k',
	'l',
	'm',
	'n',
	'o',
	'p',
	'q',
	'r',
	's',
	't',
	'u',
	'v',
	'w',
	'x',
	'y',
	'z',
}

const (
	digitShift rune = 48
	upperShift rune = 55
	lowerShift rune = 60
)

type Enum int64

var (
	errInvalidCharacter = fmt.Errorf("invalid character")
	errEmptyString      = fmt.Errorf("empty string")
	errInvalidLength    = fmt.Errorf("maximum string length reached")
)

func Encode(s string) (Enum, error) {
	switch len(s) {
	case 0:
		return 0, errEmptyString
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10:
	default:
		return 0, errInvalidLength
	}

	var result Enum
	for _, char := range s {
		fmt.Println(char)
		result <<= 6
		switch {
		case char >= 0x30 && char <= 0x39:
			result += Enum(char - digitShift)
		case char >= 0x41 && char <= 0x5A:
			result += Enum(char - upperShift)
		case char == 95: // underscore
			result += Enum(char - 59)
		case char >= 0x61 && char <= 0x7A:
			result += Enum(char - lowerShift)
		default:
			return 0, fmt.Errorf("%w: character, '%d', is not a valid character", errInvalidCharacter, char)
		}
	}

	return result, nil
}

const firstCharMask = 0b111111

func (e Enum) String() string {
	var result string
	for e > 0 {
		fmt.Println(e & firstCharMask)
		result = string(chars[e&firstCharMask]) + result
		e >>= 6
	}
	return result
}
