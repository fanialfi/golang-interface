package main

import (
	"fmt"
	"golang/interface/intf"
	strcmtd "golang/interface/strc-mtd"
)

// membuat implementasi perhitungan di atas
func main() {
	var bangunDatar intf.Hitung = strcmtd.Lingkaran{Diameter: 14.0}

	// implementasi perhitungan bangun datar persegi
	bangunDatar = strcmtd.Lingkaran{Diameter: 14.0}
	fmt.Println("===== lingkaran")
	fmt.Printf("luas\t\t: %f\n", bangunDatar.Luas())
	fmt.Printf("keliling\t: %f\n", bangunDatar.Keliling())

	// untuk mengakses method yang tidak terdefinisi di interface, variabel-nya harus di-casting terlebih dahulu ke tipe asli variabel konkrit-nya (pada kasus ini tipe-nya lingkaran)
	// setelah itu method dapat di akses
	// cara casting interface sedikit unik, yaitu dengan menuliskan nama tipe tujuan dalam kurung.
	// ditempatkan setelah nama interface, lalu diikuti dengan notasi dot (seperti cara mengakses field pada sebuah struct, hanya saja ada tanda kurungnya)
	fmt.Printf("jari jari\t: %v\n", bangunDatar.(strcmtd.Lingkaran).Jarijari())

	// implementasi perhitungan bangun datar lingkaran
	bangunDatar = strcmtd.Persegi{Sisi: 10.0}

	fmt.Println("===== persegi")
	fmt.Printf("luas\t\t: %f\n", bangunDatar.Luas())
	fmt.Printf("keliling\t: %f\n", bangunDatar.Keliling())
}
