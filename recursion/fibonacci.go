package main

func main() {
	var n = 80
	//println(fib(n))

	var memo = make([]int, n+1)

	println(fibDPMemo(n, memo))

	println(fibDPTabulation(n))
}

func fib(n int) int {
	if n <= 1 {
		return n
	}

	return fib(n-1) + fib(n-2)
}

func fibDPMemo(n int, memo []int) int {
	if n <= 1 {
		return n
	}

	if memo[n] > 0 {
		return memo[n]
	}

	memo[n] = fibDPMemo(n-1, memo) + fibDPMemo(n-2, memo)

	return memo[n]
}

func fibDPTabulation(n int) int {
	if n <= 1 {
		return n
	}

	var tab = make([]int, n+1)

	tab[0] = 0
	tab[1] = 1

	for i := 2; i < len(tab); i++ {
		tab[i] = tab[i-1] + tab[i-2]
	}

	return tab[n]
}
