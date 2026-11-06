package main

import (
	"encoding/json"
	"fmt"
)

type Siswa struct {
	Nama string
	Umur int
	Hobi []string `json:"hobi"`
}
func ubahUmur(s *Siswa, umurBaru int) {
	s.Umur = umurBaru
}


func main() {
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

	// konversi menjadi json
	jsonData, err := json.Marshal(dataSiswa)
	if err != nil {
		fmt.Println("Gagal melakukan konversi", err)
		return
	}

	fmt.Println(string(jsonData))
}