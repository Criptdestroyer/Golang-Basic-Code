package main
import "fmt"

func maps(){
	//tipe 1
	var chicken map[string]int
	chicken = map[string]int{}
	chicken["januari"]=50
	chicken["februari"]=40
	fmt.Println("januari", chicken["januari"])
	fmt.Println("mei", chicken["mei"])

	//tipe 2
	var chicken1 = map[string]int{"januari":50, "februari": 40}
	chicken = map[string]int{
		"januari":50,
		"februari":40,
		"maret":34,
    	"april":67,
	}
	// var chicken3 = map[string]int{}
	// var chicken4 = make(map[string]int)
	// var chicken5 = *new(map[string]int)
	for key, val := range chicken{
		fmt.Println(key,"  \t:",val)
	}

	//delete()
	fmt.Println(len(chicken1))
	fmt.Println(chicken1)
	delete(chicken1, "januari")
	fmt.Println(len(chicken1))
	fmt.Println(chicken1)

	//isexist
	var value, isExist = chicken1["mei"]
	if isExist{
		fmt.Println(value)
	}else{
		fmt.Println("item is not exist")
	}

	//map and slice
	var chickens = []map[string]string{
		/*map[string]string*/{"name":"chicken blue", "gender":"male", "color":"blue"},
		/*map[string]string*/{"name":"chicken red", "gender":"male"},
		/*map[string]string*/{"name":"chicken yellow", "gender":"female"},
	}
	for _, chicken := range chickens{
		fmt.Println(chicken["gender"], chicken["name"])
	}
}