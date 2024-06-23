package main

import (
	"fmt"
	"sort"
)

func main() {
	var p *int
	fmt.Println("Value of p: ", p)

	anInt:=42
	p=&anInt
	fmt.Println("Value of p: ", *p)

	val1:=42.13
	pointer1:=&val1
	fmt.Println("Value1: ", *pointer1)

	*pointer1 = *pointer1 /30
	fmt.Println("Pointer1: ", *pointer1)
	fmt.Println("Value1: ", val1)

	var colors [3]string
	colors[0] = "red"
	colors[1] = "green"
	colors[2] = "blue"
	fmt.Println(colors)

	var numbers =[3]int{3,2,1}
	fmt.Println(numbers)
	fmt.Println(len(numbers))

	//slice
	var sliceNumbers =[]int{3,2,1}
	sliceNumbers = append(sliceNumbers, 6)
	fmt.Println(sliceNumbers)

	var sliceColors =[]string{"red", "green", "blue"}
	sliceColors = append(sliceColors[1:len(sliceColors)])
	fmt.Println(sliceColors)

	sliceColors = sliceColors[1:]
	fmt.Println(sliceColors)

	sliceColors = sliceColors[:len(sliceColors)-1]
	fmt.Println(sliceColors)

	nums:= make([]int, 5,5)
	nums[0] = 134
	nums[1] = 21
	nums[2] = 98
	fmt.Println(nums)

	nums2:= make([]int, 0)
	nums2 = append(nums2, 2)
	nums2 = append(nums2, 4)
	nums2 = append(nums2, 6)
	fmt.Println(nums2)

	sort.Ints(sliceNumbers)
	fmt.Println(sliceNumbers)

	states:=make(map[string]string)
	fmt.Println(states)
	states["WA"] = "Washington"
	states["OR"] = "Oregon"
	states["CA"] = "California"
	fmt.Println(states)
	california:=states["CA"]
	fmt.Println(california)
	delete(states, "OR")
	fmt.Println(states)
    states["NY"] = "New York"
	fmt.Println(states)

	for k,v:= range states{
		fmt.Printf("%v: %v\n", k,v)
	}

	keys:=make([]string, len(states)) 
	i:=0
	for k:=range states{
		keys[i]=k
		i++
	}
	fmt.Println(keys)
	sort.Strings(keys)
	fmt.Println(keys)

	for i:=range keys{
		fmt.Println(states[keys[i]])
	}

	poodle:=Dog{"Poodle", 10}
	fmt.Println(poodle)
	fmt.Printf("%+v\n",poodle)
	fmt.Printf("The dogs breed is %v\nThe dogs weight is %v\n",poodle.Breed, poodle.Weight)
	poodle.Weight = 9
	fmt.Printf("The dogs breed is %v\nThe dogs weight is %v\n",poodle.Breed, poodle.Weight)


}

type Dog struct {
	Breed string
	Weight int

}
