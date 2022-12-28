package main

import (
	"fmt"
)

func str(any any) string {
	if any == nil {
		return "nil"
	}
	switch a := any.(type) {
	case *Cons:
		return a.Str()
	case *T:
		return "T"
	default:
		return fmt.Sprint(a)
	}
}

func print(any any) {
	fmt.Println(str(any))
}
