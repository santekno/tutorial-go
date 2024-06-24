package test


import "testing"

func Test_checkConsonant(t *testing.T) {
	type args struct {
		ch string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test untuk konsonan",
			args: args{
				ch: "k",
			},
			want: "Konsonan",
		},
		{
			name: "Test untuk konsonan 2",
			args: args{
				ch: "p",
			},
			want: "Konsonan",
		},
		{
			name: "Test untuk Vokal a",
			args: args{
				ch: "a",
			},
			want: "Vokal",
		},
		{
			name: "Test untuk Vokal i",
			args: args{
				ch: "i",
			},
			want: "Vokal",
		},
		{
			name: "Test untuk Vokal u",
			args: args{
				ch: "u",
			},
			want: "Vokal",
		},
		{
			name: "Test untuk Vokal e",
			args: args{
				ch: "e",
			},
			want: "Vokal",
		},
		{
			name: "Test untuk Vokal o",
			args: args{
				ch: "o",
			},
			want: "Vokal",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkConsonant(tt.args.ch); got != tt.want {
				t.Errorf("checkConsonant() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkConsonantWithPackage(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test untuk konsonan",
			args: args{
				input: "k",
			},
			want: "Konsonan",
		},
		{
			name: "Test untuk konsonan 2",
			args: args{
				input: "p",
			},
			want: "Konsonan",
		},
		{
			name: "Test untuk Vokal a",
			args: args{
				input: "a",
			},
			want: "Vokal",
		},
		{
			name: "Test untuk Vokal i",
			args: args{
				input: "i",
			},
			want: "Vokal",
		},
		{
			name: "Test untuk Vokal u",
			args: args{
				input: "u",
			},
			want: "Vokal",
		},
		{
			name: "Test untuk Vokal e",
			args: args{
				input: "e",
			},
			want: "Vokal",
		},
		{
			name: "Test untuk Vokal o",
			args: args{
				input: "o",
			},
			want: "Vokal",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkConsonantWithPackage(tt.args.input); got != tt.want {
				t.Errorf("checkConsonantWithPackage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkCheckConsonant(b *testing.B) {
	for i := 0; i < b.N; i++ {
		checkConsonant("a")
	}
}

func BenchmarkCheckConsonantWithPackage(b *testing.B) {
	for i := 0; i < b.N; i++ {
		checkConsonantWithPackage("a")
	}
}