package ptr

import "fmt"

func init() {
	var a int
	fmt.Println(&a)
	var p *int = &a
	*p = 1650
	fmt.Println(a)
}
