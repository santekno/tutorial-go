package konsonOrVocal

import "testing"

func BenchmarkKonsonanOrVocal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		konsonanOrVocal("a")
	}
}

func BenchmarkKonsonanOrVocalTableTest(b *testing.B) {
	type args struct {
		kalimat string
	}
	tests := []struct {
		name       string
		parameters args
	}{
		{
			name: "test number 1",
			parameters: args{
				kalimat: "a",
			},
		},
		{
			name: "test number 2",
			parameters: args{
				kalimat: "b",
			},
		},
	}
	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				konsonanOrVocal(tt.parameters.kalimat)
			}
		})
	}
}

func Test_konsonanOrVocal(t *testing.T) {
	type args struct {
		huruf string
	}
	//parameter := huruf
	tests := []struct {
		name       string
		parameters args
		want       string
	}{
		{
			name: "test number 1",
			parameters: args{
				huruf: "a",
			},
			want: "Vocal",
		},
		{
			name: "test number 2",
			parameters: args{
				huruf: "b",
			},
			want: "Konsonan",
		},
		{
			name: "test number 3",
			parameters: args{
				huruf: "yygfuyhv",
			},
			want: "Hanya bisa memasukan 1 huruf bukan kata atau kalimat",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := konsonanOrVocal(tt.parameters.huruf); got != tt.want {
				t.Errorf("konsonanOrVocal() = %v, want %v", got, tt.want)
			}
		})
	}
}
