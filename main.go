package main

import (
	"learngo/mocking"
	"os"
)

func main() {
	s := &mocking.DefaultSleeper{}
	
	mocking.Countdown(os.Stdout, s)

}