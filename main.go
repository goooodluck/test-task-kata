package main

import (
 "fmt"
 "regexp"
 "strconv"
 "strings"
 "bufio"
 "os"
)

var romanNumerals = map[int]string{
	1:   "I", 2: "II", 3: "III", 4: "IV", 5: "V",
	6:   "VI", 7: "VII", 8: "VIII", 9: "IX", 10: "X",
	11:  "XI", 12: "XII", 13: "XIII", 14: "XIV", 15: "XV",
	16:  "XVI", 17: "XVII", 18: "XVIII", 19: "XIX", 20: "XX",
	21:  "XXI", 22: "XXII", 23: "XXIII", 24: "XXIV", 25: "XXV",
	26:  "XXVI", 27: "XXVII", 28: "XXVIII", 29: "XXIX", 30: "XXX",
	31:  "XXXI", 32: "XXXII", 33: "XXXIII", 34: "XXXIV", 35: "XXXV",
	36:  "XXXVI", 37: "XXXVII", 38: "XXXVIII", 39: "XXXIX", 40: "XL",
	41:  "XLI", 42: "XLII", 43: "XLIII", 44: "XLIV", 45: "XLV",
	46:  "XLVI", 47: "XLVII", 48: "XLVIII", 49: "XLIX", 50: "L",
	51:  "LI", 52: "LII", 53: "LIII", 54: "LIV", 55: "LV",
	56:  "LVI", 57: "LVII", 58: "LVIII", 59: "LIX", 60: "LX",
	61:  "LXI", 62: "LXII", 63: "LXIII", 64: "LXIV", 65: "LXV",
	66:  "LXVI", 67: "LXVII", 68: "LXVIII", 69: "LXIX", 70: "LXX",
	71:  "LXXI", 72: "LXXII", 73: "LXXIII", 74: "LXXIV", 75: "LXXV",
	76:  "LXXVI", 77: "LXXVII", 78: "LXXVIII", 79: "LXXIX", 80: "LXXX",
	81:  "LXXXI", 82: "LXXXII", 83: "LXXXIII", 84: "LXXXIV", 85: "LXXXV",
	86:  "LXXXVI", 87: "LXXXVII", 88: "LXXXVIII", 89: "LXXXIX", 90: "XC",
	91:  "XCI", 92: "XCII", 93: "XCIII", 94: "XCIV", 95: "XCV",
	96:  "XCVI", 97: "XCVII", 98: "XCVIII", 99: "XCIX", 100: "C",
  
}

// Функция для преобразования римского числа в арабское
func romanToArab(roman string) int {
 for value, numeral := range romanNumerals {
  if numeral == roman {
   return value
  }
 }
 return 0
}

// Функция для преобразования арабского числа в римское
func arabToRoman(number int) string {
 if number < 1 {
  panic("Выдача паники, так как в римской системе нет отрицательных чисел.")
 }
 return romanNumerals[number]
}

// Функция для выполнения арифметической операции
func calculate(a, b int, operator string) int {
 switch operator {
 case "+":
  return a + b
 case "-":
  return a - b
 case "*":
  return a * b
 case "/":
  if b == 0 {
   panic("Выдача паники, так как невозможно деление на ноль.")
  }
  return a / b
 default:
  panic("Выдача паники, так как неизвестная операция.")
 }
}

// Функция для проверки, являются ли все строки числами одного типа (либо римские, либо арабские)
func validateInput(aStr, bStr string) (int, int, bool) {
 arabA, errA := strconv.Atoi(aStr)
 arabB, errB := strconv.Atoi(bStr)

 if errA == nil && errB == nil {
  if arabA < 1 || arabA > 9 || arabB < 1 || arabB > 9 {
   panic("Числа должны быть от 1 до 9.")
  }
  return arabA, arabB, true
 }

 romanA := romanToArab(aStr)
 romanB := romanToArab(bStr)

 if romanA == 0 || romanB == 0 {
  panic("Выдача паники, так как используются одновременно разные системы счисления.")
 }

 return romanA, romanB, false
}
   

func main() {
 reader := bufio.NewReader(os.Stdin)
 for {
  fmt.Print("Введите выражение (a + b, a - b, a * b, a / b): ")
  input, _ := reader.ReadString('\n')
  input = strings.TrimSpace(input)

  re := regexp.MustCompile(`^(\w+)\s*([+\-*/])\s*(\w+)$`)
  matches := re.FindStringSubmatch(input)

  if len(matches) < 1 {
	panic("Выдача паники, так как строка не является математической операцией.")
  } 

  aStr, operator, bStr := matches[1], matches[2], matches[3]
  a, b, isArabic := validateInput(aStr, bStr)

  var result int
  if isArabic {
   result = calculate(a, b, operator)
   fmt.Println(result)
  } else {
   result = calculate(a, b, operator)
   fmt.Println(arabToRoman(result))
  }
 }
}
