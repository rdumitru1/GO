package main

import "fmt"

func main() {
	var r rune = 'A'

	// In Go, a string is just a read-only slice of bytes ([]byte). It can store any kind of data, including text, binary data, or arbitrary byte sequences.
	// Here, "Hello" is stored as bytes [72 101 108 108 111] (ASCII values).
	// This means Go does not store strings as characters, but as raw bytes.
	s1 := "Hello"
	fmt.Println([]byte(s1), "\n") // Output: [72 101 108 108 111]

	//A rune is just an alias for int32. It is used to represent a single Unicode code point.
	// 'A' is stored as 65, the Unicode code point for 'A'.
	// Since a rune is int32, it can store any Unicode character (even non-ASCII ones).
	fmt.Println(r, "\n") // Output: 65

	// ASCII characters fit in 1 byte (uint8 or byte).
	// Unicode characters can take up to 4 bytes.
	// A rune (int32) is large enough (4 bytes) to store any Unicode character.
	s := "Hello, 世界" // "世界" means "World" in Chinese
	for _, r := range s {
		fmt.Printf("Char: %c, Unicode: %U\n", r, r)
	}
}
