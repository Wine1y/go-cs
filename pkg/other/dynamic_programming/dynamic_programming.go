package dynamicprogramming

func Fibonacci(n int) int {
	if n == 0 || n == 1 {
		return n
	}

	first, second := 0, 1
	for i := 2; i <= n; i++ {
		first, second = second, first+second
	}
	return second
}

func Factorial(n int) int {
	if n == 0 || n == 1 {
		return n
	}

	mem := make([]int, n+1)
	mem[0] = 1

	for i := 1; i <= n; i++ {
		mem[i] = i * mem[i-1]
	}
	return mem[n]
}

func LongestCommonSubsequence(a, b string) int {
	mem := make([][]int, len(a)+1)
	for i := 0; i < len(a)+1; i++ {
		mem[i] = make([]int, len(b)+1)
	}

	for i := 0; i <= len(a); i++ {
		for j := 0; j <= len(b); j++ {
			if i == 0 || j == 0 {
				mem[i][j] = 0
			} else if a[i-1] == b[j-1] {
				mem[i][j] = mem[i-1][j-1] + 1
			} else {
				mem[i][j] = max(mem[i-1][j], mem[i][j-1])
			}
		}
	}

	return mem[len(a)][len(b)]
}
