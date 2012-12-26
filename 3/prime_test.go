package prime

import "testing"
import "fmt"

func TestPrime(t *testing.T) {

  num := Integer(600851475143)
  var largestPrimeFactor Integer
  for n := Integer(3); (n * n) < num; n += 2 {
    if num % n == 0 && n.isPrime() {
      largestPrimeFactor = n
    }
  }

  fmt.Printf("Largest prime factor: %d\n", largestPrimeFactor)
}
