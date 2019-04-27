package main
import "fmt"

func array(){
	//tipe 1
	var names[3]string
	names[0] = "Ahmad"
	names[1] = "Emir"
	names[2] = "Alfatah"
	// names[3] = "hehe" error
	fmt.Println(names[0], names[1], names[2])

	//tipe 2
	var fruits = [4]string{"apple","grape","banana","melon"}
	fmt.Println("Jumlah element \t\t", len(fruits))
	fmt.Println("Isi semua element \t", fruits)

	//tipe 3
	// var fruits [4]string
	fruits = [4]string{
		"apple",
		"grape",
		"banana",
		"melon",
	}

	//tipe 4
	var numbers = [...]int{2,3,2,4,3}
	fmt.Println("data array \t:",numbers)
	fmt.Println("jumlah elemen \t:",len(numbers))

	//tipe 5
	var numbers1 = [2][3]int{[3]int{3,2,3}, [3]int{3,4,5}}
	var numbers2 = [2][3]int{{3,2,3},{3,4,5}}
	fmt.Println("numbers1", numbers1)
	fmt.Println("numbers2", numbers2)

	//tipe 6
	//var fruits = [4]string{"apple", "grape", "banana", "melon"}
	for i:=0;i<len(fruits);i++{
		fmt.Printf("elemen %d : %s\n", i, fruits[i])
	}
	for i, fruit := range fruits{
		fmt.Printf("elemen %d : %s\n", i, fruit)
	}
	for _, fruit := range fruits{
		fmt.Printf("nama buah : %s\n", fruit)
	}
	// for i, _ := range fruits { }
	// for i := range fruits { }

	//tipe 7
	var fruits2 = make([]string, 2)
	fruits2[0] = "apple"
	fruits2[1] = "manggo"
	fmt.Println(fruits2)
}