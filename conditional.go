package main

import (
	"fmt"
)

const standarKKM int = 75
func checkNilai(s *Siswa) string {
	if s.nilai < standarKKM {
		return "Tidak Lulus"
	}else {
		return "Lulus"
	}
}

func main() {
	s1 := Siswa{
		Nama: "Ani",
		Umur: 20,
		nilai: 80,
		Hobi: []string{"Gaming", "Coding"},
	}
	// ubahNilai(&s1, 74)
	fmt.Println(s1.Nama, checkNilai(&s1))
}