package main

import (
	"errors"
	"fmt"
)

type Student struct {
	FullName string `json:"fullname"`
	Grade    string `json:"grade"`
	Class    int    `json:"class"`
}

//go:generate moq -out main_mock_test.go . StudentRepositoryInterface
type StudentRepositoryInterface interface {
	GetAllStudents() ([]Student, error)
}

type StudentService struct {
	StudentRepositoryInterface
}

func (s StudentService) GetStudent(name string) ([]Student, error) {
	Students, err := s.StudentRepositoryInterface.GetAllStudents()
	if err != nil {
		return nil, err
	}

	var result = make([]Student, 0)

	for i := range Students {
		if Students[i].FullName == name {
			result = append(result, Students[i])
		}
	}

	if len(result) < 1 {
		return nil, errors.New("Data tidak ditemukan!")
	}

	return result, nil
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
	student, err := service.GetStudent("Ihsan Arif")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(student)
	}
}
