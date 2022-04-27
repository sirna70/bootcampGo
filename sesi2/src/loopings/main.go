// first way looping

package main

import "fmt"

func main()  {

	//first way looping

	// for i := 0; i < 3; i++{
	// 	fmt.Println("Angka:", i)
	// }


	// second way looping : like as while looping
	// var i = 0

	// for i < 3{
	// 	fmt.Println("Angka", i)
	// 	i++
	// }
	
	// third way looping using keyword break
/*
	var i = 0

	for {
		fmt.Println(" ", i)


		i++

		if i == 5 {
			break
		}
	}
*/

// penggunaan break and continue keywords
/*
for i := 1; i <= 10; i++{
		if i%2 == 1{
			continue
		}

		if i > 8 {
			break
		}

		fmt.Println("Angka", i)
	}
*/

// nested looping : loopin bersarang yang memiliki suatu proses looping di dalamnya.
/*
for i := 0; i < 3; i++{
	for j := 0 ; j <= i; j++{
		fmt.Print("*")  // jika dibuat (" ", j) maka akan menampilkan angka 0 sampai 9
	}

	fmt.Println()
}
*/

/*
for i := 0; i < 5; i++{ // i adalah baris
	for j := i ; j < 5; j++{  // j adalah column
		fmt.Print("*")  // jika dibuat (" ", j) maka akan menampilkan angka 0 sampai 9
	}

	fmt.Println()
}
*/

// looping label : break and continue berlaku pada blok perulangan dimana ia digunakan saja.

outerloop:
	for i := 0; i < 3; i++{
		fmt.Println("Perulangan ke - ", i+1)
		for j := 0; j < 3; j++{
			if i == 2{
				break outerloop
			}
			fmt.Print(j, " ")
		}
		fmt.Println()
	}

}