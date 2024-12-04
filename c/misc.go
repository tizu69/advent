package c

import "fmt"

func Output(val int, two bool) {
	if two {
		fmt.Printf("Output 2: %d\n", val)
	} else {
		fmt.Printf("Output  : %d\n", val)
	}
}

func If[T any](cond bool, then, otherwise T) T {
	if cond {
		return then
	}
	return otherwise
}
