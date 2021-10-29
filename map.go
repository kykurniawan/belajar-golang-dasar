package main

import "fmt"

func main() {
	var person map[string]string = map[string]string{
		"name": "Rizky",
		"address":"Subur Indah",
	}

	person["title"] = "Programmer Keren"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	var book map[string]string = make(map[string]string)

	book["title"] = "Belajar Go Lang"
	book["author"] = "Rizky Kurniawan"
	book["ups"] = "Salah"

	fmt.Println(book)
	fmt.Println(len(book))
	
	delete(book, "ups")
	
	fmt.Println(book)
	fmt.Println(len(book))
}
