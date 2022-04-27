// aliase adalah sebuah fitur sebagai nama alternative dari tipe data yang sudah ada
// untuk aliase memiliki tipe data yang sama tetapi nama yang berbeda
// seperti byte dari tipe data uint8
// rune dari tipe data uint32

package main

import "fmt"

func main() {
	// deklarasi variable dengan tipe data uint8
	/*	var a uint8 = 10
		var b byte // byte adalah aliase dari tipe data uint8

		b = a // no error, karena byte memiliki tipe data yang sama dengan uint8
		_ = b
	*/

	// deklarasi tipe data aliase bernama second
	// type nama_aliase = nama_tipe_data

	type second = uint
	//tujuan dari aliase ini adalah seolah olah membuat tipe data yang baru  akan tetapi fungsinya tetap sama

	var hour second = 3600
	fmt.Printf("Hour type: %T\n", hour) // => hour type: uint
}
