package main

func longestCommonPrefix(strs []string) string {
  var prefix []uint8
  isCommonPrefix := true
  var curChar uint8
  var i = 0
  for isCommonPrefix {
    if curChar != 0 {
      prefix = append(prefix, curChar)
    }
    curChar = 0
    for _, word := range strs {
      if curChar == 0 && len(word) > 0 {
        curChar = word[i]
      } else {
        if len(word) == 0 || word[i] != curChar {
          isCommonPrefix = false
        }
      }
    }
    i++
  }
  return string(prefix)
}

func main() {
  strs := []string{"a"}

  print(longestCommonPrefix(strs))
}
