package main

import (
	"fmt"
)

func main() {
	wait := make(chan struct{})

	seq := make(chan int)
	Processor(seq, wait)
	for num := 2; num < 10; num++ {
		seq <- num
	}
	close(seq)

	<-wait
}

func Processor(seq chan int, wait chan struct{}) {
	go func() {

		prime, ok := <-seq
		if !ok {
			close(wait)
			return
		}
		fmt.Println(prime)

		out := make(chan int)
		Processor(out, wait)
		for num := range seq {
			if num%prime != 0 {
				out <- num
			}
		}
		close(out)

	}()
}
