// Package ellipsis provides functions to ellipsis strings.
package ellipsis

import (
	"strings"
	"unicode/utf8"
)

// EllipsisFunc is a function type for Centering, Starting and Ending
type EllipsisFunc func(string, int) string

// Centering ellipsis a long string s -> "front...end"
func Centering(s string, n int) string {
	if n <= 3 {
		return "..."
	}

	if len(s) <= n {
		return s
	}

	n -= 3
	h := n / 2

	var sb strings.Builder
	sb.WriteString(cutString(s, h, cutLeftToRight))
	sb.WriteString("...")
	sb.WriteString(cutString(s, h, cutRightToLeft))

	return sb.String()
}

type cutDirection bool

const (
	cutLeftToRight cutDirection = true
	cutRightToLeft cutDirection = false
)

// cutString cuts a string s into a string of n utf-8 runes.
func cutString(s string, n int, leftToRight cutDirection) string {
	if n <= 0 {
		return ""
	}

	if n >= len(s) {
		return s
	}

	// var runes = make([]rune, 0, n)
	var runes = make([]rune, n)

	var idx int = 0
	if !leftToRight {
		idx = len(s)
	}

	for i := 0; i < n; i++ {
		if leftToRight {
			char, size := utf8.DecodeRuneInString(s[idx:])
			// runes = append(runes, char)
			runes[i] = char
			idx += size
		} else {
			char, size := utf8.DecodeLastRuneInString(s[:idx])
			// runes = append([]rune{char}, runes...)
			runes[n-i-1] = char
			idx -= size
		}
	}

	return string(runes)
}

// Starting ellipsis a long string s -> "...end"
func Starting(s string, n int) string {
	if n <= 3 {
		return "..."
	}

	if len(s) <= n {
		return s
	}

	n -= 3

	var sb strings.Builder
	sb.WriteString("...")
	sb.WriteString(cutString(s, n, cutRightToLeft))

	return sb.String()
}

// Ending ellipsis a long string s -> "front..."
func Ending(s string, n int) string {
	if n <= 3 {
		return "..."
	}

	if len(s) <= n {
		return s
	}

	n -= 3

	var sb strings.Builder
	sb.WriteString(cutString(s, n, cutLeftToRight))
	sb.WriteString("...")

	return sb.String()
}

// noEllipsis returns the string s, without ellipsis.
// It is used for benchmarking.
func noEllipsis(s string, n int) string {
	return s
}
