package l6range

func concurrentFib(n int) []int {
	ch := make(chan int)

	go fibonacci(n, ch)

	result := make([]int, 0, n)

	for value := range ch {
		result = append(result, value)
	}

	return result
}

// don't touch below this line

func fibonacci(n int, ch chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
	}
	close(ch)
}
