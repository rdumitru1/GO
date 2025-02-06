package main

import "fmt"

func main() {
	// initialize variables here
	var smsSendingLimit int
	var contPerSMS float64
	var hasPermission bool
	var username string

	fmt.Printf(
		"%v %f %v %q\n",
		smsSendingLimit,
		contPerSMS,
		hasPermission,
		username,
	)
}
