package main

import (
	"fmt"
	"time"
)

func sleep(value time.Duration) {
	<-time.After(value)
}

func main() {
	sleep(3 * time.Second)
	fmt.Println("Wake up")
}
