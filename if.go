package main

import "fmt"

func main() {
	var name = "Rizky"

	if name == "Rizky" {
		fmt.Println("Helo Rizky")
	} else if name == "Kurniawan" {
		fmt.Println("Hello Kurniawan")
	} else {
		fmt.Println("Hi, Kenalan dong")
	}

	if length := len(name); length > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}
