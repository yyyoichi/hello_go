package main

import (
	"fmt"
)

func main() {
	p := Parameter()
	ns := p.getIntSlice()
	fmt.Println(ns[0] + ns[1])
	fmt.Println(ns[0] * ns[1])
	solve()
}
func solve() {
}
