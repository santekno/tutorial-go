package math

func Average(xs []float64) float64 {
	total := float64(0)
	for _, x := range xs {
		total += x
	}
	return total / float64(len(xs))
}

func GanjilGenap(angka int) string {
	if angka%2 == 0 {
		return "genap"
	}
	return "ganjil"
}
