package intf

// membuat interface exported, yang berisi 2 method untuk melakukan perhitungan bangun datar
type Hitung interface {
	Luas() float64
	Keliling() float64
}
