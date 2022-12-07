package main

func romanToInt(s string) int {
  var intFromRoman int = 0
  runes := []rune(s)
  for i := 0; i < len(runes); i++ {
    switch rSymbol := string(runes[i]); rSymbol {
    case "M":
      intFromRoman += 1000
    case "D":
      intFromRoman += 500
    case "C":
      if i+1 < len(runes) && (string(runes[i+1]) == "D" || string(runes[i+1]) == "M") {
        intFromRoman -= 100
      } else {
        intFromRoman += 100
      }
    case "L":
      intFromRoman += 50
    case "X":
      if i+1 < len(runes) && (string(runes[i+1]) == "L" || string(runes[i+1]) == "C") {
        intFromRoman -= 10
      } else {
        intFromRoman += 10
      }
    case "V":
      intFromRoman += 5
    case "I":
      if i+1 < len(runes) && (string(runes[i+1]) == "V" || string(runes[i+1]) == "X") {
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
  print(romanToInt("MCMXCIV"))
}
