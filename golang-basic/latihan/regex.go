package training

import "fmt"

type Regex struct {
	Type  string `json:"type"`
	Regex string `json:"regex"`
}

//go:generate moq -out main_mock_test.go . RegexRepositoryInterface
type RegexRepositoryInterface interface {
	GetAllRegex() ([]Regex, error)
}

type RegexService struct {
	RegexRepositoryInterface
}

func (s RegexService) GetRegex() ([]Regex, error) {
	users, _ := s.RegexRepositoryInterface.GetAllRegex()
	return users, nil
}

type RegexRepository struct{}

func (r RegexRepository) GetAllRegex() ([]Regex, error) {
	users := []Regex{
		{"Vokal", "real"},
		{"Konsonan", "real2"},
	}
	return users, nil
}

func main() {
	repository := RegexRepository{}
	service := RegexService{repository}
	regex, _ := service.GetRegex()
	fmt.Println(regex)
}
