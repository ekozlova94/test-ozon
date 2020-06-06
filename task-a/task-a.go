package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strings"
)

func main() {
	//generateInput()
	input, _ := os.OpenFile("input-201.txt", os.O_RDONLY|os.O_CREATE, os.ModePerm)
	output, _ := os.OpenFile("input-201.a.txt", os.O_WRONLY|os.O_CREATE, os.ModePerm)
	//noinspection GoUnhandledErrorResult
	defer input.Close()
	//noinspection GoUnhandledErrorResult
	defer output.Close()
	r := bufio.NewReader(input)
	values := make(map[string]int64)
	for {
		key, err := r.ReadString('\n')
		key = strings.TrimSpace(key)
		if err == io.EOF {
			break
		}
		if _, ok := values[key]; ok {
			values[key]++
		} else {
			values[key] = 1
		}
	}
	for k, v := range values {
		if v%2 != 0 {
			if _, err := fmt.Fprintln(output, k); err != nil {
				fmt.Println(err)
			}
			return
		}
	}
}

func generateInput() {
	input, _ := os.OpenFile("input.txt", os.O_WRONLY|os.O_CREATE, os.ModePerm)
	for i := 1; i < 1000000; i++ {
		a := rand.Intn(100)
		_, _ = fmt.Fprintln(input, a)
		_, _ = fmt.Fprintln(input, a)
	}
	_, _ = fmt.Fprintln(input, rand.Intn(100))
	_ = input.Close()
}
