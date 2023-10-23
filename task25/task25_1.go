package main

import (
	"fmt"
	"time"
)

func mySleep(duration time.Duration) {
	<-time.After(duration)
}

func main() {
	fmt.Println("Sleeping for 2 seconds...")
	mySleep(2 * time.Second)
	fmt.Println("Awake!")
}
