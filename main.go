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

	stext := "( + ( * 10 10 ) 41 )"
	cell := ParseText(stext)
	fmt.Println(cell.Str())
	fmt.Println(eval(cell))
}
