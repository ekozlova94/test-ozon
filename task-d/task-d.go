package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := new(big.Rat)
	b := new(big.Rat)
	c := new(big.Rat)
	_, _ = fmt.Scan(a, b)
	c.Add(a, b)
	fmt.Print(c.RatString())
}
