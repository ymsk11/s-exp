package main

func fold(fn func(any, any) any, accum any, cons *Cons) any {
	if cons == nil {
		return accum
	}

	switch car := cons.Car.(type) {
	case *Cons:
		return fold(fn, fn(accum, eval(car)), cons.Cdr)
	default:
		return fold(fn, fn(accum, car), cons.Cdr)
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
		return fold(plus, 0, cons.Cdr)
	case "*":
		return fold(prod, 1, cons.Cdr)
	default:
		return nil
	}
}
