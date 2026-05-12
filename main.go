package main

import "fmt"

func main() {
	data := NewData()
	data.Print()
	data.UpdateData()
	fmt.Print("\n")
	data.Print()
}
