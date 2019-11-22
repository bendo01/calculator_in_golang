package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func repeatHelper() {
	fmt.Println("Tekan huruf Y jika ingin melakukan perhitungan")
	fmt.Println("Tekan huruf T jika ingin keluar dari applikasi")
	fmt.Println("Apakah Anda Ingin Melakukan Perhitungan ?")
}

func inputHelper() {
	fmt.Println("Perhitungan Dengan Format operand1 spasi operator spasi operand2")
	fmt.Println("contoh: 1 + 1")
}

func calculate(inputUser string) {
	var operator1 float64
	var operator2 float64
	var result float64
	var operand string
	var tempArray []string

	operator1 = 0
	operator2 = 0
	operand = "+"
	result = 0

	tempArray = strings.Split(inputUser, " ")
	if len(tempArray) == 3 {
		// fmt.Printf("%q\n", tempArray)
		operator1, _ = strconv.ParseFloat(tempArray[0], 64)
		operand = tempArray[1]
		operator2, _ = strconv.ParseFloat(tempArray[2], 64)
	}

	// fmt.Println("isi operand " + operand)

	if strings.Compare("^", operand) == 0 {
		result = math.Pow(operator1, operator2)
	} else if strings.Compare("%", operand) == 0 {
		result = math.Mod(operator1, operator2)
	} else if strings.Compare("*", operand) == 0 && operator1 != 0 {
		result = operator1 * operator2
	} else if strings.Compare("/", operand) == 0 && (operator1 != 0 || operator2 != 0) {
		result = operator1 / operator2
	} else if strings.Compare("+", operand) == 0 {
		result = operator1 + operator2
	} else if strings.Compare("-", operand) == 0 {
		result = operator1 - operator2
	}

	fmt.Printf("Hasil Perhitungan %f\n", result)
}

func main() {
	var isCalculateAgain bool
	var holderIsCalculateAgain string
	var userInput string
	isCalculateAgain = true

	reader := bufio.NewReader(os.Stdin)
	for isCalculateAgain {
		inputHelper()
		userInput, _ = reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\n", "", -1)
		calculate(userInput)
		repeatHelper()
		holderIsCalculateAgain, _ = reader.ReadString('\n')
		holderIsCalculateAgain = strings.Replace(holderIsCalculateAgain, "\n", "", -1)
		isCalculateAgain = false
		if strings.Compare("Y", holderIsCalculateAgain) == 0 || strings.Compare("y", holderIsCalculateAgain) == 0 {
			isCalculateAgain = true
		}
	}
}
