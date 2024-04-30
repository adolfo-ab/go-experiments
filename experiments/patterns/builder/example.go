package main

import "fmt"

func main() {
	// Create builder
	var nb = newNotificationBuilder()

	// Set properties
	nb.SetTitle("Example")
	nb.SetSubtitle("Example")
	nb.SetMessage("Example")
	nb.SetImage("Example")
	nb.SetIcon("Example")
	nb.SetPriority(1)
	nb.SetNotType("Example")

	// Build
	if notif, err := nb.Build(); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Notification: %+v", notif)
	}

}
