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
	m := make(map[string]int)
	stringN := nextLine()
	n := s2i(stringN)
	for i := 0; i < n; i++ {
		s := nextLine()
		v, ok := m[s]
		var num int
		if ok {
			num = v + 1
			m[s] = num
			fmt.Printf("%v(%v)\n", s, num)
		} else {
			m[s] = 0
			fmt.Println(s)
		}
	}
	solve()
}
func solve() {}
