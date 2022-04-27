//tipe data 1

/* package main

import "fmt"

func main() {

	var name string = "ARIELLIN"
	var age int = 24

	fmt.Println("Ini adalah namanya ==>", name)
	fmt.Println("Ini adalah umurnya ==>", age)
}
*/

// tipe data 2

package main

import "fmt"

func main() {

	var name string
	var age int = 24

	name = "ARIELL Anjelo"
	age = 24

	// tipe data 3 , akan terjadi error jika declare name and age dibawahnya tidak sesuai dengan tipe data yg ditentukan diawal.
	//	name = 30
	//	age = "ARIOLLIK"

	fmt.Println("Ini adalah namaku", name)
	fmt.Println("ini adalah umurku", age)

}
