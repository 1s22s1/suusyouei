package answer

import "os"

func divisor(n int) []int {
	var x []int

	for i := 1; i <= n; i++ {
		if (n % i) == 0 {
			x = append(x, i)
		}
	}

	return x
}

func isAbundant(n int) bool {
	var sum int = 0

	for _, x := range divisor(n) {
		sum += x
	}

	if sum-n > n {
		return true
	} else {
		return false
	}
}

func collatz(n int) []int {
	var x []int = []int{n}

	for {
		if (n % 2) == 0 {
			n = n / 2
		} else {
			n = n*3 + 1
		}

		x = append(x, n)

		if n == 1 {
			return x
		}
	}
}

func q1(n int) []int {
	return divisor(n)
}

func q2(n int) int {
	for {
		if isAbundant(n) {
			return n
		}

		if n < 0 {
			os.Exit(1)
		}

		n--
	}
}

func q3(n int) int {
	for {
		if len(collatz(n)) == 256 {
			return n
		}

		n++
	}
}
