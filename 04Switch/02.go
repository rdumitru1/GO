package main

import "fmt"

// If you do want a case to fall through to the next case, you can use the fallthrough keyword.
// The default case does what you’d expect: it’s the case that runs if none of the other cases match.

func getCreator(os string) string {
	var creator string
	switch os {
	case "linux":
		creator = "Linus Torvalds"
	case "windows":
		creator = "Bill Gates"

	// all three of these cases will set creator to "A Steve"
	case "macOS":
		fallthrough
	case "Mac OS X":
		fallthrough
	case "mac":
		creator = "A Steve"

	default:
		creator = "Unknown"
	}
	return creator
}

func main() {
	os := "linux"
	fmt.Println("The creator is:", getCreator(os))
}
