package stru

import "fmt"

type my_process struct {
	t int
	c float64
}

func init() {
	var p1 my_process
	p1.t = 5
	p1.c = 10.18
	fmt.Printf("p1=%#v\n", p1)

	var me struct {
		name string
		age  int
	}

	me.name = "禾八"
	me.age = 24
	fmt.Printf("me=%#v\n", me)
}
