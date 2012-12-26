package fibonacci

import "testing"
import "fmt"

func TestFibonacci(t *testing.T) {
  total := 0

  for i, n := 0, Fibonacci(0); n < 4000000; i, n = i + 1, Fibonacci(i + 1) {
    if (n % 2 == 0) {
      fmt.Printf("Adding Fibonacci(%d): %d\n", i, n)
      total += n
    }
  }

  fmt.Printf("Sum of fibonacci numbers under 4 million is %d\n", total)

}
