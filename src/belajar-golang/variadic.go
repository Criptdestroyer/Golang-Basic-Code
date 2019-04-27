package main
import (
	"fmt"
	"strings"
)
func variadic(){
	var avg = calculate2(2, 4, 3, 5, 4, 3, 3, 5, 5, 3)
	var msg = fmt.Sprintf("Rata-rata : %.2f", avg) //mengembalikan nilai dalam bentuk string (fmt.Sprint() dan fmt.Sprintln())
	fmt.Println(msg)
	var numbers = []int{2,4,3,5,4,3,3,5,5,3}
	avg = calculate2(numbers...)
	msg = fmt.Sprintf("Rata-rata : %.2f", avg)
	fmt.Println(msg)
	var hobbies = []string{"sleeping", "eating"}
	yourHobbies("emir","code", "bucin")
	yourHobbies("emir",hobbies...)
}

//fungsi variadic
func yourHobbies(name string, hobbies ...string){
	var hobbiesAsString = strings.Join(hobbies, ", ")
	fmt.Printf("Hello, my name is: %s\n", name)
    fmt.Printf("My hobbies are: %s\n", hobbiesAsString)
}

func calculate2(numbers ...int)float64{
	var total int = 0
	for _, number := range numbers{
		total+=number
	}
	var avg = float64(total)/float64(len(numbers))
	return avg
}