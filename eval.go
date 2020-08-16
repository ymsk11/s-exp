package main

func fold(fn func(any, any) any, accum any, cons *Cons, env *Env) any {
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
		return fold(fn, fn(accum, eval(car, env)), cdr, env)
	case string: // is variable
		value := env.get(car)
		var cdr *Cons
		if cons.Cdr == nil {
			cdr = nil
		} else {
			cdr = cons.Cdr.(*Cons)
		}
		return fold(fn, fn(accum, value), cdr, env)
	default:
		var cdr *Cons
		if cons.Cdr == nil {
			cdr = nil
		} else {
			cdr = cons.Cdr.(*Cons)
		}
		return fold(fn, fn(accum, car), cdr, env)
	}
}

func plus(a any, b any) any {
	return a.(int) + b.(int)
}

func prod(a any, b any) any {
	return a.(int) * b.(int)
}

func eval(any any, env *Env) any {
	cons, ok := any.(*Cons)
	if ok == false {
		return any
	}
	cdr := cons.Cdr.(*Cons)
	switch car := cons.Car.(type) {
	case string:
		return evalOperator(car, cdr, env)
	case *Cons:
		return evalCarCons(car, cdr, env)
	default:
		return nil
	}
}

func evalCarCons(car *Cons, cdr *Cons, env *Env) any {
	switch op := eval(car, env).(type) {
	case *function:
		return op.eval(cdr, env)
	default:
		return nil
	}
}

func evalOperator(op string, cdr *Cons, env *Env) any {
	switch op {
	case "+":
		return fold(plus, 0, cdr, env)
	case "*":
		return fold(prod, 1, cdr, env)
	case "%":
		a := cdr.Car.(int)
		b := cdr.Cdr.(*Cons).Car.(int)
		return a % b
	case "list":
		return cdr
	case "car":
		return eval(cdr.Car, env).(*Cons).Car
	case "cdr":
		return eval(cdr.Car, env).(*Cons).Cdr
	case "cons":
		return &Cons{
			eval(cdr.Car, env),
			eval(cdr.Cdr.(*Cons).Car, env),
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
		a := eval(cdr.Car, env)
		b := eval(cdr.Cdr.(*Cons).Car, env)
		if a == b {
			return &T{}
		}
		return nil
	case "if":
		cond := eval(cdr.Car, env)
		t := cdr.Cdr.(*Cons).Car
		f := cdr.Cdr.(*Cons).Cdr.(*Cons).Car
		switch cond.(type) {
		case *T:
			return eval(t, env)
		default:
			return eval(f, env)
		}
	case "quote":
		return cdr.Car
	case "nth":
		n := eval(cdr.Car, env).(int)
		target := eval(cdr.Cdr.(*Cons).Car, env).(*Cons)
		return eval(target.nth(n), env)
	case "lambda":
		args := cdr.Car.(*Cons)
		fn := cdr.Cdr.(*Cons).Car.(*Cons)
		return &function{
			args,
			fn,
		}
	case "define":
		name := cdr.Car.(string)
		value := eval(cdr.Cdr.(*Cons).Car, env)
		env.set(name, value)
		return nil
	default:
		value := env.get(op)
		switch op := value.(type) {
		case *function:
			return op.eval(cdr, env)
		}
		return nil
	}
}
