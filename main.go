package main

import (
	"fmt"

	go_decimal_converter "github.com/cahayoyo/go-decimal-converter"
)

func main() {
	var input int
	var angka int
	fmt.Println("Decimal Number Conversion Program")
	fmt.Println("1. Decimal To Biner")
	fmt.Println("2. Decimal To Octal")
	fmt.Println("3. Decimal To Hexa")
	fmt.Println("Enter Your Choice :")
	fmt.Scan(&input)
	fmt.Println("Enter The Number You Want To Convert")
	fmt.Scan(&angka)

	switch input {
	case 1:
		go_decimal_converter.DecimalToBiner(int64(angka))
		break
	case 2:
		go_decimal_converter.DecimalToOctal(int64(angka))
		break
	case 3:
		go_decimal_converter.DecimalToOctal(int64(angka))
		break
	default:
		fmt.Println("Menu Not Available")
	}
}
