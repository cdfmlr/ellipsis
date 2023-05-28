// Package ellipsis provides functions to ellipsis strings.
package ellipsis

import "strings"

// Centering ellipsis a long s -> "front...end"
func Centering(s string, n int) string {
	if n <= 3 {
		return "..."
	}

	r := []rune(s)

	if len(r) <= n {
		return s
	}

	n -= 3
	h := n / 2

	var sb strings.Builder
	sb.WriteString(string(r[:h]))
	sb.WriteString("...")
	sb.WriteString(string(r[len(r)-h:]))

	return sb.String()
}

// Starting ellipsis a long s -> "...end"
func Starting(s string, n int) string {
	if n <= 3 {
		return "..."
	}

	r := []rune(s)

	if len(r) <= n {
		return s
	}

	n -= 3

	var sb strings.Builder
	sb.WriteString("...")
	sb.WriteString(string(r[len(r)-n:]))

	return sb.String()
}

// Ending ellipsis a long s -> "front..."
func Ending(s string, n int) string {
	if n <= 3 {
		return "..."
	}
	
	r := []rune(s)

	if len(r) <= n {
		return s
	}

	n -= 3

	var sb strings.Builder
	sb.WriteString(string(r[:n]))
	sb.WriteString("...")

	return sb.String()
}
