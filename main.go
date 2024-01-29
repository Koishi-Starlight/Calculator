package main

import (
	"Calculator/calc"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	var num1, num2 float64
	var unconverted1, unconverted2, operator string
	var err error
	var exit = false

	fmt.Println("Application started.")
	for exit != false {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Enter mathematical operation in following format: 'a + b, a - b, a * b, a / b'\nOr type 'quit' to exit the application.")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		if text != "quit" {
			operationSlice := strings.Split(text, " ")
			if len(operationSlice) != 3 {
				fmt.Println("Incorrect amount of arguments or too many spaces between numbers and operator.")
			} else {
				unconverted1 = operationSlice[0]
				unconverted2 = operationSlice[2]
				operator = operationSlice[1]
				num1, err = strconv.ParseFloat(unconverted1, 64)
				if err != nil {
					fmt.Printf("Conversion error: #{err}")
				}
				num2, err = strconv.ParseFloat(unconverted2, 64)
				if err != nil {
					fmt.Printf("Conversion error: #{err}")
				}
				calculator := calc.NewCalc()
				result := calculator.Calculate(num1, num2, operator)
				fmt.Println("Result:", result)
				time.Sleep(time.Second)
			}
		} else {
			fmt.Println("Quitting...")
			exit = true
		}
	}
}
