package main

import "fmt"

func main() {
	name := "Kurniawan"

	switch name {
	case "Rizky":
		fmt.Println("Hello Rizky")

	case "Kurniawan":
		fmt.Println("Hello Kurniawan")

	default:
		fmt.Println("Hi, Boleh kenalan?")
	}

	// switch length := len(name); length > 5 {
	// case true:
	// 	fmt.Println("Nama terlalu panjang")
	// case false:
	// 	fmt.Println("Nama sudah benar")
	// }

	length := len(name)

	switch {
	case length > 10:
		fmt.Println("Nama terlalu panjang")
	case length > 10:
		fmt.Println("Nama lumayan panjang")
	default:
		fmt.Println("Nama sudah benar")
	}
}
