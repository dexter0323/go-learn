package main

import "fmt"

func main() {
	data := []byte("Hello World!")
	dataString := string(data)
	dataByte := []byte(dataString)
	fmt.Println(data, dataString, dataByte)
}
