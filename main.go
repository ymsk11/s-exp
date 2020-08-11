package main

import (
	"bufio"
	"os"
)

func main() {
	env := CreateEnv()
	stdin := bufio.NewScanner(os.Stdin)
	for stdin.Scan() {
		text := stdin.Text()
		print(eval(ParseText(text), env))
	}
}
