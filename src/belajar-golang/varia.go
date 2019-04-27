package main
import "fmt"

func varia(){
	// var firstName string = "Emir"
	// var firstName = "Emir"
	// lastName := "Alfatah"
	var firstName, lastName = "Emir", "Alfatah"
	var _ = "hehe"
	name := new(string) 
	exist := true
	message := `Nama saya "Ahmad Emir Alfatah".
	Teknik Informatika.
	Universitas Sriwijaya.`
	const phi = 3.14 //tidak bisa diubah

	fmt.Printf("Halo %s %s!\n", firstName, lastName)
	fmt.Println("Halo", firstName, lastName + "!")
	fmt.Println(name)
	fmt.Println(*name)
	fmt.Printf("exits? %t \n", exist)
	fmt.Println(message)
}