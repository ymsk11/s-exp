package main

type function struct {
	Args *Cons
	Body *Cons
}

func (fn *function) eval(args *Cons) any {
	if fn.Args.len() != args.len() {
		panic("invalid arguments")
	}
	body := fn.Body
	for i := 0; i < args.len(); i++ {
		body = walk(replace(fn.Args.nth(i), args.nth(i)), body)
	}
	return eval(body)
}
