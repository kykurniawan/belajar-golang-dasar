package main

import "fmt"

func main() {
	type NoKTP string
	type Married bool

	var noKtpRizky NoKTP = "013421847698146"
	var marriedStatus Married = false
	fmt.Println(noKtpRizky)
	fmt.Println(marriedStatus)
}