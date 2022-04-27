// array using... ,
// %#v memiliki fungsi untuk melakukan format tipe data array agar dapat melihat panjang dari arraynya.

package main

import (
	"fmt"
	// "strings"   // only using for string, cause golang so strict
)

func main() {

	// array biasa
	/*
		var numbers [4]int

		numbers = [4]int{1,2,3,4} // data yang dimasukkan harus sama jumlahnya dengan length yang di imput agar tidak terjadi pengosongan alias nilai 0

		var strings = [3]string{"Androl","Defok","Napitu"}

		fmt.Printf("%#v\n", numbers)
		fmt.Printf("%#v\n", strings)
	*/

	// array menggunakan modify element through index
	// maksudnya dapat memodifikasi elemet" yang berada didalam array
	// sehingga ketika data yang telah di isi pertama kali dimodifikasi menjadi yang lain maka yang akan tampil ketika di compile ialah yang dimodifikasi tersebut
	/*
		var fruits = [3]string{"apel","jeruk","pisang"}
		fruits[0] = "apple"
		fruits[1] = "orange"
		fruits[2] = "banana"

		fmt.Printf("%#v\n", fruits)
	*/

	// penggunaan range loop dan loop biasa menggunakan len
	/*
		var fruits = [3]string{"apple","banana","mango"}

		//cara pertama menggunakan range loop dan lebih mudah digunakan
		for i,v := range fruits{
			fmt.Printf("Index: %d, Value: %s\n",i,v)
		}

		fmt.Println(strings.Repeat("#", 25)) // 25 merupakan value, yang maksudnya # akan ditampilkan sebanyak 25 kali

		//cara kedua looping biasa dengan dibantu fungsi len agar mendapatkan panjang dari array len<nama array>
		for i := 0; i < len(fruits); i++{
			fmt.Printf("Index: %d, Value: %s\n",i,fruits[i])
		}
	*/

	// multidimensional Array
	//maksudnya ialah terdapat array didalam array

	balances := [2][3]int{{5, 6, 7}, {8, 9, 10}} // maksudnya ialah memiliki 2 baris , 3 kolom sehingga datanya akan ditampilkan sesuai baris dan kolom yang di input

	for _, arr := range balances {
		for _, value := range arr {
			fmt.Printf("%d ", value)
		}
		fmt.Println()
	}

}
