package main

import (
	"strconv"
	"strings"
)

func parseText(tokens []string) *Cons {
	parent := &Cons{}
	result := parent
	var now *Cons
	nestCount := 0
	var startIndex int
	for i, token := range tokens {
		switch token {
		case "(":
			if nestCount == 0 {
				startIndex = i
			}
			nestCount++
		case ")":
			nestCount--
			if nestCount == 0 {
				now = &Cons{
					parseText(tokens[startIndex+1 : i]),
					nil,
				}
				parent.Cdr = now
				parent = now
			}
		default:
			if nestCount == 0 {
				i, err := strconv.Atoi(token)
				var now *Cons
				if err != nil {
					now = &Cons{
						token,
						nil,
					}
				} else {
					now = &Cons{
						i,
						nil,
					}
				}
				parent.Cdr = now
				parent = now
			}
		}
	}
	return result.Cdr.(*Cons)
}

// ParseText parse string
func ParseText(text string) *Cons {
	return parseText(strings.Fields(text)).Car.(*Cons)
}
