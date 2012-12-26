package fibonacci

var CACHE = map[int]int {}

func Fibonacci(n int) int {
  if (n == 0) {
    return 0;
  } else if (n == 1) {
    return 1;
  } else if (CACHE[n] == 0) {
    CACHE[n] = Fibonacci(n-1) + Fibonacci(n-2)
  }

  return CACHE[n]
}
