package main

import "fmt"

func main() {

	// penggunaan conditions dengan temporary variable
	// penggunaan kondisional if else

	// var msg1 = "Kamu belum boleh membuat kartu sim"
	// var msg2 = "Kamu sudah boleh membuat kartu sim"

	// var currentYear = 2022

	// if age := currentYear - 1997; age < 18 {  // age merupakan sebuah variable yang terbuat dari scope block kondisional if else nya . (;) merupakan tanda semicolon
	// 	fmt.Println(msg1)
	// } else{
	// 	fmt.Println(msg2)
	// }


	// penggunaan kondisional switch

	// var score = 9

	// switch score{
	// case 9:
	// 	fmt.Println("Perfecto")
	// case 7:
	// 	fmt.Println("Good job")
	// case 6:
	// 	fmt.Println("Great")
	// default:
	// 	fmt.Println("Not bad")
	// }


	//penggunaan kondisional switch relational operators
	/*
	var scorer = 4

	switch{
	case scorer >= 8:
		fmt.Println("PERFEK BANGET")
	case (scorer < 8) && (scorer >4):
		fmt.Println("Lumayan")
		fallthrough					// fungsi fallthrough adalah  ketika case sudah terpenuhi dia akan tetap melakukan pengecekan case berikutnya
	case scorer < 5:
		fmt.Println("It's Okay , Belajar lebih rajin lagi")
	default:
		{
			fmt.Println("Belajar giat lagi")
			fmt.Println("Kamu masih punya waktu")
		}
	}
	*/

	// nested conditions : kondisional bersarang yang maksudnya mampu menambahkan semua kondisional didalamnya , baik itu if else ataupun switch keduanya.


	var skor = 0

	if skor > 7{
		switch skor{
		case 10:
			fmt.Println("Perfect")
		default:
			fmt.Println("Nice!!")
		}
	} else{
		if skor == 5 {
			fmt.Println("not bad")
		} else if skor == 3{
			fmt.Println("Tetap coba lagi")
		} else {
			fmt.Println("Kamu pasti bisa")
			if skor == 0{
				fmt.Println("Coba kerja keras lagi!!")
			}
		}
	}
}