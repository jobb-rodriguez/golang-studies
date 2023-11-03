package main

import "fmt"

func main() {
	// initialize variables here
	var smsSendingLimit int = 100
	var costPerSMS float64 = 0.2
	var hasPermission bool = true
	var username string = "John Smith"

	fmt.Printf("%v %f %v %q\n", smsSendingLimit, costPerSMS, hasPermission, username)
}