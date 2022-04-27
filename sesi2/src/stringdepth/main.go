// strings in depth

package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {

	// dalam case ini kita menampilkan byte dari pada karakter jakarta
	// dengan menggunakan angka desimal ascii code
	/*	city := "Jakarta"

		for i := 0; i < len(city); i++ {
			fmt.Printf("Index: %d, byte: %d\n", i, city[i])
		}
	*/

	// selanjutnya mengubah kode ascii menjadi sebuah karakter
	// konsep ini disebut dengan slice of byte

	// tipe data string pada Go merupakan kumpulan byte yang berada didalam slice/slice of bytes

	/*	var city []byte = []byte{74, 97, 107, 97, 114, 116, 97}

		fmt.Println(string(city))
		// agar slice of byte dapat mencetak nilai stringnya, maka perlu diletakkan kedalam func.string(nama variabel)
	*/

	// rune by rune , maksudnya ketika ingin mendapatkan panjang karakter
	// function len() sebenarnya tidak mendapatkan panjang dari string berdasarkan karakternya, tetapi dari jumlah byte pada string.
	str1 := "makan"
	str2 := "mânca" // â merupakan accented-character bukan termasuk ascii code sehingga memiliki 2 byte

	fmt.Printf("str1 byte length => %d\n", len(str1))
	fmt.Printf("str2 byte length => %d\n", len(str2))

	// penggunaan rune by rune
	fmt.Printf("str1 charac. length => %d\n", utf8.RuneCountInString(str1))
	fmt.Printf("str2 charac. length => %d\n", utf8.RuneCountInString(str2))

	//penggunaan range loop terhadap rune by rune
	//tujuannya memilah atau menjabarkan karakter tersebut berdasarkan index dan byte daripada string
	fmt.Println(strings.Repeat("#", 25))
	for i, s := range str2 {
		fmt.Printf("Index => %d, string => %s\n", i, string(s))
	}
}
