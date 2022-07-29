package common

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GetLine() string {
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
	return string(buf)
}

func GetIntSlice() []int {
	slice := make([]int, 0)
	line := GetLine()
	lines := strings.Split(line, " ")
	for _, v := range lines {
		slice = append(slice, s2i(v))
	}
	return slice
}

func s2i(s string) int {
	v, ok := strconv.Atoi(s)
	if ok != nil {
		panic("Faild : " + s + " can't convert to int")
	}
	return v
}
