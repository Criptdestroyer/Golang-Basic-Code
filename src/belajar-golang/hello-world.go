package main

import (
	"fmt"
	"strings"
	"math/rand"
	"time"
	"math"
)

func main(){
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

	/*
	var diameter float64 = 15
	var area, circumference = calculate(diameter)
	fmt.Printf("luas lingkaran\t\t: %.2f \n",area)
	fmt.Printf("keliling lingkaran\t: %.2f \n", circumference)
	*/

	/*
	divideNumber(10,2)
	divideNumber(4,0)
	divideNumber(8, -4)
	*/

	rand.Seed(time.Now().Unix())//membuat random benar2 acak
	/*
	var randomValue int
	randomValue = randomWithRange(2,10)
	fmt.Println("random number:", randomValue)
	randomValue = randomWithRange(2,10)
	fmt.Println("random number:", randomValue)
	randomValue = randomWithRange(2,10)
	fmt.Println("random number:", randomValue)
	// func nameOfFunc(paramA type, paramB type, paramC type) returnType
	// func nameOfFunc(paramA, paramB, paramC type) returnType
	// func randomWithRange(min int, max int) int
	// func randomWithRange(min, max int) int
	*/

	/*
	var names = []string{"john","wick"}
	printMessage("halo", names)
	*/

	/*
	fmt.Println("hello world");
	varia()
	ifstatement()
	perulangan()
	array()
	slice()
	maps()
	*/
}

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

func calculate(d float64) (area float64,circumference float64){
	area = math.Pi * math.Pow(d/2,2)
	circumference = math.Pi*d
	return //area, circumference
}

func divideNumber(m,n int){
	if n == 0{
		fmt.Printf("invalid divider. %d cannot divide by %d\n", m, n)
		return
	}
	var res = m/n
	fmt.Printf("%d / %d = %d\n", m, n, res)
}

func randomWithRange(min, max int) int{
	var value = rand.Int()%(max-min+1)+min
	return value
}

func printMessage(message string, arr[]string){
	var nameString = strings.Join(arr, " ")
	fmt.Println(message, nameString)
}