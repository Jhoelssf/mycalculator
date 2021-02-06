package mycalculator

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Elija una operacion:")
	fmt.Println("1: suma")
	fmt.Println("2: resta")
	fmt.Println("3: multiplicacion")
	fmt.Println("4: division")
	operator, _ := strconv.Atoi(ReadKeyboard())
	fmt.Println("ingresa la cadena Ejm ====> 4+4+4")

	switch operator {
	case 1:
		values := strings.Split(ReadKeyboard(), "+")
		valuesInteger := valuesIntegers(values)
		fmt.Println(sum(valuesInteger))
	case 2:
		values := strings.Split(ReadKeyboard(), "-")
		valuesInteger := valuesIntegers(values)
		fmt.Println(subs(valuesInteger))
	case 3:
		values := strings.Split(ReadKeyboard(), "*")
		valuesInteger := valuesIntegers(values)
		fmt.Println(mul(valuesInteger))
	case 4:
		values := strings.Split(ReadKeyboard(), "/")
		valuesInteger := valuesIntegers(values)
		fmt.Println(div(valuesInteger))
	default:
		fmt.Println("Invalid input")
	}
}

func ReadKeyboard() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func valuesIntegers(numbers []string) []int {
	var arrayNumbers []int
	for _, value := range numbers {
		operator, _ := strconv.Atoi(value)
		arrayNumbers = append(arrayNumbers, operator)
	}
	return arrayNumbers
}

func sum(numbers []int) int {
	result := 0
	for _, value := range numbers {
		result += value
	}
	return result
}

func mul(numbers []int) int {
	result := 0
	previewValue := 0
	for i, value := range numbers {
		if i == 0 {
			previewValue = value
		} else {
			result = previewValue * value
			previewValue = result
		}
	}
	return result
}

func subs(numbers []int) int {
	result := 0
	previewValue := 0
	for i, value := range numbers {
		if i == 0 {
			previewValue = value
		} else {
			result = previewValue - value
			previewValue = result
		}
	}
	return result
}

func div(numbers []int) float64 {
	result := 0.0
	previewValue := 0.0
	for i, value := range numbers {
		if i == 0 {
			previewValue = float64(value)
		} else {
			result = previewValue / float64(value)
			previewValue = result
		}
	}
	return result
}