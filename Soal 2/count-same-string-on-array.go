package main

import "fmt"

//buatlah sebuah program yang dapat menghitung berapa banyak sebuah string yang sama didalam sebuah slice!

func Mapping(slice []string) map[string]int {
	//your code here

	// Menyimpan jumlah kemunculan setiap string
	Count := make(map[string]int)

	// Menghitung kemunculan setiap string
	for _, string := range slice {
		Count[string]++
	}

	return Count
}

func main() {
	fmt.Println(Mapping([]string{"asd", "qwe", "asd", "adi", "qwe", "qwe"}))
	//map[adi:1 asd:2 qwe:3]

	fmt.Println(Mapping([]string{"asd", "qwe", "asd"}))
	//map[asd:2 qwe:1]

	fmt.Println(Mapping([]string{}))
	//map[]
}
