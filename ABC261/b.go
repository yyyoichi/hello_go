package abc261

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextLine() string {
	sc.Scan()
	return sc.Text()
}
func s2i(s string) int {
	v, ok := strconv.Atoi(s)
	if ok != nil {
		panic("Faild : " + s + " can't convert to int")
	}
	return v
}
func main() {
	l1 := nextLine()
	n := s2i(l1)
	board := make([]string, n)
	for i := 0; i < n; i++ {
		board[i] = nextLine()
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if i == j {
				continue
			}
			a := board[i][j : j+1]
			b := board[j][i : i+1]
			if a == "W" && b == "L" {
				continue
			}
			if a == "L" && b == "W" {
				continue
			}
			if a == "D" && b == "D" {
				continue
			}
			fmt.Println("incorrect")
			return
		}
	}
	fmt.Println("correct")
}
