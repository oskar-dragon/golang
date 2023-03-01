package mocking

import (
	"fmt"
	"io"
	"time"
)

const finalWorld = "Go!"
const countdownStart = 3

type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func Countdown(out io.Writer, sleeper Sleeper) {
	
	for i := countdownStart; i > 0;i-- {
		fmt.Fprint(out, i)
		sleeper.Sleep()
	}

	fmt.Fprint(out, finalWorld)
}