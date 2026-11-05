package main

import "fmt"

func array() {
	// Array
	// Deklarasi dengan ukuran (nill/zero value)
	var buah [3]string
	buah[0] = "Apple"
	buah[1] = "Jeruk"
	buah[2] = "Mangga"

	// Mengubah data
	// buah[0] = "Durian"

	// Deklarasi langsung
	// angka := [3]int{10,20,30}
	// fmt.Println(angka)

	// Deklarasi tanpa ukuran (otomatis dihitung untuk jumlahnya)
	// hobi := [...]string{"Coding", "Gaming", "Membaca"}
	// fmt.Println(hobi)
	// fmt.Println("jumlah hobi:", len(hobi))

	// Menyalin Array
	buahBaru := buah
	buahBaru[1] = "Durian"
	fmt.Println(buahBaru)
	fmt.Println(buah) // array baru tidak mengubah array lama

}

func slice() {
	// Slice
	angka := make([]int, 3, 5)
	fmt.Println(angka)
	fmt.Println(len(angka)) // panjang slice
	fmt.Println(cap(angka)) // kapasitas slice

	buah := []string{"Apple", "Jeruk", "Mangga"}
	fmt.Println(buah)
	fmt.Println(len(buah))
	fmt.Println(cap(buah))

	// Append / Menambah data
	buahBaru := append(buah, "Durian")
	fmt.Println(buahBaru)
	fmt.Println(len(buahBaru))
	fmt.Println(cap(buahBaru))

	// String juga array perlu di konversi menjadi string
	domisili := "Sukabumi"
	fmt.Println(string(domisili[0]))
	fmt.Println(len(domisili))
}

func mapData() {
	// Map
	nilaiUjian := map[string]int{
		"Budi": 85,
		"Ani": 90,
		"Joko": 78, // harus ada koma di akhir data
	}
	fmt.Println(nilaiUjian)
	// Map dengan make (function dari go untuk menginisialisasi data ke zero value)
	ibukota := make(map[string]string)
	ibukota["Indonesia"] = "Jakarta"
	fmt.Println(ibukota)
}

func commaIdiom() {
	// Comma-Ok Idiom
	nilai := map[string]int{
		"Bahar": 85,
	}
	
	skorBahar, adaBahar := nilai["Bahar"]
	fmt.Println(skorBahar, adaBahar)
	skorBudi, adaBudi := nilai["budi"]
	fmt.Println(skorBudi, adaBudi)
}