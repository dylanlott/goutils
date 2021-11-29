package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("hello Dylan")

	// values.
	fmt.Println("go" + "lang")

	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0) // float example float= decimal number.

	fmt.Println(true && false)
	fmt.Println(true || false) //*I could not fiqure out how to do the two straight lines and had to copy and paste
	fmt.Println(!true)
	//? I could not understand line 17-19 the code work as it did in the example but i did not grasp what I did.
	// found some more examples online so i did these to help build my understanding of boolean principles.

	var num1 = 10
	var num2 = 20

	var res1 = num1 == num2 // equal
	var res2 = num1 != num2 // not eaqual
	var res3 = num1 > num2  // greater than
	var res4 = num1 < num2  // less than
	var res5 = num1 <= num2 // less then equal to
	var res6 = num1 >= num2 // greater than equal to

	fmt.Println(res1, res2, res3, res4, res5, res6)

	//Variables
	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "apples"
	fmt.Println(f)

	// Constant
	const s string = "constant"

	fmt.Println(s)

	const n = 500000000

	const d1 = 3e20 / n
	fmt.Println(d1)

	fmt.Println(int64(d1))

	fmt.Println(math.Sin(n))
	fmt.Println("branch merge test")
}
