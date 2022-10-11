package loops

func SumOfFirst(n int) int {
	sum := 0
	for i := 0; i <= n; i++ {
		sum += i
	}

	return sum
}
