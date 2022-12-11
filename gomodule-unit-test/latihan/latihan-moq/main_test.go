package main

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type StudentRepositoryMock struct {
	mock.Mock
}

func (r StudentRepositoryMock) GetAllStudents() ([]Student, error) {
	args := r.Called()
	if err != nil {
		return nil, err
	}

	return Students, args.Error(1)
}

func TestStudentService_GetStudent(t *testing.T) {
	repository := StudentRepositoryMock{}
	repository.On("GetAllStudents").Return([]Student{}, nil)


	service := StudentService{repository}
	Students, _ := service.GetStudent()
	if err != nil {
		return nil, err
	}
	fmt.Println(users)
}


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
			name: "case ambil data student",
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