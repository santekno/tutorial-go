package main

import (
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type StudentRepositoryMock struct {
	mock.Mock
}

func (s StudentRepositoryMock) GetAllStudents() ([]Student, error) {
	args := s.Called()
	students := []Student{
		{FullName: "Ihsan Arif", Grade: "B", Class: 1},
		{FullName: "Tono", Grade: "A", Class: 2},
		{FullName: "Andi", Grade: "C", Class: 3},
		{FullName: "John Doe", Grade: "A", Class: 4},
	}
	return students, args.Error(1)
}

func TestService_GetStudent(t *testing.T) {
	repository := StudentRepositoryMock{}
	repository.On("GetAllStudents").Return(Student{}, nil)

	service := StudentService{&repository}
	students, err := service.GetStudents()
	expected := []Student{
		{FullName: "Ihsan Arif", Grade: "B", Class: 1},
		{FullName: "Tono", Grade: "A", Class: 2},
		{FullName: "Andi", Grade: "C", Class: 3},
		{FullName: "John Doe", Grade: "A", Class: 4},
	}

	assert.Equal(t, expected, students, fmt.Sprintf("result must be %v", expected))
	assert.Nil(t, err, "error must be nil")
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
		// TODO: Add test cases.
		{
			name: "Success get student",
			fields: fields{
				StudentRepositoryInterface: &StudentRepositoryInterfaceMock{
					GetAllStudentsFunc: func() ([]Student, error) {
						return []Student{
							{FullName: "Ihsan Arif", Grade: "B", Class: 1},
						}, nil
					},
				},
			},
			want: []Student{
				{FullName: "Ihsan Arif", Grade: "B", Class: 1},
			},
			wantErr: false,
		},
		{
			name: "Failed get student",
			fields: fields{
				StudentRepositoryInterface: &StudentRepositoryInterfaceMock{
					GetAllStudentsFunc: func() ([]Student, error) {
						return nil, errors.New("[500] internal server error")
					},
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := StudentService{
				StudentRepositoryInterface: tt.fields.StudentRepositoryInterface,
			}
			got, err := s.GetStudents()
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

func TestStudentRepository_GetAllStudents(t *testing.T) {
	tests := []struct {
		name    string
		s       StudentRepository
		want    []Student
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Get all students",
			want: []Student{
				{FullName: "Ihsan Arif", Grade: "B", Class: 1},
				{FullName: "Tono", Grade: "A", Class: 2},
				{FullName: "Andi", Grade: "C", Class: 3},
				{FullName: "John Doe", Grade: "A", Class: 4},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := StudentRepository{}
			got, err := s.GetAllStudents()
			if (err != nil) != tt.wantErr {
				t.Errorf("StudentRepository.GetAllStudents() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StudentRepository.GetAllStudents() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{
			name: "Success",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
