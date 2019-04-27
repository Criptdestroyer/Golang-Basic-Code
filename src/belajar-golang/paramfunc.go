package main
import "fmt"
import "strings"

type FilterCallback func(string) bool
func param(){
	var data = []string{"Ahmad", "Emir", "Alfatah", "owo"}
	var dataContainsO = filter(data, func(each string)bool{
		return strings.Contains(each, "o")
	})
	var dataLenght5 = filter(data, func(each string)bool{
		return len(each) == 5
	})
	fmt.Println("data asli \t\t:", data)
    fmt.Println("filter ada huruf \"o\"\t:", dataContainsO)
	fmt.Println("filter jumlah huruf \"5\"\t:", dataLenght5)
	
	

}

	// func filter(data []string, callback FilterCallback) []string {
	// 	// ...
	// }

func filter(data[]string, callback func(string)bool)[]string{
	var result []string
	for _, each := range data{
		if filtered:=callback(each);filtered{
			result = append(result, each)
		}
	}
	return result
}