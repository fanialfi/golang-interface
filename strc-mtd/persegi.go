package strcmtd

import "math"

// menyiapkan struct bangun datar persegi,
// struct ini memiliki method yang beberapa diantaranya sudah terdefinisi di interface hitung
type Persegi struct {
	Sisi float64
}

// menyiapkan method struct persegi
// untuk melakukan perhitungan luas
func (p Persegi) Luas() float64 {
	return math.Pow(p.Sisi, 2)
}

// menyiapkan method struct persegi
// untuk melakukan perhitungan keliling
func (p Persegi) Keliling() float64 {
	return p.Sisi * 4
}
