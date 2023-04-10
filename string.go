package main

import "fmt"

func main() {
	var namaDepan = "Erwin"
	var namaBelakang string
	namaBelakang = "Rahayu"
	fmt.Println(namaDepan, namaBelakang)
	// count string
	fmt.Println("panjang karakter", len(namaDepan+namaBelakang))
}
