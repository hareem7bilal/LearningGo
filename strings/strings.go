package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
)

const aConst = 64

func main() {
	var aString string = "This is a go!"
	fmt.Println(aString)
	fmt.Printf("The variables type is a %T\n", aString)
	var anInteger int = 42
	fmt.Printf("The variables type is a %T\n", anInteger)
	var defaultInteger int
	fmt.Println(defaultInteger)
	var anotherString = "This is another string"
	fmt.Printf("The variables type is a %T\n", anotherString)
	var anotherInteger = 42
	fmt.Printf("The variables type is a %T\n", anotherInteger)
	myString := "This is also a string"
	fmt.Println(myString)
	fmt.Println(aConst)
	fmt.Printf("The variables type is a %T\n", aConst)

	var reader = bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	input, _ := reader.ReadString('\n')
	fmt.Println("You entered ", input)

	fmt.Print("Enter a number: ")
	numInput, _ := reader.ReadString('\n')
	aFloat, err:= strconv.ParseFloat(strings.TrimSpace(numInput), 64)
	if err != nil {
		fmt.Println(err)
	}else{
		fmt.Println("Value of number: ", aFloat)
	}

}
