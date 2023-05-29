package ellipsis

import (
	"fmt"
)

func ExampleCentering() {
	// Centering ellipsis a long string s -> "begin...end"
	s := Centering("1234567890", 7)
	fmt.Println(s)
	// Output: 12...90
}

func ExampleStarting() {
	// Starting ellipsis a long string s -> "...end"
	s := Starting("1234567890", 7)
	fmt.Println(s)
	// Output: ...7890
}

func ExampleEnding() {
	// Ending ellipsis a long string s -> "begin..."
	s := Ending("1234567890", 7)
	fmt.Println(s)
	// Output: 1234...
}

