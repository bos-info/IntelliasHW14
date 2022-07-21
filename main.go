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
	var wg sync.WaitGroup
	n := [][]int{
		{2, 6, 9, 24},
		{7, 3, 94, 3, 0},
		{4, 2, 8, 35},
	}
	wg.Add(len(n))
	for i := 0; i < len(n); i++ {
		//оскільки вимагається передача ЛИШЕ слайсу інтів до функції sum
		//то за допомогою замикання ми створюємо модифікований слайс, де перше(0) значення завжди номер слайсу
		//якщо додалиб рядок створення модифікованого слайсу всередину функції отримали б гонку
		//оскільки апенд дорога операція створюємо слайс заздалегідь відомого розміру
		modSlice := make([]int, 1, len(n[i]))
		modSlice[0] = i + 1
		go func(incr int) {
			modSlice = append(modSlice, n[incr]...)
			defer wg.Done()
			sum(modSlice)
		}(i)
	}
	wg.Wait()
}

func sum(sl []int) {
	summa := 0
	counter := sl[0]
	for _, v := range sl[1:] {
		summa += v
	}
	fmt.Printf("slice %d: %d\n", counter, summa)
}
