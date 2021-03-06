package main

import (
	"fmt"
	"time"
)

func main() {

	i := 2
	fmt.Print("write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("its the weekend")
	default:
		fmt.Println("its the weekday")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("ist before noon")
	default:
		fmt.Println("its after noon")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm a int")
		default:
			fmt.Printf("Don't know what type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
	// Passed test in playground, dont really understand what "Switch" is doing.
	// Originally messed us by not capitalizing Println as well as did a / instead of a \ and it took me forever to find haha.

}
