package main

import (
	"fmt"
	"math"
)

// mempersiapkan interface dan mendefinisikan method,
// pendefinisian interface menggunakan keyword type dan interface
// variabel yang menggunakan tipe data interface hitung, ketika diisi menggunakan struct, minimal harus mempunyai method yang terdefinisikan di interface hitung
type hitung interface {
	luas() float64
	keliling() float64
	// jarijari() float64
}

// menyiapkan struct bangun datar lingkaran,
// struct ini memiliki method yang beberapa diantaranya sudah terdefinisi di interface hitung.
type lingkaran struct {
	diameter float64
}

// menyiapkan method struct lingkaran
// untuk melakukan perhitungan jari jari
func (l lingkaran) jarijari() float64 {
	return l.diameter / 2
}

// menyiapkan method struct lingkaran
// untuk melakukan perhitungan luas lingkaran
func (l lingkaran) luas() float64 {
	return math.Pi * math.Pow(l.jarijari(), 2)
}

// menyiapkan method struct lingkaran
// untuk melakukan perhitungan keliling
func (l lingkaran) keliling() float64 {
	return math.Pi * l.diameter
}

// menyiapkan struct bangun datar lingkaran,
// struct ini memiliki method yang beberapa diantaranya sudah terdefinisi di interface hitung
type persegi struct {
	sisi float64
}

// menyiapkan method struct persegi
// untuk melakukan perhitungan luas
func (p persegi) luas() float64 {
	return math.Pow(p.sisi, 2)
}

// menyiapkan method struct persegi
// untuk melakukan perhitungan keliling
func (p persegi) keliling() float64 {
	return p.sisi * 4
}

// membuat implementasi perhitungan di atas
func main() {
	var bangunDatar hitung = lingkaran{diameter: 14.0}

	// implementasi perhitungan bangun datar persegi
	bangunDatar = lingkaran{diameter: 14.0}
	fmt.Println("===== lingkaran")
	fmt.Printf("luas\t\t: %f\n", bangunDatar.luas())
	fmt.Printf("keliling\t: %f\n", bangunDatar.keliling())

	// untuk mengakses method yang tidak terdefinisi di interface, variabel-nya harus di-casting terlebih dahulu ke tipe asli variabel konkrit-nya (pada kasus ini tipe-nya lingkaran)
	// setelah itu method dapat di akses
	// cara casting interface sedikit unik, yaitu dengan menuliskan nama tipe tujuan dalam kurung.
	// ditempatkan setelah nama interface, lalu diikuti dengan notasi dot (seperti cara mengakses field pada sebuah struct, hanya saja ada tanda kurungnya)
	fmt.Printf("jari jari\t: %v\n", bangunDatar.(lingkaran).jarijari())

	// implementasi perhitungan bangun datar lingkaran
	bangunDatar = persegi{sisi: 10.0}

	fmt.Println("===== persegi")
	fmt.Printf("luas\t\t: %f\n", bangunDatar.luas())
	fmt.Printf("keliling\t: %f\n", bangunDatar.keliling())
}
