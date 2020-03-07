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

type ConfigurableSleepr struct {
	durationTime time.Duration
	sleep func(time.Duration)
}

func (cs *ConfigurableSleepr) Sleep()  {
	cs.sleep(cs.durationTime)
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
	t := 1 * time.Second
	s := &ConfigurableSleepr{t, time.Sleep}
	CountDown(os.Stdout, s)
}
