package main

import (
	"fmt"
)

func main() {
	sexps := []string{
		"(+ (* 10 10) 41)",
		"(list 1 2 3 4 5 )",
		"(car (list 1 2 3 4 5 ))",
		"(cdr (list 1 2 3 4 5 ))",
		"(car (cdr (list 1 2 3 )))",
		"(* (car (list 3 4 )) (car (cdr (list 3 4 ))))",
		"(cons 1 2 )",
		"(cons (car (list 1 2 )) (cdr (list 1 2 )))",
		"(atom nil)",
		"(atom (list 1 2 3))",
		"(atom 3)",
		"(eq (* 2 3) (* 3 2))",
		"(if (atom 2) (* 1 100) (hoge))",
		"(if nil (hoge) (car (list 1 2 3)))",
		"(quote (1 2 3 4 5))",
		"(quote 1)",
		"(quote (* 1 100))",
		"(nth 3 (quote ( 0 1 2 3 4 5 )))",
		"((lambda (x) (* x x)) 10)",
		"((lambda (x y) (* x x (+ y y))) 3 2)",
		"(if t 1 2)",
		"(define x (* 10 10))",
		"(* x x)",
		"(define f (lambda (x) (* x x)))",
		"(f 10)",
		"(define second (lambda (x) (car (cdr x))))",
		"(second (list 1 2 3 4 5))",
	}

	env := &Env{
		map[string]any{},
	}
	for i, text := range sexps {
		fmt.Println(i, "> ", text)
		parsed := ParseText(text)
		print(parsed)
		print(eval(parsed, env))
		fmt.Println("")
	}
}
