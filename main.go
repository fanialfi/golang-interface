package main

import (
	"fmt"
	"golang/interface/intf"
	strcmtd "golang/interface/strc-mtd"
)

// func main() {
// 	// inisialisasi variabel dengan menggunakan tipe data Hitung interface
// 	var bangunDatar intf.Hitung = strcmtd.Lingkaran{Diameter: 20.0}

// 	// menggunakan method yang sudah dibuat
// 	// implementasi lingkaran
// 	bangunDatar = strcmtd.Lingkaran{Diameter: 14.0}
// 	fmt.Println("=====> lingkaran")

// 	// karena method Jarijari() tidak tersedia di interface intf.Hitung, maka untuk mengakses-nya harus di-casting terlebih dahulu ke tipe variabel asli nya (pada kasus ini bertpe strcmtd.Lingkaran)
// 	// setelah itu method nya baru bisa di akses
// 	// untuk melakukan casting interface yaitu dengan menuliskan nama tipe tujuan dalam tanda kurung
// 	// ditempatkan setelah nama interface dengan menggunakan notasi titik (seperti cara megakses field pada sebuah struct, hanya saja ada tanda kurungnya)
// 	fmt.Printf("\n\tjari-jari\t: %.3f\n", bangunDatar.(strcmtd.Lingkaran).Jarijari())
// 	fmt.Printf("\tluas\t\t: %.3f\n", bangunDatar.Luas())
// 	fmt.Printf("\tkeliling\t: %.3f\n\n", bangunDatar.Keliling())

// 	// implementasi persegi
// 	bangunDatar = strcmtd.Persegi{Sisi: 10.0}
// 	fmt.Println("=====> persegi")
// 	fmt.Printf("\n\tluas\t\t: %.3f\n", bangunDatar.Luas())
// 	fmt.Printf("\tkeliling\t: %.3f\n", bangunDatar.Keliling())
// }

// implementasi embeded interface yang telah dibuat
func main() {
	// variabel yang dibuat menggunakan struct yang memiliki method pointer,
	// jika di tampung dalam variabel interface harus diambil referensinya terlebih dahulu
	var bangunRuang intf.Hitung = &strcmtd.Kubus{Sisi: 4}

	fmt.Println("=====> Kubus")
	fmt.Println("luas 		: ", bangunRuang.Luas())
	fmt.Println("keliling 	: ", bangunRuang.Keliling())
	fmt.Println("volume 		: ", bangunRuang.Volume())
}
