package main

import (
	"fmt"
	"sync"
)

// Конкурентно порахувати суму кожного слайсу int, та роздрукувати результат.
// Потрібно використовувати WaitGroup.
// Приклад:
// [ [ 4, 6 ], [ 7, 9 ] ]
// Результат друку:
// Порядок друку не важливий.
// “slice 1: 10”
// “slice 2: 16”

func main() {
	// Розкоментуй мене)
	var wg sync.WaitGroup
	n := [][]int{
		{2, 6, 9, 24},
		{7, 3, 94, 3, 0},
		{4, 2, 8, 35},
	}
	//use channels only for test  purpose
	ch := make(chan int, len(n))
	wg.Add(len(n))
	for i := 0; i < len(n); i++ {
		go func(incr int) {
			ch <- incr + 1
			defer wg.Done()
			sum(n[incr], ch)
		}(i)
	}
	wg.Wait()
}

func sum(sl []int, ch chan int) {
	summa := 0
	counter := <-ch
	for _, v := range sl {
		summa += v
	}
	fmt.Printf("slice %d: %d\n", counter, summa)
}
