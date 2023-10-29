package c

import "fmt"

func Output(val int, two bool) {
	fmt.Printf("Output")
	if two {
		fmt.Printf(" for part 2")
	}
	fmt.Printf(": %d\n", val)
}
