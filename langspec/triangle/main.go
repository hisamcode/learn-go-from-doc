package triangle

import "fmt"

// Triangle1
func Triangle1() {
	for i := 0; i < 5; i++ {
		a := "*"
		for z := 0; z < i; z++ {
			a = a + "*"
		}
		fmt.Println(a)
	}
}

// extend
func extend(slice []int, element int) []int {
	n := len(slice)
	slice = slice[0 : n+1]
	slice[n] = element
	return slice

}

// Triangle2
func Triangle2() {
	var iBuffer [10]int
	slice := iBuffer[0:0]
	fmt.Println(len(slice), len(iBuffer))
	for i := 0; i < len(iBuffer); i++ {
		slice = extend(slice, i)
		fmt.Println(slice)
	}
}

// Triangle3
func Triangle3() {
	cap := [10]any{}
	triangle := cap[0:0]
	for i := 0; i < 9; i++ {
		n := len(triangle)
		triangle = triangle[0 : n+1]
		triangle[n] = "*"
		fmt.Println(triangle...)
	}
}
