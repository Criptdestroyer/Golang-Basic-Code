package main

import (
	"fmt"
	"strings"
	"math/rand"
	"time"
	"math"
)

func main(){
	var diameter float64 = 15
	var area, circumference = calculate(diameter)
	fmt.Printf("luas lingkaran\t\t: %.2f \n",area)
	fmt.Printf("keliling lingkaran\t: %.2f \n", circumference)

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