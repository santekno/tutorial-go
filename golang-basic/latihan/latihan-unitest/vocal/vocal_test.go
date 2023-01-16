package vocal

import "testing"

func TestVocal(t *testing.T) {
	type args struct {
		letter string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "should be vocal",
			args: args{
				letter: "A",
			},
			want: "vocal",
		},
		{
			name: "should be consonant",
			args: args{
				letter: "B",
			},
			want: "consonant",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Vocal(tt.args.letter); got != tt.want {
				t.Errorf("Vocal() = %v, want %v", got, tt.want)
			}
		})
	}
}
