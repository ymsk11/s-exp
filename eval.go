package main

func fold(fn func(any, any) any, accum any, cons *Cons) any {
	if cons == nil {
		return accum
	}

	switch car := cons.Car.(type) {
	case *Cons:
		var cdr *Cons
		if cons.Cdr == nil {
			cdr = nil
		} else {
			cdr = cons.Cdr.(*Cons)
		}
		return fold(fn, fn(accum, eval(car)), cdr)
	default:
		var cdr *Cons
		if cons.Cdr == nil {
			cdr = nil
		} else {
			cdr = cons.Cdr.(*Cons)
		}
		return fold(fn, fn(accum, car), cdr)
	}
}

func plus(a any, b any) any {
	return a.(int) + b.(int)
}

func prod(a any, b any) any {
	return a.(int) * b.(int)
}

func eval(any any) any {
	cons, ok := any.(*Cons)
	if ok == false {
		return any
	}
	op := cons.Car.(string)
	cdr := cons.Cdr.(*Cons)
	switch op {
	case "+":
		return fold(plus, 0, cdr)
	case "*":
		return fold(prod, 1, cdr)
	case "list":
		return cdr
	case "car":
		return eval(cdr.Car).(*Cons).Car
	case "cdr":
		return eval(cdr.Car).(*Cons).Cdr
	case "cons":
		return &Cons{
			eval(cdr.Car),
			eval(cdr.Cdr.(*Cons).Car),
		}
	case "atom":
		if cdr.Car == nil {
			return nil
		}
		switch cdr.Car.(type) {
		case *Cons:
			return nil
		default:
			return &T{}
		}
	case "eq":
		a := eval(cdr.Car)
		b := eval(cdr.Cdr.(*Cons).Car)
		if a == b {
			return &T{}
		}
		return nil
	case "if":
		cond := eval(cdr.Car)
		t := cdr.Cdr.(*Cons).Car
		f := cdr.Cdr.(*Cons).Cdr.(*Cons).Car
		switch cond.(type) {
		case *T:
			return eval(t)
		default:
			return eval(f)
		}
	case "quote":
		return cdr.Car
	default:
		return nil
	}
}
