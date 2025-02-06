package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	const name = "boots"
	const bear = "🐻"
	fmt.Printf("constant 'name' byte length: %d\n", len(name))
	fmt.Printf("constant 'name' rune length: %d\n", utf8.RuneCountInString(name))
	fmt.Println("=====================================")
	fmt.Printf("Hi %s, so good to have you back in the arcanum\n", name)

	fmt.Printf("constant 'bear' byte length: %d\n", len(bear))
	fmt.Printf("constant 'bear' rune length: %d\n", utf8.RuneCountInString(bear))
	fmt.Println("=====================================")
	fmt.Printf("Hi %s, so good to have you back in the arcanum\n", bear)

}
