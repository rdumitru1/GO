package main

import "fmt"

// Notice that in Go, the break statement is not required at the end of a case to stop it from falling through to the next case. The break statement is implicit in Go.

func getCreator(os string) string {
	var creator string
	switch os {
	case "linux":
		creator = "Linus Torvalds"
	case "windows":
		creator = "Bill Gates"
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
