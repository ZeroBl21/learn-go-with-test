package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

// Sleeper allows you to put delays.
type Sleeper interface {
	Sleep()
}

// DefaultSleeper is an implementation of Sleeper with a predefined delay.
type DefaultSleeper struct{}

// Sleep will puase execution for the defined Duration.
func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

const countdownStart = 3
const finalWord = "Go!"

// Countdown prints a countdown from 3 to out with a delay between count provided by Sleeper
func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
    sleeper.Sleep()
	}

	fmt.Fprint(out, finalWord)
}

func main() {
	sleeper := &DefaultSleeper{}
	Countdown(os.Stdout, sleeper)
}
