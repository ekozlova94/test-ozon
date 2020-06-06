package main

import (
	"fmt"
	"os"
)

func Merge2Channels(f func(int) int, in1 <-chan int, in2 <-chan int, out chan<- int, n int) {
	var c0 = make(chan []int, n)
	var c1 = make(chan []int, n)

	go readData(in1, in2, n, c0)
	go calcF(f, n, c0, c1)
	go writeData(out, n, c1)
}

func readData(in1, in2 <-chan int, n int, downstream chan []int) {
	for i := 0; i < n; i++ {
		x1 := <-in1
		x2 := <-in2
		downstream <- []int{x1, x2, i}
	}
	close(downstream)
}

func calcF(f func(int) int, n int, upstream chan []int, downstream chan []int) {
	results := make(chan []int)
	for i := 0; i < n; i++ {
		go func() {
			for numbers := range upstream {
				y1 := f(numbers[0])
				y2 := f(numbers[1])
				y := y1 + y2
				results <- []int{y, numbers[2]}
			}
		}()
	}
	counter := 0
	for r := range results {
		counter++
		downstream <- r
		if counter == n {
			close(results)
		}
	}
	close(downstream)
}

func writeData(out chan<- int, n int, upstream chan []int) {
	results := make([]int, n)
	for answer := range upstream {
		results[answer[1]] = answer[0]
	}
	output, err := os.OpenFile("см формат вывода", os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	for _, r := range results {
		out <- r
		_, _ = fmt.Fprintln(output, r)
	}
	_ = output.Close()
}
