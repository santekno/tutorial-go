package student

import (
	"fmt"

	"github.com/hendralw/tutorial-go-unit-test/student"
)

func main() {
	repository := student.StudentRepository{}
	service := student.StudentService{repository}
	Students, _ := service.GetStudent()
	fmt.Println(Students)
}
