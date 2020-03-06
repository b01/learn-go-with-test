package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "Go!"
const countStart = 3

type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct {}

func (s *DefaultSleeper) Sleep ()  {
	time.Sleep(1 * time.Second)
}

func CountDown(w io.Writer, s Sleeper)  {
	for i := countStart; i > 0; i-- {
		s.Sleep()
		fmt.Fprintln(w, i)
	}

	s.Sleep()
	fmt.Fprint(w, finalWord)
}

func main()  {
	s := &DefaultSleeper{}
	CountDown(os.Stdout, s)
}
