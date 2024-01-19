package enum

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

func Encode(s string) (Enum, error) {
	if len(s) == 0 {
		return 0, ErrEmptyString
	}
	if len(s) > 10 {
		return 0, ErrInvalidLength
	}

	var result Enum
	for _, char := range s {
		result <<= 6
		switch {
		case char >= 0x61 && char <= 0x7A:
			result += Enum(char - lowerShift)
		case char >= 0x30 && char <= 0x39:
			result += Enum(char - digitShift)
		case char >= 0x41 && char <= 0x5A:
			result += Enum(char - upperShift)
		case char == 95: // underscore
			result += Enum(36)
		case char == 32: // space
			result += Enum(63)
		default:
			return 0, invalidCharacter(char)
		}
	}

	return result, nil
}

const firstCharMask = 0b111111

func (e Enum) String() string {
	var result string
	for e > 0 {
		result = string(chars[e&firstCharMask]) + result
		e >>= 6
	}
	return result
}
