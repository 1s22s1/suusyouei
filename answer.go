package answer

func divisor(n int) []int {
	var x []int

	for i := 1; i <= n; i++ {
		if (n % i) == 0 {
			x = append(x, i)
		}
	}

	return x
}

func q1(n int) []int {
	return divisor(n)
}
