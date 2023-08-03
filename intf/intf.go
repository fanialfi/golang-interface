package intf

// mempersiapkan interface dan mendefinisikan method,
// pendefinisian interface menggunakan keyword type dan interface
// variabel yang menggunakan tipe data interface hitung, ketika diisi menggunakan struct, minimal harus mempunyai method yang terdefinisikan di interface hitung
type Hitung interface {
	Luas() float64
	Keliling() float64
}
