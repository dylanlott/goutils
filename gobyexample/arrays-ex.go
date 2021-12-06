package main

import "fmt"

func main() {

	var a [5]int
	fmt.Println("emp:", a)

	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	fmt.Println("len:", len(a))

	q := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", q)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	fmt.Println()

	var b [14]string
	b[0] = "f"
	b[1] = "u"
	b[2] = "c"
	b[3] = "k"
	b[4] = "i"
	b[5] = "n"
	b[6] = "g"
	b[7] = " "
	b[8] = "d"
	b[9] = "i"
	b[10] = "d"
	b[11] = " "
	b[12] = "i"
	b[13] = "t"
	fmt.Println(b)
	//arrays examples went really well and I feel like i understood it. test passed in playground.
}
