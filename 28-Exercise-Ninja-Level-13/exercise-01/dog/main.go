package dog

func Years(n int) int {
	return n * 7
}

func YearTwo(n int) int {
	count := 0
	for i := 0; i < n; i++ {
		count += 7
	}
	return count
}
