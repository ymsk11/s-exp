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

func eval(cons *Cons) any {
	op := cons.Car.(string)
	switch op {
	case "+":
		cdr := cons.Cdr.(*Cons)
		return fold(plus, 0, cdr)
	case "*":
		cdr := cons.Cdr.(*Cons)
		return fold(prod, 1, cdr)
	default:
		return nil
	}
}
