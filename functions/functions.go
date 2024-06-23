package main

import (
	"fmt"
)

func main() {
	doSomething()
	sum:= addValues(5,8)
	fmt.Println(sum)
	multiSum, multiCount:= addAllValues(5,8,4,9)
	fmt.Println("Sum of multiple values: ", multiSum)
	fmt.Println("Count of multiple values: ", multiCount)

	poodle:=Dog{"Poodle", 10, "Woof"}
	poodle.Speak()
	poodle.Sound = "Arf"
	poodle.Speak()
	poodle.SpeakThreeTimes()
}

func doSomething(){
	fmt.Println("Doing Something")
}

func addValues(val1, val2 int) int {
	return val1+val2
}

func addAllValues(vals ...int) (int, int) {
	total:=0
	for _, val := range vals{
		total+=val
	}
	return total, len(vals)
}

type Dog struct {
	Breed string
	Weight int
	Sound string
}

func (d Dog) Speak(){
	fmt.Println(d.Sound)
}

func (d Dog) SpeakThreeTimes(){
	d.Sound = fmt.Sprintf("%v %v %v",d.Sound,d.Sound,d.Sound)
	fmt.Println(d.Sound)
}
