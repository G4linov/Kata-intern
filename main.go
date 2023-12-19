package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Roman(number int) string {
	conversions := []struct {
		value int
		digit string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	roman := ""
	for _, conversion := range conversions {
		for number >= conversion.value {
			roman += conversion.digit
			number -= conversion.value
		}
	}
	return roman
}

func main() {
	var input, outstr, swingtype string
	var out int
	rim := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
	}
	for {
		input, _ = bufio.NewReader(os.Stdin).ReadString('\n')
		num1 := 0
		num2 := 0
		arr := strings.Fields(input)
		if arr[0] == "exit" {
			break
		}
		if len(arr) != 3 {
			fmt.Println("Input operation error")
			continue
		}
		if _, err := strconv.Atoi(arr[0]); err != nil {
			if _, err := strconv.Atoi(arr[2]); err != nil {
				prev := 0
				for _, el := range arr[0] {
					if prev < rim[el] {
						num1 = num1 - prev*2
					}
					num1 = num1 + rim[el]
				}
				prev = 0
				for _, el := range arr[2] {
					if prev < rim[el] {
						num2 = num2 - prev*2
					}
					num2 = num2 + rim[el]
				}
				swingtype = "rom"
			} else {
				fmt.Println("Type error")
				continue
			}
		} else {
			if _, err := strconv.Atoi(arr[2]); err == nil {
				num1, _ = strconv.Atoi(arr[0])
				num2, _ = strconv.Atoi(arr[2])
				swingtype = "arab"
			} else {
				fmt.Println("Type error")
				continue
			}
		}
		if num1 > 10 || num1 < 1 || num2 > 10 || num2 < 1 {
			fmt.Println("Value error, some number not in range")
			continue
		}
		switch swingtype {
		case "arab":
			switch arr[1] {
			case "+":
				fmt.Println(num1 + num2)
			case "-":
				fmt.Println(num1 - num2)
			case "/":
				fmt.Println(num1 / num2)
			case "*":
				fmt.Println(num1 * num2)
			default:
				fmt.Println("Unknow function")
			}
		case "rom":
			switch arr[1] {
			case "+":
				out = num1 + num2
				outstr = Roman(out)
				fmt.Println(outstr)
			case "-":
				if num1-num2 < 1 {
					fmt.Println("No negative or zero in roman")
				} else {
					out = num1 - num2
					outstr = Roman(out)
					fmt.Println(outstr)
				}
			case "/":
				if num1-num2 < 1 {
					fmt.Println("No negative or zero in roman")
				} else {
					out = num1 / num2
					outstr = Roman(out)
					fmt.Println(outstr)
				}
			case "*":
				out = num1 * num2
				outstr = Roman(out)
				fmt.Println(outstr)
			default:
				fmt.Println("Unknow function")
			}
		}
	}
}
