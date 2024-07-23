package main

import (
	"fmt"
	"time"
)

type Sleep struct {
	duration time.Duration
}

func (s *Sleep) Run() {
	time.Sleep(s.duration)
}

func NewSleep(duration time.Duration) Sleep {
	return Sleep{duration}
}

func main() {
	s := NewSleep(time.Second * 5)
	s.Run()
	fmt.Println("Проснулись")
}
