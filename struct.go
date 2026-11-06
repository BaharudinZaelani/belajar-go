package main

import "fmt"

type Siswa struct {
	Nama string
	Umur int
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
	fmt.Println(dataSiswa[0])
	ubahUmur(&dataSiswa[0], 20)
	fmt.Println(dataSiswa)
	fmt.Println(len(dataSiswa))
}