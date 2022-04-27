// multiple variable declarations - 1
// can be used more variable on one line

/* package main

import "fmt"

func main() {

	var stud1, stud2, stud3 string = "satu", "dua", "tiga"
	var first, second, third = 1, 2, 3

	fmt.Println(stud1, stud2, stud3)

	fmt.Println(first, second, third)

}
*/

// multiple variable declarations - 2
//
// can be used more variable on one line

/* package main

import "fmt"

func main() {

	var name, age, address = "Ariell", 24, "Jl.Balige"
	first, second, third := "1", 2, "3"

	fmt.Println(name, age, address)

	fmt.Println(first, second, third)
	fmt.Printf("%T, %T, %T\n", first, second, third)

}
*/

// underscore variable -> we can't unusing 1 or more variable if declarate - 3
// but wwhen using underscore for variable , it's make compiling success

package main

import "fmt"

func main() {

	var firstVariable string

	var name, age, address = "Arokell", 24, "Jalan BALIGE"

	_, _, _, _ = firstVariable, name, age, address

	fmt.Printf("HAlo nama saya %s, umurku adalah %d, aku tinggal di %s\n", name, age, address)

}

// fungsi fmt.printf -> mengetahui tipe data suatu variable, struktur ouput yang dihasilkan akan bergantung pada flag yang dipakai.
// %s memiliki fungsi utk melakukan format suatu nilai dengan tipe data string
// %d digunakan untuk format tipe data int dengan base 10 atau angka 0 sampai 9.
