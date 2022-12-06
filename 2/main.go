package main

import "strconv"

func isPalindrome(x int) bool {
  var invStr string = ""
  for _, v := range strconv.Itoa(x) {
    invStr = string(v) + invStr
  }
  invInt, _ := strconv.Atoi(invStr)

  if x == invInt {
    return true
  }
  return false
}

func main() {
  print(isPalindrome(0))
}
