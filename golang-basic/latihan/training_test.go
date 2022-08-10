package training

import (
	"testing"

	"github.com/stretchr/testify/mock"
)

func TestHello(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "test 1",
			want: "Hello, world.",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Hello(); got != tt.want {
				t.Errorf("Hello() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsPolindrome(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test 1",
			args: args{
				word: "malam",
			},
			want: true,
		},
		{
			name: "test 2",
			args: args{
				word: "hutan",
			},
			want: false,
		},
		{
			name: "test 3",
			args: args{
				word: "mama",
			},
			want: false,
		},
		{
			name: "test 4",
			args: args{
				word: "ibu",
			},
			want: false,
		},
		{
			name: "test 5",
			args: args{
				word: "noon",
			},
			want: true,
		},
		{
			name: "test 6",
			args: args{
				word: "",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPolindrome(tt.args.word); got != tt.want {
				t.Errorf("IsPolindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

type RegexRepositoryMock struct {
	mock.Mock
}

func (r RegexRepositoryMock) GetAllRegex() ([]Regex, error) {
	args := r.Called()
	regex := []Regex{
		{"Vokal", "[${aiueo}]"},
	}
	return regex, args.Error(1)
}

func TestTypeChar(t *testing.T) {
	type args struct {
		char string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "test 1",
			args: args{
				char: "!",
			},
			wantErr: true,
		},
		{
			name: "test 2",
			args: args{
				char: "!",
			},
			wantErr: true,
		},
		{
			name: "test 3",
			args: args{
				char: "asd",
			},
			wantErr: true,
		},
		{
			name: "test 4",
			args: args{
				char: "t",
			},
			want: "Konsonan",
		},
		{
			name: "test 5",
			args: args{
				char: "T",
			},
			want: "Konsonan",
		},
		{
			name: "test 6",
			args: args{
				char: "a",
			},
			want: "Vokal",
		},
	}
	for _, tt := range tests {
		repository := RegexRepositoryMock{}
		repository.On("GetAllRegex").Return([]Regex{}, nil)

		service := RegexService{repository}
		regex, _ := service.GetRegex()

		t.Run(tt.name, func(t *testing.T) {
			got, err := TypeChar(tt.args.char, regex)
			if (err != nil) != tt.wantErr {
				t.Errorf("TypeChar() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("TypeChar() = %v, want %v", got, tt.want)
			}
		})
	}
}
