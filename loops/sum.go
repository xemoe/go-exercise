package loops

func SumOfFirst(n int, sum int) int {
	if n >= 0 {
		sum += n
		n--
		return SumOfFirst(n, sum)
	}

	return sum
}
