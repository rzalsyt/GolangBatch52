package main

import (
	"fmt"
)

func main() {
	// deklarasi array
	fruits := [4]string{
		"apple",
		"grape",
		"banana",
		"melon",
	}

	fmt.Println(fruits)

	//array tampa jumlah elemt [..]
	var numbers = [...]int{2, 3, 2, 4, 3}

	fmt.Println("data array \t:", numbers)
	fmt.Println("jumlah elemen \t:", len(numbers))

	var numbers2 = [2][3]int{{3, 2, 3}, {3, 4, 5}}
	// buat variable = yang ada 2 elemt dan isinya {3,2,3} , {3,4,5)}
	fmt.Println("numbers2", numbers2)

	//Slice
	// Slice adalah reference elemen array. Slice bisa dibuat,
	//  atau bisa juga dihasilkan dari manipulasi sebuah array ataupun slice lainnya. Karena merupakan data reference,
	//  menjadikan perubahan data di tiap elemen slice akan berdampak pada slice lain yang memiliki alamat memori yang sama.

	//COntoh membuat slice, tidak perlu mendefinisikan jumlah elemetnya

	var fruits2 = []string{"apple", "grape", "banana", "melon"}
	//untuk mereferance index0 sampai 2
	var newFruits2 = fruits2[0:2]

	fmt.Println(newFruits2) // "apple"

}
