package strcmtd

import "math"

// menyiapkan struct bangun datar lingkaran,
// struct ini memiliki method yang beberapa diantaranya sudah terdefinisi di interface hitung.
type Lingkaran struct {
	Diameter float64
}

// keliling implements intf.Hitung.
func (Lingkaran) keliling() float64 {
	panic("unimplemented")
}

// luas implements intf.Hitung.
func (Lingkaran) luas() float64 {
	panic("unimplemented")
}

// menyiapkan method struct lingkaran
// untuk melakukan perhitungan jari jari
func (l Lingkaran) Jarijari() float64 {
	return l.Diameter / 2
}

// menyiapkan method struct lingkaran
// untuk melakukan perhitungan luas lingkaran
func (l Lingkaran) Luas() float64 {
	return math.Pi * math.Pow(l.Jarijari(), 2)
}

// menyiapkan method struct lingkaran
// untuk melakukan perhitungan keliling
func (l Lingkaran) Keliling() float64 {
	return math.Pi * l.Diameter
}
