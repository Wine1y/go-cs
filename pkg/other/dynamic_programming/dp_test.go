package dynamicprogramming

import (
	"testing"

	tu "github.com/Wine1y/go-cs/internal/testing_utils"
)

func TestFibonacci(t *testing.T) {
	tu.AssertEqualsNamed(t, "Fibonacci returns correct number for N=0", Fibonacci(0), 0)
	tu.AssertEqualsNamed(t, "Fibonacci returns correct number for N=1", Fibonacci(1), 1)
	tu.AssertEqualsNamed(t, "Fibonacci returns correct number for N=2", Fibonacci(2), 1)
	tu.AssertEqualsNamed(t, "Fibonacci returns correct number for N=3", Fibonacci(3), 2)
	tu.AssertEqualsNamed(t, "Fibonacci returns correct number for N=4", Fibonacci(4), 3)
	tu.AssertEqualsNamed(t, "Fibonacci returns correct number for N=10", Fibonacci(10), 55)
	tu.AssertEqualsNamed(t, "Fibonacci returns correct number for N=30", Fibonacci(30), 832040)
}

func TestFactorial(t *testing.T) {
	tu.AssertEqualsNamed(t, "Factorial returns correct number for N=0", Factorial(0), 0)
	tu.AssertEqualsNamed(t, "Factorial returns correct number for N=1", Factorial(1), 1)
	tu.AssertEqualsNamed(t, "Factorial returns correct number for N=2", Factorial(2), 2)
	tu.AssertEqualsNamed(t, "Factorial returns correct number for N=3", Factorial(3), 6)
	tu.AssertEqualsNamed(t, "Factorial returns correct number for N=4", Factorial(4), 24)
	tu.AssertEqualsNamed(t, "Factorial returns correct number for N=8", Factorial(8), 40320)
	tu.AssertEqualsNamed(t, "Factorial returns correct number for N=10", Factorial(10), 3628800)
}

func TestLongestCommonSubsequence(t *testing.T) {
	tu.AssertEqualsNamed(t, "LCS returns correct sequnce length", LongestCommonSubsequence("ABAB", "ABAB"), 4)
	tu.AssertEqualsNamed(t, "LCS returns correct sequnce length", LongestCommonSubsequence("AQBWAEBR", "ATBYAUBI"), 4)
	tu.AssertEqualsNamed(t, "LCS returns correct sequnce length", LongestCommonSubsequence("ABAB", "QWERTYUIOPABAB"), 4)
	tu.AssertEqualsNamed(t, "LCS returns correct sequnce length", LongestCommonSubsequence("ABAB", "ABABQWERTYUIOP"), 4)
	tu.AssertEqualsNamed(t, "LCS returns correct sequnce length", LongestCommonSubsequence("ABAB", "QWERTYUIOP"), 0)
}
