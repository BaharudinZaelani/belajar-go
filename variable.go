package main

import "fmt"

// NOTE : Semua variable harus digunakan jika di bahasa GO

func variable() {
	// Eksplisit
	// var nama string = "Baharudin Zaelani"
	// var umur int8 = 25
	// var apakahAktif bool = true

	// Short Declaration / Implisit
	// nama := "Baharudin Zaelani"
	// umur := 25
	// apakahAktif := true

	// Block Declaration
	// var (
	// 	nama 				= "Baharudin Zaelani"
	// 	umur				= 25
	// 	apakahAktif	= true
	// )
	// var a,b,c,d int8 = 1,2,3,4
	// fmt.Println(nama, umur, apakahAktif)
	// fmt.Println(a,b,c,d)

	// Konstanta
	// const namaAplikasi string = "BaharBelajar Go"
	// const pi float32 = 3.14
	// fmt.Println(namaAplikasi, pi)

	// Zero Values
	// var nama string // string kosong
	// var umur int8 // 0
	// var status bool // false
	// fmt.Println(nama, umur, status)

	// IOTA
	// Setiap blok variable memiliki identifier dimulai dari 0 khusus untuk const
	const (
		senin = iota
		selasa
		rabu
		kamis
		jumat
		sabtu
		minggu
	)
	fmt.Println("senin", senin)
	fmt.Println("selasa", selasa)
	fmt.Println("rabu", rabu)
	fmt.Println("kamis", kamis)
	fmt.Println("jumat", jumat)
	fmt.Println("sabtu", sabtu)
	fmt.Println("minggu", minggu)
}