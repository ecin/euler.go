package palindrome

import "testing"
import "fmt"

func TestReverse(t *testing.T) {

  pairs := map[string]string {"a":"a", "ab": "ba", "aaa": "aaa", "abcd": "dcba"}

  for s, result := range pairs {
    if Reverse(s) != result {
      t.Errorf("%s is not the reverse of %s", Reverse(s), s)
    }
  }

}

func TestPalindrome(t *testing.T) {

  if !isPalindrome(999) {
    t.Error("999 should be a palindrome\n")
  }

  if isPalindrome(998) {
    t.Error("998 shouldn't be a palindrome\n")
  }

  if !isPalindrome(9) {
    t.Error("9 should be a palindrome\n")
  }

}

func TestMain(t *testing.T) {
  max := 0

  for i := 100; i < 1000; i++ {
    for j := 100; j < 1000; j++ {
      if isPalindrome(i * j) && max < i * j {
        max = i * j
      }
    }
  }

  fmt.Printf("Largest palindrome product of two 3-digit numbers: %d\n", max)
}
