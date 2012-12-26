package prime

type Integer int64

func (n Integer) isPrime() bool {
  if (n < Integer(3)) {
    return true
  }

  for i := Integer(3); (i * i) < Integer(n); i += 2 {
    if n % i == 0 {
      return false
    }
  }

  return true
}
