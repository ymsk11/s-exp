package main

import (
	"fmt"
)

func main() {

	s := &Cons{
		"+",
		&Cons{
			41,
			&Cons{
				&Cons{
					"*",
					&Cons{
						10,
						&Cons{
							10,
							nil,
						},
					},
				},
				nil,
			},
		},
	}
	fmt.Println(s.Str())
	fmt.Println(eval(s))

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
	}
	for i, text := range sexps {
		fmt.Println(i, "> ", text)
		parsed := ParseText(text)
		print(parsed)
		print(eval(parsed))
		fmt.Println("")
	}
}
