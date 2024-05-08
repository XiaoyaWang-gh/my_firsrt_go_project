package str

import "fmt"

func traversalString() {
	s := "行步至春深"
	for i := 0; i < len(s); i++ {
		fmt.Printf("%v(%c)", s[i], s[i])
	}
	fmt.Println()
	for _, r := range s {
		fmt.Printf("%v(%c)", r, r)
	}
	fmt.Println()
}

func changeString() {
	s := "操前曲而后晓声"
	fmt.Println("before change:", s)
	// 强制转化成[]rune
	rune_s := []rune(s)
	rune_s[0] = '练'
	rune_s[1] = '千'
	rune_s[2] = '剑'
	rune_s[3] = '而'
	rune_s[4] = '后'
	rune_s[5] = '识'
	rune_s[6] = '器'
	fmt.Println("after change:", string(rune_s))
}

func init() {
	traversalString()
	changeString()
}
