package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"reflect"
	"testing"
)

type StudentRepositoryMock struct {
	mock.Mock
}

func (r StudentRepositoryMock) GetAllStudents() ([]Student, error) {
	args := r.Called()
	students := []Student{
		{"Tono", "A", 2},
	}
	return students, args.Error(1)
}

func TestService_GetStudent(t *testing.T) {
	repository := StudentRepositoryMock{}
	repository.On("GetAllStudents").Return([]Student{}, nil)

	service := StudentService{repository}
	students, _ := service.GetStudent()
	for i := range students {
		assert.Equal(t, students[i].FullName, "Tono", "nama harus sama")
		assert.Equal(t, students[i].Grade, "A", "grade harus A")
		assert.Equal(t, students[i].Class, "1", "grade harus 1")
	}
	fmt.Println(students)
}

func TestStudentService_GetStudent(t *testing.T) {
	type fields struct {
		StudentRepositoryInterface StudentRepositoryInterface
	}
	tests := []struct {
		name    string
		fields  fields
		want    []Student
		wantErr bool
	}{
		{
			name: "case ambil data user",
			fields: fields{
				StudentRepositoryInterface: StudentRepositoryMock{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := StudentService{
				StudentRepositoryInterface: tt.fields.StudentRepositoryInterface,
			}
			got, err := s.GetStudent()
			if (err != nil) != tt.wantErr {
				t.Errorf("StudentService.GetStudent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StudentService.GetStudent() = %v, want %v", got, tt.want)
			}
		})
	}
}
