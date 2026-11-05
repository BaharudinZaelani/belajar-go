package main

import "fmt"

/*
Tipe data integer
int 	: default data type 32bit/64bit sesuai arsitektur komputer
int8	: -128 sampai 127
int16	: -2.147.483.648 sampai 2.147.483.647
int64 : -9.223.372.036.854.775.808 sampai 9.223.372.036.854.775.807

Unsigned Integer hanya positive value
uint, uint8, uint16, uint32, uint64, serta uintptr

Float
float32 : menggunakan memory 32-bit
float64 : menggunakan memory 64-bit

Complex
complex64 : Terdiri dari dua buah angka float32 (32-bit untuk bagian riil dan 32-bit untuk bagian imajiner, total 64-bit).
complex128: Terdiri dari dua buah angka float64 (64-bit untuk bagian riil dan 64-bit untuk bagian imajiner, total 128-bit). Ini adalah tipe default jika Anda mendeklarasikan bilangan kompleks secara langsung.

Alias
byte : untuk uint8
rune : untuk int32

Boolean = bool
Text/String = string

*/
func dataType() {
	/*
	// Mendeklarasikan bilangan kompleks dengan tipe complex128 (default)
	var c1 complex128 = 5 + 3i
	
	// Menggunakan type inference (:=)
	c2 := 2.5 + 1.5i

	fmt.Println("Bilangan Kompleks 1:", c1)
	fmt.Println("Bilangan Kompleks 2:", c2)

	// Operasi aritmatika dasar
	tambah := c1 + c2
	kali := c1 * c2

	fmt.Println("Hasil Penjumlahan:", tambah)
	fmt.Println("Hasil Perkalian:", kali)

	stringBiasa := "\tHallo namaku bahar\nsedang belajar \"Golang\"."
	rawString := `Hallo aku bahar
	Aku sedang belajar GO
	baris ini otomatis ada di bawah \n`

	fmt.Println("--- String Biasa ---")
	fmt.Println(stringBiasa)

	fmt.Println("\n--- Raw String Literal ---")
	fmt.Println(rawString)

	fmt.Println("\n--- Contoh penggunaan untuk RAW biasanya untuk sql query ---")
	query := `SELECT id, nama, email 
    FROM users 
    WHERE status = 'active' 
    ORDER BY created_at DESC;
	`
	fmt.Println(query)
	*/

	// Type Conversion
	var angkaInt int = 10
	var angkaFloat float64 = float64(angkaInt)

	var desimal float64 = 9.99
	var dibulatkan int = int(desimal) // Pembulatan mendekati angka 0 atau kebawah (truncate)

	fmt.Println("Int:", angkaInt)
	fmt.Println("Float:", angkaFloat)
	fmt.Println("Desimal asli:", desimal)
	fmt.Println("Setelah dikonversi ke int:", dibulatkan)
}