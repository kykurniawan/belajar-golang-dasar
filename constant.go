package main

import "fmt"

func main() {
	const firstName string = "Rizky"
	const lastName = "Kurniawan"
	const value = 1000

	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(value)

	const (
		rrs string = "RRS"
	)

	fmt.Println(rrs)
}
