package abc261

import (
	"example/hello/common"
	"fmt"
)

func a() {
	ns := common.GetIntSlice()
	l1 := ns[0]
	r1 := ns[1]
	l2 := ns[2]
	r2 := ns[3]
	l := l1
	if l2 > l1 {
		l = l2
	}
	r := r1
	if r2 < r1 {
		r = r2
	}
	ans := r - l
	if ans < 0 {
		ans = 0
	}
	fmt.Println(ans)
}
