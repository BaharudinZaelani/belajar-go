package main

import "fmt"

type Siswa struct {
	Nama string
	Umur int
}
func main() {
	dataSiswa := []Siswa{
		{
			Nama: "Budi",
			Umur: 16,
		},
	}

	fmt.Println(dataSiswa)
	fmt.Println(len(dataSiswa))
}