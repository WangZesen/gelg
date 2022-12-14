package template

var TUtils string = `package log

import (
	"unicode/utf8"
)

// borrowed from Zerolog: https://github.com/rs/zerolog/blob/c2b9d0e2defd04e1afb9712132a66681d077eb42/internal/json/string.go
var noEscapeTable = [256]bool{}
const hex = "0123456789abcdef"

func appendString(buf []byte, data string) []byte {
	// Loop through each character in the string.
	for i := 0; i < len(data); i++ {
		// Check if the character needs encoding. Control characters, slashes,
		// and the double quote need json encoding. Bytes above the ascii
		// boundary needs utf8 encoding.
		if !noEscapeTable[data[i]] {
			// We encountered a character that needs to be encoded. Switch
			// to complex version of the algorithm.
			return appendStringComplex(buf, data, i)
		}
	}
	// The string has no need for encoding and therefore is directly
	// appended to the byte slice.
	return append(buf, data...)
}

//// appendStringComplex is used by appendString to take over an in
// progress JSON string encoding that encountered a character that needs
// to be encoded.
func appendStringComplex(dst []byte, s string, i int) []byte {
	start := 0
	for i < len(s) {
		b := s[i]
		if b >= utf8.RuneSelf {
			r, size := utf8.DecodeRuneInString(s[i:])
			if r == utf8.RuneError && size == 1 {
				// In case of error, first append previous simple characters to
				// the byte slice if any and append a replacement character code
				// in place of the invalid sequence.
				if start < i {
					dst = append(dst, s[start:i]...)
				}
				// Changed by Zesen: don't include the error if encoutered
				// dst = append(dst, ` + "`\ufffd`" + `...)
				i += size
				start = i
				continue
			}
			i += size
			continue
		}
		if noEscapeTable[b] {
			i++
			continue
		}
		// We encountered a character that needs to be encoded.
		// Let's append the previous simple characters to the byte slice
		// and switch our operation to read and encode the remainder
		// characters byte-by-byte.
		if start < i {
			dst = append(dst, s[start:i]...)
		}
		switch b {
		case '"', '\\':
			dst = append(dst, '\\', b)
		case '\b':
			dst = append(dst, '\\', 'b')
		case '\f':
			dst = append(dst, '\\', 'f')
		case '\n':
			dst = append(dst, '\\', 'n')
		case '\r':
			dst = append(dst, '\\', 'r')
		case '\t':
			dst = append(dst, '\\', 't')
		default:
			dst = append(dst, '\\', 'u', '0', '0', hex[b>>4], hex[b&0xF])
		}
		i++
		start = i
	}
	if start < len(s) {
		dst = append(dst, s[start:]...)
	}
	return dst
}`
