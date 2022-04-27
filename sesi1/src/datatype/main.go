package main

import "fmt"

func main() {

	var first uint8 = 89  // untuk menghemat memori , karena value yg di input berada di interval value uint8
	var second int8 = -12 // untuk menghemat memori , karena value yang di input berada di interval value int8
	var decimal float32 = 3.80

	fmt.Printf("tipe data pertama adalah %T\n", first) // %T : untuk menampilkan jenis tipe data ; %d : untuk menampilkan value int ; %f , %3f , %2f : untuk menampilkan value float baik itu 3 dibelakang koma, atau seterusnya
	fmt.Printf("tipe data kedua adalah %T\n", second)
	fmt.Printf("tipe data ketiga adalah %T\n", decimal)

	//penggunaan boolean
	var condition bool = true
	fmt.Printf("Is the permitted? %t \n", condition) // fungsi dari %t ialah untuk menampilkan value boolean apakah bernilai true atau false

	//penggunaan string
	var message string = "hallo gais\n"
	var msg2 = `Semuanya baik baik saja kan "aku sangat merindukan kalian 
	jangan lupa untuk menjaga pola makan".`
	fmt.Print(message) // string dalam go lang sangat unik , karena backtiks/grave accent (``) bisa digunakan sebagai string dan lainnya.
	fmt.Println(msg2)
}
