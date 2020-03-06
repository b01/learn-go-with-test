package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct {
	
}

func (s *DefaultSleeper) Sleep ()  {
	time.Sleep(1 * time.Second)
}

func CountDown(w io.Writer, s Sleeper)  {
	for i := 3; i > 0; i-- {
		s.Sleep()
		fmt.Fprintln(w, i)
	}

	s.Sleep()
	fmt.Fprint(w, "Go!")
}

func main()  {
	s := &DefaultSleeper{}
	CountDown(os.Stdout, s)
}
