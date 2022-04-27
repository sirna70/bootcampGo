package main

import (
	"fmt"
)

func main() {

	// slice with backing array
	// metode ini sangat lebih hemat memori dibanding menggunakan array
	// means fruits22 melakukan slicing terhadap fruits11 untuk mendapatkan element dari index ke- 1 sampai ke-3
	// rambutan berada pada index 1 di fruits11 ketika ditampilkan, karena masih dalam satu backing array
	/*	var fruits11 = []string{"apple", "orange", "mango", "durian", "Potatoes"}

		var fruits22 = fruits11[2:4]

		fruits22[0] = "rambutan"

		fmt.Println("fruits11 =>", fruits11)
		fmt.Println("fruits22 =>", fruits22)
	*/

	//slice with cap function
	// untuk mengetahui kapasitas dari sebuah array maupun slice
	// cap berfungsi untuk melakukan slicing terhadap array tanpa mengurangi apapun
	// len berfungsi untuk melakukan slicing terhadap data array dengan mengurangi 1 data dari value index inputnya
	/*
		var fruits = []string{"apple", "manggoes", "banana", "durian"}

		fmt.Println("Fruits1 cap:", cap(fruits)) //4
		fmt.Println("Fruits1 len:", len(fruits)) //4

		fmt.Println(strings.Repeat("#", 20))

		var fruits2 = fruits[0:3]

		fmt.Println("fruits2 cap:", cap(fruits2)) //4
		fmt.Println("fruits2 len:", len(fruits2)) //3

		fmt.Println(strings.Repeat("#", 20))

		var fruits3 = fruits[1:]

		fmt.Println("Fruits3 cap:", cap(fruits3)) //3
		fmt.Println("Fruits3 len:", len(fruits3)) //3
	*/

	// creating a new backing array
	// when kita ingin mendapatkan element2 dari slice yg sudah ada. namun ingin membuat backing array yang baru,
	//maka dapat menggunakan fungsi append untuk melakukannya

	cars := []string{"Hyundai", "Honda", "Audi", "Range Rover"}
	newCars := []string{}

	newCars = append(newCars, cars[0:2]...)
	// newcars akan melakukan slicing terhadap data array yang ada pada cars dengan index yang telah ditentukan

	cars[0] = "Nissan"
	fmt.Println("cars:", cars)
	fmt.Println("newCars:", newCars)

}
