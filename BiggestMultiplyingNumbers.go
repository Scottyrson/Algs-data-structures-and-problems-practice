// Задача: дан список целых чисел. Найти два числа, произведение которых максимально.
//
// Формат ввода:
//   В единственной строке через пробел вводятся целые числа — элементы списка.
//   Список содержит не менее двух и не более 100 000 чисел.
//   Сами элементы по модулю не превышают 1 000 000.
//
// Формат вывода:
//   Выведите эти два числа в порядке неубывания.
//
// Примечание:
//   Решение должно иметь сложность O(n), где n — размер списка.
//   Гарантируется, что во всех тестах ответ однозначен.

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 0, 1<<20), 1<<20)
	if !scanner.Scan() {
		fmt.Println("No input")
		return
	}
	var numbers []int
	for _, field := range strings.Fields(scanner.Text()) {
		num, err := strconv.Atoi(field)
		if err != nil {
			fmt.Printf("Warning: '%s' is not a valid integer. Skipping.\n", field)
			continue
		}
		numbers = append(numbers, num)
	}

	if len(numbers) < 2 {
		fmt.Println("List must contain at least two numbers.")
		return
	}

	sort.Ints(numbers)

	if numbers[0]*numbers[1] <= numbers[len(numbers)-1]*numbers[len(numbers)-2] {
		fmt.Print(numbers[len(numbers)-2], " ", numbers[len(numbers)-1])
	} else {
		fmt.Print(numbers[0], " ", numbers[1])
	}
}
