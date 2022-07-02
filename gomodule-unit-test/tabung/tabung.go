package math

import "math"

type Tabung struct {
	Jarijari, Tinggi float64
}

func (t Tabung) Volume() float64 {
	return math.Phi * math.Pow(t.Jarijari, 2) * t.Tinggi
}

func (t Tabung) Luas() float64 {
	return 2 * math.Phi * t.Jarijari * (t.Jarijari + t.Tinggi)
}

func (t Tabung) KelilingAlas() float64 {
	return 2 * math.Phi * t.Jarijari
}
