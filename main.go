package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	var spliterinput []string

	fmt.Print("Введи операнды и оператор (например, 5 + 3 или V + II): ")

	scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    input := scanner.Text()

	spliterinput = strings.Split(input, " ")


	if len(spliterinput) < 3 {
		fmt.Print("Выдача паники, так как строка не является математической операцией")
		return
	}
	if len(spliterinput) > 3 {
		fmt.Print("Выдача паники, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *)")
		return
	}

	operand1 := spliterinput[0]
	operator := spliterinput[1]
	operand2 := spliterinput[2]

	/*fmt.Println(operand1)
	fmt.Println(operand2)
	fmt.Print(operator)*/

	operand1INT, isRoman1, err1 := toint(operand1)
	operand2INT, isRoman2, err2 := toint(operand2)

	if isRoman1 != isRoman2{
		fmt.Println("Разные системы исчисления")
		return
	}
	
	if err1 != nil {
		fmt.Println("Ошибка в первом операнде:", err1)
		return
	}
	if err1 != nil {
		fmt.Println("Ошибка во втором операнде:", err2)
		return
	}
	if operand1INT > 10 {
		fmt.Println("Ошибка: Первое число больше 10")
		return
	}
	if operand2INT > 10 {
		fmt.Println("Ошибка: Второе число больше 10")
		return
	}

	if operand1INT <= 0 {
		fmt.Println("Ошибка: Первое число равно 0 или меньше")
		return
	}
	
	if operand2INT <= 0 {
		fmt.Println("Ошибка: Второе число равно 0 или меньше")
		return
	}


	result := calculate(operand1INT, operand2INT, operator)
	if isRoman1 {
		fmt.Print("Результат", toRoman(result))
		return
	}
	fmt.Println("Результат", result)

}

func toint(operand string) (int, bool, error) {

	if intValue, err := strconv.Atoi(operand); err == nil {
		return intValue, false, nil
	}

	mapRoman := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
	}

	result := 0
	prevValue := 0

	for i := len(operand) - 1; i >= 0; i-- {
		currentValue, exists := mapRoman[operand[i]]
		if !exists {
			return 0, true, fmt.Errorf("Неверный символ в римском числе: %c", operand[i])
		}
		if currentValue < prevValue {
			result -= currentValue
		} else {
			result += currentValue
		}
		prevValue = currentValue
	}
	return result, true, nil
}

func calculate(operand1INT int, operand2INT int, operator string) int {
	var result int
	switch operator {
	case "+":
		result = operand1INT + operand2INT
	case "-":
		result = operand1INT - operand2INT
	case "*":
		result = operand1INT * operand2INT
	case "/":
		if operand2INT == 0 {
			fmt.Print("Так делить нельзя")
			return 0
		}
		result = operand1INT / operand2INT
	default:
		fmt.Println("Ошибка неверный оператор")
		return 0
	}
	return result
}

func toRoman(num int) string {
	if num < 1 {
		return "Выдача паники, так как в римской системе нет отрицательных чисел."
	}
	var result strings.Builder
	romanNumerals := []struct {
		Value  int
		Symbol string
	}{
		{100, "C"}, {90, "XC"}, {50, "L"}, {40, "XL"},
		{10, "X"}, {9, "IX"}, {5, "V"}, {4, "IV"}, {1, "I"},
	}
	for _, numeral := range romanNumerals {
		for num >= numeral.Value {
			result.WriteString(numeral.Symbol)
			num -= numeral.Value
		}
	}
	return result.String()
}
