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
				now := &Cons{
					castToken(token),
					nil,
				}
				parent.Cdr = now
				parent = now
			}
		}
	}
	return result.Cdr.(*Cons)
}

func castToken(token string) any {
	if token == "t" || token == "T" {
		return &T{}
	}
	i, err := strconv.Atoi(token)
	if err == nil {
		return i
	}
	f, err := strconv.ParseFloat(token, 64)
	if err == nil {
		return f
	}
	return token
}

// ParseText parse string
func ParseText(text string) *Cons {
	text = strings.Replace(text, "(", " ( ", -1)
	text = strings.Replace(text, ")", " ) ", -1)
	return parseText(strings.Fields(text)).Car.(*Cons)
}
