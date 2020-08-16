package main

import "testing"

func TestMain(t *testing.T) {
	cases := []struct {
		before string
		text   string
		result string
	}{
		{
			text:   "(+ (* 4 10) 1)",
			result: "41",
		},
		{
			text:   "(list 1 2 3 4 5)",
			result: "(1 . (2 . (3 . (4 . (5 . nil)))))",
		},
		{
			text:   "(car (cdr (list 1 2 3 4 5)))",
			result: "2",
		},
		{
			text:   "(cons 1 2)",
			result: "(1 . 2)",
		},
		{
			text:   "(cons (* 1 2) (+ 1 3))",
			result: "(2 . 4)",
		},
		{
			text:   "(atom t)",
			result: "T",
		},
		{
			text:   "(atom (list 1 2))",
			result: "nil",
		},
		{
			text:   "(eq (* 2 3) (+ 1 2 3))",
			result: "T",
		},
		{
			text:   "(if T (* 1 2) (* 3 4))",
			result: "2",
		},
		{
			text:   "(if nil (* 1 2) (* 3 4))",
			result: "12",
		},
		{
			text:   "(quote (1 2 3))",
			result: "(1 . (2 . (3 . nil)))",
		},
		{
			text:   "(quote (* 1 10))",
			result: "(* . (1 . (10 . nil)))",
		},
		{
			text:   "(nth 3 (list 0 1 2 3 4 5))",
			result: "3",
		},
		{
			text:   "((lambda (x y) ( * x y )) 3 2)",
			result: "6",
		},
		{
			text:   "((lambda (x y) ( * x y)) (+ 1 2) (+ 3 4))",
			result: "21",
		},
		{
			before: "(define x 10)",
			text:   "(* x x)",
			result: "100",
		},
		{
			before: "(define second (lambda (x) (car (cdr x))))",
			text:   "(second (quote (1 2 3 4 5)))",
			result: "2",
		},
		{
			text:   "(% 3 2)",
			result: "1",
		},
		{
			before: "(define fb (lambda (n) (if (eq (% n 15) 0) fizzbuzz (if (eq (% n 3) 0) fizz (if (eq (% n 5) 0) buzz n)))))",
			text:   "(fb 30)",
			result: "fizzbuzz",
		},
		{
			before: "(define fb (lambda (n) (if (eq (% n 15) 0) fizzbuzz (if (eq (% n 3) 0) fizz (if (eq (% n 5) 0) buzz n)))))",
			text:   "(fb 9)",
			result: "fizz",
		},
		{
			before: "(define fb (lambda (n) (if (eq (% n 15) 0) fizzbuzz (if (eq (% n 3) 0) fizz (if (eq (% n 5) 0) buzz n)))))",
			text:   "(fb 10)",
			result: "buzz",
		},
		{
			before: "(define fb (lambda (n) (if (eq (% n 15) 0) fizzbuzz (if (eq (% n 3) 0) fizz (if (eq (% n 5) 0) buzz n)))))",
			text:   "(fb 23)",
			result: "23",
		},
	}

	for _, c := range cases {
		env := CreateEnv()
		if c.before != "" {
			eval(ParseText(c.before), env)
		}
		r := str(eval(ParseText(c.text), env))
		if r != c.result {
			t.Error(c.text, " is expected ", c.result, "but is actual ", r)
		}
	}
}
