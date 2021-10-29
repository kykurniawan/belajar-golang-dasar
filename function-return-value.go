package main

import "fmt"

func getHello(name string) string {
	if name == "" {
		return "Hello bro"
	} else {
		return "Hello " + name
	}
}

func main() {
	fmt.Println(getHello(""))
	fmt.Println(getHello("Rizky Kurniawan"))
}
