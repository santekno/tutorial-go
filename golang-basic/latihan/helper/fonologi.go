package helper

func cekFonologi(input string) bool {
	var x [5]string
	x[0] = "a"
	x[1] = "i"
	x[2] = "u"
	x[3] = "e"
	x[4] = "o"

	for i := 0; i < 5; i++ { //Looping array
		if input == x[i] { //Perbandingan input dgn array
			return true
		}

	}
	return false
}
