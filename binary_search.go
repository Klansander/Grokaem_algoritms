package main

import (
	"bufio"
	"fmt"
	"os"
)

func binary_search(num int, slice []int) (pos int) {
	low := 0
	high := len(slice)
	med := 0
	for low < high {
		med = (low + high) / 2
		if slice[med] == num {
			pos = med + 1
			break
		} else if slice[med] > num {
			high = med
		} else {
			low = med

		}
	}
	return

}
func main() {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26}
	num := 0
	in := bufio.NewReader(os.Stdin)
	fmt.Println("Введите искомое число")
	fmt.Fscan(in, &num)
	fmt.Println("Позиция заданного числа(значение 0, сигнализирует об отсутствии искомого числа):", binary_search(num, array))

}
