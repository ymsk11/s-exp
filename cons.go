package main

import "fmt"

type any interface{}

// Cons cell
type Cons struct {
	Car any
	Cdr any
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