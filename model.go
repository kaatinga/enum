package enum

import (
	"bytes"
)

var chars = [64]rune{
	0:  '0',
	1:  '1',
	2:  '2',
	3:  '3',
	4:  '4',
	5:  '5',
	6:  '6',
	7:  '7',
	8:  '8',
	9:  '9',
	10: 'A',
	11: 'B',
	12: 'C',
	13: 'D',
	14: 'E',
	15: 'F',
	16: 'G',
	17: 'H',
	18: 'I',
	19: 'J',
	20: 'K',
	21: 'L',
	22: 'M',
	23: 'N',
	24: 'O',
	25: 'P',
	26: 'Q',
	27: 'R',
	28: 'S',
	29: 'T',
	30: 'U',
	31: 'V',
	32: 'W',
	33: 'X',
	34: 'Y',
	35: 'Z',
	36: '_',
	37: 'a',
	38: 'b',
	39: 'c',
	40: 'd',
	41: 'e',
	42: 'f',
	43: 'g',
	44: 'h',
	45: 'i',
	46: 'j',
	47: 'k',
	48: 'l',
	49: 'm',
	50: 'n',
	51: 'o',
	52: 'p',
	53: 'q',
	54: 'r',
	55: 's',
	56: 't',
	57: 'u',
	58: 'v',
	59: 'w',
	60: 'x',
	61: 'y',
	62: 'z',
	63: ' ',
}

const (
	digitShift rune = 48
	upperShift rune = 55
	lowerShift rune = 60
)

type Enum int64

// MarshalText and UnmarshalText are used by encoding/json for string encoding.
// This causes Enum to be marshaled as its original string value in JSON, not as a number.
func (e Enum) MarshalText() ([]byte, error) {
	return e.Bytes(), nil
}

func (e *Enum) UnmarshalText(text []byte) (err error) {
	name := string(text)
	*e, err = Encode(name)
	return
}

// Encode encodes a string into an Enum.
func Encode(s string) (Enum, error) {
	if len(s) > 10 {
		return 0, ErrInvalidLength
	}

	var result uint64
	for _, char := range s {
		result <<= 6
		switch {
		case char >= 0x61 && char <= 0x7A:
			result += uint64(char - lowerShift)
		case char >= 0x30 && char <= 0x39:
			result += uint64(char - digitShift)
		case char >= 0x41 && char <= 0x5A:
			result += uint64(char - upperShift)
		case char == 95: // underscore
			result += 36
		case char == 32: // space
			result += 63
		default:
			return 0, invalidCharacter(char)
		}
	}

	result |= uint64(len(s)) << 60 // Store the length of the string in the top 4 bits

	return Enum(result), nil
}

// MustEncode is like Encode but panics if Encode returns an error.
func MustEncode(s string) Enum {
	result, err := Encode(s)
	if err != nil {
		panic(err)
	}
	return result
}

const firstCharMask = 0b111111

// String returns the string representation of the Enum.
func (e Enum) String() string {
	var builder bytes.Buffer
	topBit := int64(((uint64(e) >> 60) - 1) * 6)
	builder.Grow(int(topBit))
	for ; topBit >= 0; topBit -= 6 {
		builder.WriteRune(chars[(uint64(e)>>topBit)&firstCharMask])
	}
	return builder.String()
}

// Bytes returns the []byte representation of the Enum.
func (e Enum) Bytes() []byte {
	var builder bytes.Buffer
	topBit := int64(((uint64(e) >> 60) - 1) * 6)
	builder.Grow(int(topBit))
	for ; topBit >= 0; topBit -= 6 {
		builder.WriteRune(chars[(uint64(e)>>topBit)&firstCharMask])
	}
	return builder.Bytes()
}
