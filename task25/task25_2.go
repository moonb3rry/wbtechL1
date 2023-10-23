package main

import (
	"fmt"
	"time"
)

func mySleep2(duration time.Duration) {
	done := make(chan bool)
	time.AfterFunc(duration, func() {
		done <- true
	})
	<-done
}

func main() {
	fmt.Println("Sleeping for 2 seconds...")
	mySleep2(2 * time.Second)
	fmt.Println("Awake!")
}
