package strcmtd

import "math"

// membuat struct Persegi
type Persegi struct {
	Sisi float64
}

// method untuk melakukan perhitungan luas persegi
func (p Persegi) Luas() float64 {
	return math.Pow(p.Sisi, 2)
}

// method untuk melakukan perhitungan keliling persegi
func (p Persegi) Keliling() float64 {
	return p.Sisi * 4
}
