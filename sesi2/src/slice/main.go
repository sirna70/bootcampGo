// slice mirip dengan array : menyimpan data satu atau lebih
// slice tidak memiliki fixed-length berarti panjang slice tidak tetap
// slice termasuk reference type

// slice make function
package main

import (
	"fmt"
	"strings"
)

func main() {

	//penggunaan make function

	// var fruits = make([]string,0) // tidak perlu mengisi value / panjangnya disini [] karena slice membuat kita leluasa dalam menentukan panjangnya
	// dilakukan string, 0 karena agar tidak menambah variabel baru di tampilan ketika bersamaan digunakan dengan append function

	// _ = fruits

	// fruits[0] = "Potatoes"
	// fruits[1] = "Banana"

	//penggunaan append function
	// fungsinya adalah mengembalikan nilai dari slice yang ditambahkannya, maka harus disimpan ke dalam suatu variable
	//	fruits = append(fruits, "apple", "banana","orange")

	// append function with ellipsis
	/*	var fruits = []string{"durian", "jagung"}

		var fruits2 = []string{"jeruk", "mangga"}

		fruits = append(fruits, fruits2...) // tanda elipsis adalah ... digunakan setelah variabel tanpa spasi

		fmt.Printf("%#v\n", fruits)
	*/

	//slice with copy function
	// maksudnya mengembalikan jumlah elemen yang berhasil ter-copy.
	// variable1 := copy (fruits1, fruits2) => maksudnya fruits1 suda ter replace oleh element yang ada di fruits2
	/*
		var fruits1 = []string{"apple", "banana", "strawberry"}
		var fruits2 = []string{"Potatoes", "Lengkuas", "Advokat"}

		nm := copy(fruits1, fruits2)

		fmt.Println("fruits1 =>", fruits1)
		fmt.Println("fruits2 =>", fruits2)
		fmt.Println("Copied elements =>", nm)
	*/

	//slicing {menggunakan metode start;stop}
	// ketika ingin menampilkan data dalam array,
	// misal kita mau dari index 0 - 2, maka cukup ketik 0:3 atau :3 , karena metode slicing adalah memotong angka inputannya atau value input nya tidak dihitung.

	var fruits = []string{"Mango", "Dragon", "Orange", "Pineaple", "Biwa"}

	var fruits1 = fruits[1:4]
	fmt.Printf("%#v\n", fruits1)

	var fruits2 = fruits[0:]
	fmt.Printf("%#v\n", fruits2)

	var fruits3 = fruits[:]
	fmt.Printf("%#v\n", fruits3)

	var fruits4 = fruits[:3]
	fmt.Printf("%#v\n", fruits4)

	// combining with slicing and append

	fmt.Println(strings.Repeat("#", 60))

	fruits = append(fruits[:3], "Durian", "Anggur")
	fmt.Printf("%#v\n", fruits)

}
