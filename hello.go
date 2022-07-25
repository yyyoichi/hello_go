package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	p := parameter()
	fmt.Println(p.getIntSlice())
	solve()
}
func solve() {
}

type Param struct {
	getLine     func() string
	getIntSlice func() []int
}

func parameter() Param {
	const BUFSIZE = 10000000
	rdr := bufio.NewReaderSize(os.Stdin, BUFSIZE)
	buf := make([]byte, 0, 16)
	for {
		l, p, e := rdr.ReadLine()
		if e != nil {
			fmt.Println(e.Error())
			panic(e)
		}
		buf = append(buf, l...)
		if !p {
			break
		}
	}
	line := string(buf)
	tc := typeConv()
	var param Param
	param.getLine = func() string {
		return line
	}
	param.getIntSlice = func() []int {
		slice := make([]int, 0)
		lines := strings.Split(line, " ")
		for _, v := range lines {
			slice = append(slice, tc.s2i(v))
		}
		return slice
	}
	return param
}

type TPCV struct {
	s2i func(string) int
}

func typeConv() TPCV {
	var tc TPCV
	tc.s2i = func(s string) int {
		v, ok := strconv.Atoi(s)
		if ok != nil {
			panic("Faild : " + s + " can't convert to int")
		}
		return v
	}
	return tc
}
