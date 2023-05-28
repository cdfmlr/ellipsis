// Package ellipsis provides functions to ellipsis strings.
package ellipsis

import (
	"strings"
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

const utf8CharMaxSize = 4

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

	maxLen := n * utf8CharMaxSize
	if maxLen >= len(s) {
		maxLen = len(s)
	}

	var runes []rune
	if leftToRight {
		runes = []rune(s[:maxLen])
		if len(runes) > n {
			runes = runes[:n]
		}
	} else {
		runes = []rune(s[len(s)-maxLen:])
		if len(runes) > n {
			runes = runes[len(runes)-n:]
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
