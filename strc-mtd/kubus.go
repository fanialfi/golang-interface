package strcmtd

import "math"

// menyiapkan struct yang berisi method Luas() Keliling() dan Volume()
type Kubus struct {
	Sisi float64
}

// method untuk melakukan perhitungan Luas sebuah kubus
func (k *Kubus) Luas() float64 {
	return math.Pow(k.Sisi, 2) * 6
}

// method untuk melakukan perhitungan Volume sebuah kubus
func (k *Kubus) Volume() float64 {
	return math.Pow(k.Sisi, 3)
}

// method untuk melakukan perhitungan Keliling sebuah kubus
func (k Kubus) Keliling() float64 {
	return k.Sisi * 12
}

/*
	variabel hasil dari struct Kubus nantinya akan ditampung oleh variabel bertipe data interface Hitung
	yang isinya merupakan gabungan dari Hitung2d dan Hitung3d

	sebelumnya sudah tau bahwa method pointer bisa diakses lewat variabel biasa dan variabel pointer
*/
