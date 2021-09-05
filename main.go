package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"time"
)

func main(){
	println("Hello world.")
	arr,err := getArray()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(arr[0], arr[1]) //=> GoLange Java
	fmt.Println(arr) //=> [GoLange Java]
	err = output(arr)
	if err != nil {
		return
	}

}
func getArray() ([]string, error) {

	arr := []string{"1", "Java"}
	return arr, nil

}
func output(data []string) error {
	now := time.Now()
	path := "./output"
	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		fmt.Println(err)
		return err
	}
	f, err := os.Create(fmt.Sprintf("./output/test%02d%02d.txt", now.Month(), now.Day()))
	if err != nil {
		return err
	}
	bw := bufio.NewWriter(f)
	header := [][]string{
		{
			"id",
			"name",
		},
	}
	content := append(header, data)

	w := csv.NewWriter(bw)
	err = w.WriteAll(content)
	if err != nil {
		return err
	}
	w.Flush()

	return nil
}