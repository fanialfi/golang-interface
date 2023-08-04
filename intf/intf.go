package intf

// membuat interface exported, yang berisi 2 method untuk melakukan perhitungan bangun datar
// type Hitung interface {
// 	Luas() float64
// 	Keliling() float64
// }

// melakukan embeded interface

// berisi method untuk kalkulasi luas dan keliling
type hitung2d interface {
	Luas() float64
	Keliling() float64
}

// berisi method untuk mencari volume bidang
type hitung3d interface {
	Volume() float64
}

// meng embed intercase Hitung2d dan Hitung3d
type Hitung interface {
	hitung2d
	hitung3d
}
