package main

import (
	"encoding/json"
	"fmt"
)

type Siswa struct {
	Nama string
	Umur int
	nilai int
	Hobi []string
	piket []string
	Laptop
}
func ubahUmur(s *Siswa, umurBaru int) {
	s.Umur = umurBaru
}
func ubahNilai(s *Siswa, nilaiBaru int) {
	s.nilai = nilaiBaru
}

func dataStruct() {
	dataSiswa := []Siswa{
		{
			Nama: "Budi",
			Umur: 16,
		},
	}
	// Mengubah data
	fmt.Println(dataSiswa[0])
	ubahUmur(&dataSiswa[0], 20)
	fmt.Println(dataSiswa)
	fmt.Println(len(dataSiswa))
}

func structToJson() {
	dataSiswa := Siswa{
		Nama: "Budi",
		Umur: 17,
		Hobi: []string{"Membaca", "Coding"},
	}
	// konversi menjadi json
	jsonData, err := json.Marshal(dataSiswa)
	if err != nil {
		fmt.Println("Gagal melakukan konversi", err)
		return
	}

	fmt.Println(string(jsonData))
}

func jsonToStruct() {
	jsonString := `{"nama": "Ani", "umur": 16}`
	jsonData := []byte(jsonString)
	var siswaAni Siswa

	err := json.Unmarshal(jsonData, &siswaAni)
	if err != nil {
		fmt.Println("Gagal mengambil unmarshal:", err)
	}

	fmt.Println(siswaAni)
}

// Embeding Struct
type Laptop struct {
	Merk string
	harga int
	Spesifikasi
}
type Spesifikasi struct {
	Ram int
	Storage int
}
func embedingStruct() {
	l1 := Laptop{
		Merk: "Asus",
		harga: 4000000,
		Spesifikasi: Spesifikasi{
			Ram: 8,
			Storage: 128,
		},
	}

	fmt.Println(l1)
}