package main

import (
	"fmt"
	"strconv"
)

const standarKKM int = 75
func checkNilai(s Siswa) string {
	if s.nilai < standarKKM {
		return "Tidak Lulus"
	}else {
		return "Lulus"
	}
}
func checkHobi(s Siswa) string {
	if hobiCount := len(s.Hobi); hobiCount > 1 {
		return "dan memiliki " + strconv.Itoa(hobiCount) + " Hobi"
	}else {
		return "dan tidak memiliki hobi"
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
	fmt.Println(s1.Nama, checkNilai(s1))
	// fmt.Println(len(s1.Hobi))
	fmt.Println(checkHobi(s1))
}