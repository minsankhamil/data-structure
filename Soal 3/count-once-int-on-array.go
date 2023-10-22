package main

import (
	"fmt"
	"strconv"
)

//Buat program sesuai dengan deskripsi di bawah. Input merupakan variable string berisi kumpulan angka.
//Output merupakan list / array berisi angka yang hanya muncul 1 kali pada input.

func munculSekali(angka string) []int {
	// Melacak berapa kali setiap angka muncul
	find := make(map[int]int)

	// Menghitung kemunculan setiap angka
	for _, number := range angka {
		intNumber, err := strconv.Atoi(string(number))
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return nil
		}
		find[intNumber]++
	}

	// Menyimpan angka yang hanya muncul satu kali
	result := []int{}

	// Memeriksa angka yang hanya muncul sekali
	for _, number := range angka {
		intNumber, _ := strconv.Atoi(string(number))
		if find[intNumber] == 1 {
			result = append(result, intNumber)
		}
	}

	return result
}

func main() {
	fmt.Println(munculSekali("1234123"))
	//[4]

	fmt.Println(munculSekali("76523752"))
	//[6 3]

	fmt.Println(munculSekali("12345"))
	//[1 2 3 4 5]

	fmt.Println(munculSekali("1122334455"))
	//[]

	fmt.Println(munculSekali("0872504"))
	//[8 7 2 5 4]
}
