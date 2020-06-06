package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	//generateInput()
	input, _ := os.OpenFile("input.txt", os.O_RDONLY|os.O_CREATE, os.ModePerm)
	output, _ := os.OpenFile("output.txt", os.O_WRONLY|os.O_CREATE, os.ModePerm)
	//noinspection GoUnhandledErrorResult
	defer input.Close()
	//noinspection GoUnhandledErrorResult
	defer output.Close()

	r := bufio.NewReader(input)
	t, _ := r.ReadString('\n')
	t = strings.TrimSpace(t)
	target, _ := strconv.Atoi(t)
	sequence := make(map[int]bool)
	for {
		st, err := r.ReadString(' ')
		st = strings.TrimSpace(st)
		key, errParse := strconv.Atoi(st)
		if errParse != nil {
			if err == io.EOF {
				break
			}
			continue
		}
		if _, ok := sequence[(target - key)]; ok {
			if _, err := fmt.Fprintln(output, 1); err != nil {
				fmt.Println(err)
			}
			return
		}
		sequence[key] = true
	}
	if _, err := fmt.Fprintln(output, 0); err != nil {
		fmt.Println(err)
	}
	return
}

func generateInput() {
	input, _ := os.OpenFile("input.txt", os.O_WRONLY|os.O_CREATE, os.ModePerm)
	_, _ = fmt.Fprintln(input, 999999999)
	for i := 1; i < 5000000; i++ {
		_, _ = fmt.Fprint(input, i)
		_, _ = fmt.Fprint(input, " ")
	}
	_, _ = fmt.Fprint(input, '\n')
	_ = input.Close()
}
