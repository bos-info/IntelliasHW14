package main

import (
	"fmt"
)

// Конкурентно порахувати суму усіх слайсів int, та роздрукувати результат.
// Приклад:
// [ [ 4, 6 ], [ 7, 9 ] ]
// Результат друку:
// “result: 26”
func main() {
	// Розкоментуй мене)
	n := [][]int{
		{2, 6, 9, 24},
		{7, 3, 94, 3, 0},
		{4, 2, 8, 35},
	}
	// Ваша реалізація
	result := 0
	ch := make(chan int, len(n))
	for _, v := range n {
		go func(sl []int) {
			sum := 0
			for _, v := range sl {
				sum += v
			}
			ch <- sum
		}(v)
	}
	for i := 0; i < len(n); i++ {
		result += <-ch
	}
	fmt.Println("result:", result)
}
