package ellipsis

import (
	"fmt"
)

func ExampleCentering() {
	// Centering ellipsis a long string s -> "begin...end"
	s := Centering("0123456789零一二三四五六七八九", 7)
	fmt.Println(s)
	// Output: 01...八九
}

func ExampleStarting() {
	// Starting ellipsis a long string s -> "...end"
	s := Starting("0123456789零一二三四五六七八九", 7)
	fmt.Println(s)
	// Output: ...六七八九
}

func ExampleEnding() {
	// Ending ellipsis a long string s -> "begin..."
	s := Ending("0123456789零一二三四五六七八九", 7)
	fmt.Println(s)
	// Output: 0123...
}

