package array

import "fmt"

var str = [5]string{3: "hello world", 4: "tom"}

func printArr(arr *[5]int) {
	arr[0] = 10
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func init() {
	for index, s := range str {
		fmt.Printf("%d : %s ", index, s)
	}
	fmt.Println()

	struct_a := [...]struct {
		name string
		age  uint8
	}{
		{"Sheldon", 31}, // 可省略元素类型。
		{"Amy", 29},     // 别忘了最后一行的逗号。
	}

	fmt.Println(struct_a)

	d := [3][4]int{{1, 2, 3, 4}, {1, 2, 3, 4}, {1, 2, 3, 4}}

	for k1, v1 := range d {
		for k2, v2 := range v1 {
			fmt.Printf("(%d,%d)=%d ", k1, k2, v2)
		}
		fmt.Println()
	}

	var arr1 [5]int
	printArr(&arr1)
	fmt.Println(arr1)
	arr2 := [...]int{2, 4, 6, 8, 10}
	printArr(&arr2)
	fmt.Println(arr2)
}
