package multiples

func sum(max int) int {
  total := 0
  for i := 0; i < max; i++ {
    if (i % 3 == 0 || i % 5 == 0 ) {
      total += i
    }
  }

  return total
}
