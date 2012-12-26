package palindrome

import "strconv"

func Reverse(s string) string {
  runes := make([]rune, len(s))

  for i, char := range s {
    runes[len(s)-i-1] = char
  }

  return string(runes)
}

func isPalindrome(n int) bool {
  s := strconv.Itoa(n)

  return s == Reverse(s)
}
