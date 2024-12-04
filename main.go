package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func intToRoman(num int) string {
	var roman string = ""
	var numbers = []int{1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000}
	var romans = []string{"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C", "CD", "D", "CM", "M"}
	var index = len(romans) - 1
	if num < 1 {
		log.Panic("В римской системе счисления нет отрицательных цифр и цифр меньше 1 (I)")
	} else {
		for num > 0 {
			for numbers[index] <= num {
				roman += romans[index]
				num -= numbers[index]
			}
			index -= 1
		}
	}

	return roman
}

func isParsable(s string) bool {
	roman := map[string]int{"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10}
	if _, key := roman[s]; key {
		return true
	}
	if value, err := strconv.Atoi(s); err == nil {
		if value <= 10 && value >= 1 {
			return true
		} else {
			log.Panic("Дозволены цифры от 1 до 10")
			return false
		}
	}
	return false
}

func main() {
	for true {

		roman := map[string]int{"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10}
		operators := [4]string{"+", "-", "*", "/"}
		var str1, str2 string
		//var num1, num2 int
		var operator string
		var expr []string
		var sum int

		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Введите выражение:")

		input, err := reader.ReadString('\n')
		if err != nil {
			log.Panic(err)
		}
		if len(input) <= 2 {
			log.Panic("Строка не является математической операцией.")
		}

		for i := range operators {
			if strings.ContainsAny(input, operators[i]) {
				operator = operators[i]
				expr = strings.Split(strings.ToUpper(input), operator)
				expr[0] = strings.TrimSpace(expr[0])
				expr[1] = strings.TrimSpace(expr[1])

				if len(expr) > 2 {
					log.Panic("Допускается только два операнда (от 1 или I до 10 или X) и один оператор")
				}

				if isParsable(expr[0]) && isParsable(expr[1]) {
					str1, str2 = strings.TrimSpace(expr[0]), strings.TrimSpace(expr[1])
				} else {
					log.Panic("Строка не является математической операцией")
				}

				if isParsable(str1) && isParsable(str2) {
					if value1, key1 := roman[str1]; key1 {
						if value2, key2 := roman[str2]; key2 {
							// Оба операнда римские
							switch operator {
							case "+":
								sum = value1 + value2
								fmt.Println(intToRoman(sum))
							case "-":
								sum = value1 - value2
								if sum < 1 {
									log.Panic("Паника: в римской системе нет отрицательных чисел.")
								}
								fmt.Println(intToRoman(sum))
							case "*":
								sum = value1 * value2
								fmt.Println(intToRoman(sum))
							case "/":
								sum = value1 / value2
								fmt.Println(intToRoman(sum))
							}
						} else {
							log.Panic("Разные системы счисления")
						}
					} else if value1, err := strconv.Atoi(str1); err == nil {
						if value2, err := strconv.Atoi(str2); err == nil {
							// Оба операнда арабские
							switch operator {
							case "+":
								fmt.Println(value1 + value2)
							case "-":
								fmt.Println(value1 - value2)
							case "*":
								fmt.Println(value1 * value2)
							case "/":
								if value2 == 0 {
									log.Panic("Деление на ноль.")
								}
								fmt.Println(value1 / value2)
							}
						} else {
							log.Panic("Разные системы счисления")
						}
					}
				}
			}
		}
	}
}
