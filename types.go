package main

import "fmt"

type any interface{}

// T is T
type T struct{}

// Cons cell
type Cons struct {
	Car any
	Cdr any
}

func print(any any) {
	switch a := any.(type) {
	case *Cons:
		fmt.Println(a.Str())
	case *T:
		fmt.Println("T")
	default:
		fmt.Println(a)
	}
}

func (cons *Cons) nth(n int) any {
	p := cons
	for i := 0; i < n; i++ {
		p = p.Cdr.(*Cons)
	}
	return p.Car
}

// Str to string
func (cons *Cons) Str() string {

	var carText string
	switch car := cons.Car.(type) {
	case *Cons:
		carText = car.Str()
	default:
		carText = fmt.Sprint(car)
	}
	var cdrText string
	if cons.Cdr == nil {
		cdrText = "nil"
	} else {
		switch cdr := cons.Cdr.(type) {
		case *Cons:
			cdrText = cdr.Str()
		default:
			cdrText = fmt.Sprint(cdr)
		}
	}

	return "(" + carText + " . " + cdrText + ")"
}
