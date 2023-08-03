package strcmtd

import "math"

// membuat struct Lingkaran
type Lingkaran struct {
	Diameter float64
}

// method untuk melakukan perhitungan mencari jari jari sebuah lingkaran
func (l Lingkaran) Jarijari() float64 {
	return l.Diameter / 2
}

// method untuk melakukan perhitungan luas sebuah lingkaran
func (l Lingkaran) Luas() float64 {
	return math.Pi * math.Pow(l.Jarijari(), 2)
}

// method untuk melakukan perhitungan keliling sebuah lingkaran
func (l Lingkaran) Keliling() float64 {
	return math.Pi * l.Diameter
}
