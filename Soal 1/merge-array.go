package main

import (
	"fmt"
)

//Buatlah sebuah program menggabungkan 2 array yang diberikan,
//dan jangan sampai terdapat nama yang sama di data yang sudah tergabung tadi!

func ArrayMerge(arrayA, arrayB []string) []string {
	//your code here

	// Melacak nama yang sudah ada
	find := make(map[string]bool)

	// Menyimpan hasil penggabungan array
	result := []string{}

	// Menggabungkan arrayA
	for _, unique := range arrayA {
		if !find[unique] {
			find[unique] = true
			result = append(result, unique)
		}
	}

	// Menggabungkan arrayB
	for _, unique := range arrayB {
		if !find[unique] {
			find[unique] = true
			result = append(result, unique)
		}
	}

	return result
}

func main() {
	//test cases
	fmt.Println(ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))
	//["king", "devil jin", "akuma", "eddie", "steve", "geese"]

	fmt.Println(ArrayMerge([]string{"sergei", "jin"}, []string{"jin", "steve", "bryan"}))
	//["sergei", "jin", "steve", "steve"]

	fmt.Println(ArrayMerge([]string{"alisa", "yoshimitsu"}, []string{"devil jin", "yoshimitsu", "alisa", "law"}))
	//["alisa", "yoshimitsu", "devil jin", "law"]

	fmt.Println(ArrayMerge([]string{}, []string{"devil jin", "sergei"}))
	//["devil jin", "sergei"]

	fmt.Println(ArrayMerge([]string{"hwoarang"}, []string{}))
	//["hwoarang"]

	fmt.Println(ArrayMerge([]string{}, []string{}))
	//[]
}
