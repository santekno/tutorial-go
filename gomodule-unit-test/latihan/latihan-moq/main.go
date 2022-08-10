package main

import "fmt"

type Student struct {
	FullName string `json:"fullname"`
	Grade    string `json:"grade"`
	Class    int    `json:"class"`
}

type StudentRepositoryInterface interface {
	GetAllStudents() ([]Student, error)
}

type StudentService struct {
	StudentRepositoryInterface
}

func (s StudentService) GetStudent() ([]Student, error) {
	Students, err := s.StudentRepositoryInterface.GetAllStudents()
	if err != nil {
		return nil, err
	}

	return Students, nil
}

type StudentRepository struct{}

func (r StudentRepository) GetAllStudents() ([]Student, error) {
	Students := []Student{
		{FullName: "Ihsan Arif", Grade: "B", Class: 1},
		{FullName: "Tono", Grade: "A", Class: 2},
		{FullName: "Andi", Grade: "C", Class: 3},
	}
	return Students, nil
}

func main() {
	repository := StudentRepository{}
	service := StudentService{repository}
	Students, _ := service.GetStudent()
	fmt.Println(Students)
}
