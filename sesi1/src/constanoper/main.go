package main

import "fmt"

func main() {

	const fullName string = "Ariell navy" // ketika menggunakan const , wajib langsung mengisi value dari si const variable tersebut agar tidak error

	fmt.Printf("Hello %s\n", fullName)

	// penggunaan operator aritmatika

	var value = 2 + 2*5      // Golang akan membaca kode program dari kiri ke kanan, akan tetapi untuk case perkalian.. akan diutamakan proses perkalian lalu di ikuti lainnya.
	var value2 = (2 + 2) * 5 // hati hati saat penggunaan buka kurung
	fmt.Println(value)
	fmt.Println(value2)

	//penggunaan operator relational / perbandingan

	var pertama bool = 3 < 10
	var kedua bool = "Rick" == "Ricky"
	var ketiga bool = 10 != 5.5
	var keempat bool = 11 <= 7

	fmt.Println("pertama adalah: ", pertama)
	fmt.Println("kedua adalah: ", kedua)
	fmt.Println("ketiga adalah: ", ketiga)
	fmt.Println("keempat adalah: ", keempat)

	//operator logika

	var right = true
	var left = false

	var and = right && left
	fmt.Printf("resultnya adalah \t(%t)\n", and)

	var or = right || left
	fmt.Printf("resultnya adalah \t(%t)\n", or)

	var notsame = !left
	fmt.Printf("resultnya adalah \t(%t)\n", notsame)

}
