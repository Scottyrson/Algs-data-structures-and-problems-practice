// Package brackets проверяет правильность скобочных последовательностей,
// состоящих из круглых (), квадратных [] и фигурных {} скобок.
//
// Последовательность считается правильной, если:
//
//  1. Пустая последовательность является правильной.
//  2. Если A — правильная последовательность, то (A), [A] и {A} — правильные.
//  3. Если A и B — правильные последовательности, то AB — правильная.
//
// Иными словами, каждая открывающая скобка должна иметь парную закрывающую
// того же типа, и пары должны быть правильно вложены.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReaderSize(os.Stdin, 1<<20)

	var s string
	fmt.Fscan(in, &s)

	if balanced(s) {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
}

func balanced(s string) bool {
	stack := make([]byte, 0, len(s))
	pair := [256]byte{
		')': '(',
		']': '[',
		'}': '{',
	}

	for i := 0; i < len(s); i++ {
		c := s[i]
		switch c {
		case '(', '[', '{':
			stack = append(stack, c)
		case ')', ']', '}':
			if len(stack) == 0 || stack[len(stack)-1] != pair[c] {
				return false
			}
			stack = stack[:len(stack)-1]
		default:
			return false
		}
	}

	return len(stack) == 0
}
