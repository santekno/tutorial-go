package helper

func IsPalindrome(input string) bool {
	// cek palindrome dengan melakukan looping, kemudian men cocokkan ASCII dari index awal dan index akhir
	// pertama lakukan looping sejumlah panjang isi variable input, untuk mengambil index input yang dimulai dari awal (disini hasilnya variable i)
	// kedua buat variable j yang di dapat dari jumlah panjang input dan dikurangi 1(karena index mulai dari 0), kemudian kurangi juga dengan variable i
	// hal ini untuk mencocokkan indexnya, contoh panjang input adalah 10 maka saat i berada di index pertama yaitu 0 maka j akan di index terakhir yaitu 9
	// j = 10 - 1 - 0, hasilnya 9. Dan jika i berada di index 1 maka j = 10 - 1 - 1 = 8
	// jadi i adalah representasi index yang dimulai dari awal ke akhir dan j adalah index dari akhir ke awal
	// kemudian langkah ke tiga cocokkan ASCII kedua index, jika terdapat 1 perbadaan dalam looping maka langkung print false dan return
	// dan karena kita hanya perlu mencocokkan setengah dari tiap sisi saja maka lakukan pengkondisian dengan mengecek jika i lebih dari setengah maka lakukan break untuk menghentikan perulangan
	// terakhir jika perulangan selesai tanpa ada kesalahan, maka return true

	for i := 0; i < len(input); i++ {
		j := len(input) - 1 - i
		if input[i] != input[j] {
			return false
		}

		if i > (len(input)/2)+1 {
			break
		}
	}

	return true
}
