package multiples

import "testing"
import "fmt"

func TestSum(t *testing.T) {

  samples := []int{10, 1000}

  for _, n := range samples {
    fmt.Printf("Sum of %d is %d\n", n, sum(n))
  }
}
