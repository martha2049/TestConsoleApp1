package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Hello and welcome to the Calculator!\n" + //greetings and instructions for a user
		"Enter two numbers in a line and one of the arithmetic operators (+,-,*,/) between them for calculation.\n" +
		"You can use both Arabic and Roman numerals, but of the same number system.\n" +
		"\n\n" +
		"Enter the expression:")
	intType, num1, num2, sign, err := readLine() // read the input and return a string
	if err != nil {                              // error testing
		fmt.Println("Input error:\n", err)
	}
	if intType == "arab" { //for arabic numerals
		num1Num, err1 := strconv.Atoi(num1) //convert alphanumeric to integer
		if err1 != nil {                    // error testing
			fmt.Println("Error when converting a string to a number:\n", err1)
			panic(err1)
		}
		num2Num, err2 := strconv.Atoi(num2)
		if err2 != nil { // error testing
			fmt.Println("Error when converting a string to a number:\n", err2)
			panic(err2)
		}
		res, err3 := calculator(num1Num, num2Num, sign)
		if err3 != nil { // error testing
			fmt.Println("Error when working with the calculator:\n", err3)
			panic(err3)
		} else {
			fmt.Println("Result: ", res) //result
		}
	} else {
		num1Num := fromRomanToInt(num1) //conversion
		num2Num := fromRomanToInt(num2)
		res, err1 := calculator(num1Num, num2Num, sign)
		if err1 != nil { // error testing
			fmt.Println("Error when working with the calculator:\n", err1)
			panic(err1)
		} else {
			final, err2 := fromIntToRoman(res)
			if err2 != nil { // error testing
				fmt.Println("Error when working with the calculator:\n", err2)
				panic(err2)
			}
			fmt.Println("Result: ", final)
		}
	}
}

func calculator(num1 int, num2 int, sign string) (int, error) {
	if num1 > 10 || num2 > 10 {
		fmt.Println("Please, enter the numbers from 1 to 10 inclusively.")
		panic(num1)
	} //limit for the operands from 1 to 10
	switch { //cases of arithmetic operations possible
	case sign == "+":
		return num1 + num2, nil
	case sign == "-":
		return num1 - num2, nil
	case sign == "*":
		return num1 * num2, nil
	case sign == "/" && num2 != 0:
		return num1 / num2, nil
	case sign == "/" && num2 == 0:
		fmt.Println("Cannot be divided by zero.")
		panic(num2)
	default:
		fmt.Println(":( Something went wrong during the calculation. Need time to resolve.")
		panic(num2)
	}
} // read the input and return a string
func readLine() (string, string, string, string, error) {
	stdin := bufio.NewReader(os.Stdin) //buffered reader
	usInput, _ := stdin.ReadString('\n')
	usInput = strings.TrimSpace(usInput)
	intType, num1, num2, sign, err := checkInput(usInput)
	if err != nil {
		return "", "", "", "", err
	}
	return intType, num1, num2, sign, err
}

// check input
func checkInput(input string) (string, string, string, string, error) {
	r := regexp.MustCompile("\\s+") // regular expressions, return an error
	replace := r.ReplaceAllString(input, "")
	arr := strings.Split(replace, "")
	var intType, num1, num2, sign string
	for index, value := range arr { //iterate and return as integer
		isN := isNumber(value)
		isS := isSign(value)
		isR := isRomanNumber(value)
		if !isN && !isS && !isR {
			fmt.Println("Unrecognized characters. Please, use only Arabic or Roman numerals and arithmetic operators +, -, /, *.")
			panic(intType)
		} //func from errors pack to resolve panic with an error message
		if isS {
			if sign != "" {
				fmt.Println("Undetected input. Please, enter two numbers and one arithmetic operator only.")
				panic(intType)
			} else {
				sign = arr[index]
			}
		}
		if (isN && intType != "roman") || (isR && intType != "arab") {
			if intType == "" {
				if isN {
					intType = "arab"
				} else {
					intType = "roman"
				}
			}
			if num1 == "" && !(index+1 == len(arr)) && isSign(arr[index+1]) {
				slice := arr[0:(index + 1)]
				num1 = strings.Join(slice, "") //join in one string
			} else if index+1 == len(arr) && num1 != "" {
				slice := arr[(len(num1) + 1):]
				num2 = strings.Join(slice, "") //join in one string
			}
		} else if (intType == "arab" && isR) || (intType == "roman" && isN) {
			fmt.Println("Wrong input. Please, use numerals of the same number system (only Arabic or only Roman).")
			panic(intType)
		}
	}
	if num2 == "" || num1 == "" || sign == "" {
		fmt.Println("Wrong number of arguments. Please, enter two numbers and an arithmetic operator.")
		panic(intType)
	}
	return intType, num1, num2, sign, nil
}

func isNumber(c string) bool {
	if c >= "0" && c <= "9" {
		return true
	} else {
		return false
	}
}

func isSign(c string) bool {
	if c == "+" || c == "-" || c == "/" || c == "*" {
		return true
	} else {
		return false
	}
}
func isRomanNumber(c string) bool {
	_, ok := dict[c]
	if ok {
		return true
	} else {
		return false
	}
}

// variables dictionary for roman numerals
var dict = map[string]int{
	"M":  1000,
	"CM": 900,
	"D":  500,
	"CD": 400,
	"C":  100,
	"XC": 90,
	"L":  50,
	"XL": 40,
	"X":  10,
	"IX": 9,
	"V":  5,
	"IV": 4,
	"I":  1,
}

// conversions roman-arabic arabic-roman
func fromRomanToInt(roman string) int {
	var res int
	arr := strings.Split(roman, "")
	for index, value := range arr {
		if index+1 != len(arr) && dict[value] < dict[arr[index+1]] {
			res -= dict[value]
		} else {
			res += dict[value]
		}
	}
	return res
}

func fromIntToRoman(number int) (string, error) {
	if number <= 0 {
		fmt.Println("The result can't be a negative number!")
		panic(number)
	}
	arr1 := [13]int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	arr2 := [13]string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	var str string
	for number > 0 {
		for i := 0; i < 13; i++ {
			if arr1[i] <= number {
				str += arr2[i]
				number -= arr1[i]
				break //stop the cycle
			}
		}
	}
	return str, nil
}
