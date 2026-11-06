package main

import (
	"fmt"
	"strconv"
)

const standarKKM int = 75
const hariSekarang string = "Senin"
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
func checkHari() {
	switch hariSekarang {
    case "Senin":
        fmt.Println("Semangat awal pekan!")
    case "Rabu":
        fmt.Println("Sudah separuh jalan menuju akhir pekan!")
    case "Jumat":
        fmt.Println("Besok sudah akhir pekan!")
    default:
        fmt.Println("Hari biasa, tetap semangat!")
    }
}

func conditionals() {
	s1 := Siswa{
		Nama: "Ani",
		Umur: 20,
		nilai: 80,
		Hobi: []string{"Gaming", "Coding"},
		piket: []string{
			"Senin",
			"Sabtu",
		},
		Laptop: Laptop{
			Merk: "Asus",
			harga: 4000000,
			Spesifikasi: Spesifikasi{
				Ram: 8,
				Storage: 500,
			},
		},
	}
	ubahNilai(&s1, 74)
	fmt.Println(s1.Nama, checkNilai(s1))
	fmt.Println(checkHobi(s1))
	fmt.Println("Data Lengkap", s1)
	checkHari()
}