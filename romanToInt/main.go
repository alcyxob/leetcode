package main

import "strconv"

func romanToInt(s string) int {
  var intFromRoman int = 0
  runes := []rune(s)
  for i := 0; i < len(runes); i++ {
    print(s[i])
    switch rSymbol := runes[i]; rSymbol {
    case "M":
      intFromRoman += 1000
    case "D":
      intFromRoman += 500
    case "C":
      intFromRoman += 100
    case "L":
      intFromRoman += 50
    case "X":
      intFromRoman += 10
    case "V":
      intFromRoman += 5
    case "I":
      if strconv.Itoa(int(s[i+1])) != "I" {
        intFromRoman -= 1
      } else {
        intFromRoman += 1
      }
    default:
      continue
    }
  }
  return intFromRoman
}

func main() {
  //print(string(179))
  print(romanToInt("LVIII"))
}
