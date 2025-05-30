// Package kebabcase - Super-Fast snake-case implementation.
package kebabcase

const separatorByte = '-'

// Kebabcase the given string.
func Kebabcase(s string) string { //nolint:gocognit
	idx := 0
	hasLower := false
	hasUnderscore := false
	lowercaseSinceUnderscore := false

	// loop through all good characters:
	// - lowercase
	// - digit
	// - underscore (as long as the next character is lowercase or digit)
	for ; idx < len(s); idx++ {
		if isLower(s[idx]) {
			hasLower = true
			if hasUnderscore {
				lowercaseSinceUnderscore = true
			}
			continue
		} else if isDigit(s[idx]) {
			continue
		} else if s[idx] == separatorByte && idx > 0 && idx < len(s)-1 && (isLower(s[idx+1]) || isDigit(s[idx+1])) {
			hasUnderscore = true
			lowercaseSinceUnderscore = false
			continue
		}
		break
	}

	if idx == len(s) {
		return s // no changes needed, can just borrow the string
	}

	// if we get here then we must need to manipulate the string
	b := make([]byte, 0, 64)
	b = append(b, s[:idx]...)

	if isUpper(s[idx]) && (!hasLower || hasUnderscore && !lowercaseSinceUnderscore) {
		for idx < len(s) && (isUpper(s[idx]) || isDigit(s[idx])) {
			b = append(b, asciiLowercaseArray[s[idx]])
			idx++
		}

		for idx < len(s) && (isLower(s[idx]) || isDigit(s[idx])) {
			b = append(b, s[idx])
			idx++
		}
	}

	for idx < len(s) {
		if !isAlphanumeric(s[idx]) {
			idx++
			continue
		}

		if len(b) > 0 {
			b = append(b, separatorByte)
		}

		for idx < len(s) && (isUpper(s[idx]) || isDigit(s[idx])) {
			b = append(b, asciiLowercaseArray[s[idx]])
			idx++
		}

		for idx < len(s) && (isLower(s[idx]) || isDigit(s[idx])) {
			b = append(b, s[idx])
			idx++
		}
	}
	return string(b) // return manipulated string
}

func isAlphanumeric(c byte) bool {
	return isLower(c) || isUpper(c) || isDigit(c)
}

func isUpper(c byte) bool {
	return c >= 'A' && c <= 'Z'
}

func isLower(c byte) bool {
	return c >= 'a' && c <= 'z'
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

var asciiLowercaseArray = [256]byte{
	0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07,
	0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f,
	0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17,
	0x18, 0x19, 0x1a, 0x1b, 0x1c, 0x1d, 0x1e, 0x1f,
	' ', '!', '"', '#', '$', '%', '&', '\'',
	'(', ')', '*', '+', ',', '-', '.', '/',
	'0', '1', '2', '3', '4', '5', '6', '7',
	'8', '9', ':', ';', '<', '=', '>', '?',
	'@',

	'a', 'b', 'c', 'd', 'e', 'f', 'g',
	'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o',
	'p', 'q', 'r', 's', 't', 'u', 'v', 'w',
	'x', 'y', 'z',

	'[', '\\', ']', '^', '_',
	'`', 'a', 'b', 'c', 'd', 'e', 'f', 'g',
	'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o',
	'p', 'q', 'r', 's', 't', 'u', 'v', 'w',
	'x', 'y', 'z', '{', '|', '}', '~', 0x7f,
	0x80, 0x81, 0x82, 0x83, 0x84, 0x85, 0x86, 0x87,
	0x88, 0x89, 0x8a, 0x8b, 0x8c, 0x8d, 0x8e, 0x8f,
	0x90, 0x91, 0x92, 0x93, 0x94, 0x95, 0x96, 0x97,
	0x98, 0x99, 0x9a, 0x9b, 0x9c, 0x9d, 0x9e, 0x9f,
	0xa0, 0xa1, 0xa2, 0xa3, 0xa4, 0xa5, 0xa6, 0xa7,
	0xa8, 0xa9, 0xaa, 0xab, 0xac, 0xad, 0xae, 0xaf,
	0xb0, 0xb1, 0xb2, 0xb3, 0xb4, 0xb5, 0xb6, 0xb7,
	0xb8, 0xb9, 0xba, 0xbb, 0xbc, 0xbd, 0xbe, 0xbf,
	0xc0, 0xc1, 0xc2, 0xc3, 0xc4, 0xc5, 0xc6, 0xc7,
	0xc8, 0xc9, 0xca, 0xcb, 0xcc, 0xcd, 0xce, 0xcf,
	0xd0, 0xd1, 0xd2, 0xd3, 0xd4, 0xd5, 0xd6, 0xd7,
	0xd8, 0xd9, 0xda, 0xdb, 0xdc, 0xdd, 0xde, 0xdf,
	0xe0, 0xe1, 0xe2, 0xe3, 0xe4, 0xe5, 0xe6, 0xe7,
	0xe8, 0xe9, 0xea, 0xeb, 0xec, 0xed, 0xee, 0xef,
	0xf0, 0xf1, 0xf2, 0xf3, 0xf4, 0xf5, 0xf6, 0xf7,
	0xf8, 0xf9, 0xfa, 0xfb, 0xfc, 0xfd, 0xfe, 0xff,
}
