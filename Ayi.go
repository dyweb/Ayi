package main

import "fmt"

func weekend() (string, string) {
	return  "Saturday", "Sunday"
}

func main() {
	var Ayi = "xyp"
	x := 1 + 1
	const c = "I am constant"
	fmt.Println(Ayi)
	fmt.Println("1 + "  + "1 = ", x)

	for i := 1; i <= 3; i++ {
		fmt.Println(i)
	}

	if 1 < 2 {
		fmt.Println("1 < 2")
	}

	// array have fixed length
	var arr [2]int
	for j := 0; j < len(arr); j++ {
		fmt.Println(arr[j])
	}

	// slice
	var s []int
	s = make([]int ,2)
	s = append(s,3,2)

	for k := 0; k < len(s); k++ {
		fmt.Println(s[k])
	}

	// map
	var m = make(map[string]string)

	m["ayi"] = "xyp"

	sat, sun := weekend()
	fmt.Println(sat,sun)
}