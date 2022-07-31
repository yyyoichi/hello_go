package abc261

import (
	"fmt"
	// "bufio"
	// "os"
	// "strconv"
)

func b() {
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
