package main

import (
	"fmt"
)
func print(text string){
	fmt.Println(text)
}
func basicLoop() {
	for i := 0; i < 5; i++ {
		fmt.Println("Pengulangan ke", i)
	}
}
func infinityLoop(){
	angka := 1
	for {
		fmt.Println("Angka", angka)
		if angka == 3 {
			fmt.Println("Selesai ....")
			break
		}
		angka++
	}
}
func segitigaRataKiri() {
	panjang := 5
	s := ""
	for i := 0; i < panjang; i++ {
		s += "*"
		print(s)
	}
}
func segitigaRataKiriTerbalik() {
	panjang := 5
	s := ""
	for i := 0; i < panjang; i++ {
		for j := panjang; j > i; j--{
			s += "*"
		}
		print(s)
		s = ""
	}
}
func segitigaRataKanan() {
	panjang := 5
	for i := 0; i < panjang; i++ {
		s := ""
		for j := 0; j < panjang - i - 1; j++ {
			s += " "
		}
		for k := 0; k <= i; k++ {
			s += "*"
		}
		print(s)
	}
}
func segitigaRataKananTerbalik() {
	panjang := 5
	for i := 0; i < panjang; i++ {
		s := ""
		for j := panjang; j > i; j-- {
			s += "*"
		}
		print(s)
	}
}
func main() {
	basicLoop()
	infinityLoop()
	segitigaRataKiri()
	segitigaRataKiriTerbalik()
	segitigaRataKanan()
	segitigaRataKananTerbalik()
}