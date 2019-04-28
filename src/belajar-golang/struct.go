package main
import "fmt"
func struc(){
	//tipe 1
	var s1 student
	s1.name = "Ahmad Emir Alfatah"
	s1.grade = 1
	fmt.Println("name :",s1.name)
	fmt.Println("grade :",s1.grade)

	//tipe 2
	var s2 = student{}
	s2.name = "jeju"
	s2.grade = 1

	//tipe 3
	var s3 = student{"jeje", 1}
	var s4 = student{name: "ame"}
	fmt.Println("student 2 :", s2.name)
	fmt.Println("student 3 :", s3.name)
	fmt.Println("student 4 :", s4.name)

	//tipe 4
	var s5 *student = &s1
	fmt.Println("student 1, name :", s1.name)
	fmt.Println("student 5, name :", s5.name)
	s5.name = "ethan"
	fmt.Println("student 1, name :", s1.name)
	fmt.Println("student 5, name :", s5.name)

	//tipe 5
	var s6 = student2{}
	s6.name = "wickwick"
	s6.age = 21
	s6.grade =2
	s6.person.age = 22
	fmt.Println("name  :", s6.name)
    fmt.Println("age   :", s6.age)
    fmt.Println("age   :", s6.person.age)
	fmt.Println("grade :", s6.grade)
	
	//tipe 6
	var p1 = person{name: "alfath", age: 21}
	var s7 = student2{person: p1, grade: 2}
	fmt.Println("name  :", s7.name)
	fmt.Println("age   :", s7.age)
	fmt.Println("grade :", s7.grade)

	//tipe 7
	var s8 = struct{
		person
		grade int
	}{}
	s8.person = person{"wikwik",21}
	s8.grade = 2
	fmt.Println("name  :", s8.person.name)
    fmt.Println("age   :", s8.person.age)
    fmt.Println("grade :", s8.grade)
	var s9 = struct {
		person
		grade int
	}{
		person: person{"wick", 21},
		grade:  2,
	}
	fmt.Println("name  :", s9.person.name)
    fmt.Println("age   :", s9.person.age)
	fmt.Println("grade :", s9.grade)
	
	//tipe 8
	var allStudents = []person{
		{name:"emir", age:23},
		{name: "Ethan", age:23},
		{name: "bourne", age:22},
	}
	for _, student2 := range allStudents{
		fmt.Println(student2.name, "age is", student2.age)
	}

	//tipe 9
	var allStudents2 = []struct{
		person
		grade int
	}{
		{person: person{"wick",21}, grade:2},
		{person: person{"wick2",21}, grade:3},
		{person: person{"wick3",21}, grade:3},
	}
	for _, student := range allStudents2{
		fmt.Println(student)
	}

	//tipe 10
	var student3 struct{
		person
		grade int
	}
	student3.person = person{"wick", 21}
	student3.grade = 2

	//tipe 11
	// var p2 = struct{name string; age int}{age:22, name:"emir"}
	// var p3 = struct{name string;age int}{"hehe", 23}

	//tipe 12
	
	var p4 = person{"hehe", 21}
	fmt.Println(p4)
	// var p5 = People{"wick", 21}
	// fmt.Println(p5)
}
// type People = person

type person2 struct {name string; age int; hobbies []string}

type student4 struct{
	person struct{
		name string //`tag1`
		age int //`tag2`
	}
	grade int
	hobbies []string
}

type person struct{
	name string
	age int
}

type student2 struct{
	grade int
	// age int //bila sama
	person
}

type student struct{
	name string
	grade int
}