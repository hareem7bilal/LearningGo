package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"net/http"
	"encoding/json"
	"strings"

)

const url = "http://services.explorecalifornia.org/json/tours.php"

func main() {
	content:="Hello from Go!"
	file, err:=os.Create("./fromString.txt")
	checkError(err)
	len, err := io.WriteString(file, content)
	checkError(err)
	fmt.Printf("Wrote a file with %v characters\n", len)
	defer file.Close()
	defer readFile("./fromString.txt")

	res, err:=http.Get(url)
	checkError(err)

	fmt.Printf("Response type: %T\n", res)
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
        fmt.Printf("Error: received status code %d\n", res.StatusCode)
        return
    }
	
	bytes, err:=ioutil.ReadAll(res.Body)
	checkError(err)
	data:=string(bytes)
	fmt.Print(data)

	tours:= toursFromJSON(data)
	for _, tour:=range tours{
		fmt.Println(tour.Name)
	}

}

func checkError(err error){
	if err!=nil{
		panic(err)
	}
}

func readFile(fileName string){
	data, err:= ioutil.ReadFile(fileName)
	checkError(err)
	fmt.Println("Text read from file:",string(data))
}

func toursFromJSON(content string) []Tour{
	tours:=make([]Tour, 0, 20)
	decoder:=json.NewDecoder(strings.NewReader(content))
	_, err:=decoder.Token()
	checkError(err)

	var tour Tour
	for decoder.More(){
		err:=decoder.Decode(&tour)
		checkError(err)
		tours = append(tours, tour)
	}
	return tours
}

type Tour struct {
	Name, Price string
}
