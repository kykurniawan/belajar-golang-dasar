package main

import "fmt"

func main() {
	var names [2]string

	names[0] = "Rizky"
	names[1] = "Kurniawan"

	fmt.Println(names[0])
	fmt.Println(names[1])

	var values = [3]int{
		90,
		70,
		85,
	}

	fmt.Println(values)

	fmt.Println(len(names))
	fmt.Println(len(values))
	
	var lagi [10]string
	
	fmt.Println(len(lagi))

}
