package main

import (
	"fmt"

	"github.com/zenandibarkah/latihanzenandibarkahgomoduleunittest/student"
)

func main() {
	repository := student.StudentRepository{}
	service := student.StudentService{repository}
	Students, _ := service.GetStudent()
	fmt.Println(Students)
}
