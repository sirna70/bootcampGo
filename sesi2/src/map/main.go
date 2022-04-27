// map sama dengan konsep array dan slice yaitu menyimpan satu atau lebih data namun, map disimpan
// sebagai key:value pairs(pasang key dan value)
// zero value dari tipe data map adalah nil
// map salah satu tipe data yang masuk kategori reference type = slice type data
// jika map mengcopy map lainnya, dan map tsb mengganti value pada suatu key, maka value dari map lainnya akan terganti juga

package main

import (
	"fmt"
)

func main() {

	// cara pertama
	/*
		var person map[string]string // Deklarasi

		person = map[string]string{} //inisialisasi

		person["name"] = "Rakitic"

		person["age"] = "23"

		person["address"] = "Jalan suparman"
	*/

	// cara kedua
	/*	var person = map[string]string{
			"name":    "Ariell",
			"age":     "23",
			"address": "Jalan Wahidin",
		}

		fmt.Println("name:", person["name"])
		fmt.Println("age:", person["age"])
		fmt.Println("address:", person["address"])
	*/

	// looping with map
	/*
		var person = map[string]string{
			"name":    "Ariel",
			"age":     "23",
			"address": "Jalan andusman",
		}

		for key, value := range person {
			fmt.Println(key, ":", value)
		}
	*/

	// deleting a value
	/*	var person = map[string]string{
			"name":    "Airel",
			"age":     "25",
			"address": "Jalan Agung",
		}

		fmt.Println("Before deleting:", person)

		delete(person, "name")

		fmt.Println("After deleting:", person)
	*/

	// detecting a value
	// hanya mendeteksi keberadaan suatu value saja
	// variabel yg diberikan akan mengembalikan value asli dari mapnya jika memang ada, jika tidak ada maka kita akan mendapat zero value dari valuenya
	// variabel exist yang kita berikan akan mengembalikan tipe data boolean
	/*
		var person = map[string]string{
			"name":    "Swatiastu",
			"age":     "25",
			"address": "Jalan waerpot",
		}

		value, exist := person["age"]

		if exist {
			fmt.Println(value)
		} else {
			fmt.Println("Value doesn't exist")
		}

		delete(person, "age")

		value, exist = person["age"]

		if exist {
			fmt.Println(value)
		} else {
			fmt.Println("Value has been deleted")

		}
	*/

	// combining slice with map
	var people = []map[string]string{
		{"name": "Ariel", "age": "22"},
		{"name": "Jade", "age": "25"},
		{"name": "coolek", "age": "21"},
	}

	for i, person := range people {
		fmt.Printf("Index: %d, name: %s, age: %s\n", i, person["name"], person["age"])
	}
}
