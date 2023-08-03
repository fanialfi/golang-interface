package main

import (
	"fmt"
	"golang/interface/intf"
	strcmtd "golang/interface/strc-mtd"
)

func main() {
	// inisialisasi variabel dengan menggunakan tipe data Hitung interface
	var bangunDatar intf.Hitung = strcmtd.Lingkaran{Diameter: 20.0}

	// menggunakan method yang sudah dibuat
	// implementasi lingkaran
	bangunDatar = strcmtd.Lingkaran{Diameter: 14.0}
	fmt.Println("=====> lingkaran")
	fmt.Printf("\n\tjari-jari\t: %.3f\n", bangunDatar.(strcmtd.Lingkaran).Jarijari())
	fmt.Printf("\tluas\t\t: %.3f\n", bangunDatar.Luas())
	fmt.Printf("\tkeliling\t: %.3f\n\n", bangunDatar.Keliling())

	// implementasi persegi
	bangunDatar = strcmtd.Persegi{Sisi: 10.0}
	fmt.Println("=====> persegi")
	fmt.Printf("\n\tluas\t\t: %.3f\n", bangunDatar.Luas())
	fmt.Printf("\tkeliling\t: %.3f\n", bangunDatar.Keliling())
}
