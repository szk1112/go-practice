package main

import "fmt"

func main(){
	println("Hello world.")
	arr,err := getArray()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(arr[0], arr[1]) //=> GoLange Java
	fmt.Println(arr) //=> [GoLange Java]
}
func getArray() ([2]string, error) {

	arr := [...] string{"Golang", "Java"}
	return arr, nil

}