package main

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type StudentRepositoryMock struct {
	mock.Mock
}

func (r StudentRepositoryMock) GetAllStudents() ([]Student, error) {
	args := r.Called()
	studentList := []Student{
		{FullName: "Ihsan Arif", Grade: "B", Class: 1},
		{FullName: "Tono", Grade: "A", Class: 2},
		{FullName: "Andi", Grade: "C", Class: 3},
	}
	return studentList, args.Error(1)
}

func TestService_GetStudent(t *testing.T) {
	repository := StudentRepositoryMock{}
	repository.On("GetAllStudents").Return([]Student{}, nil)

	service := StudentService{repository}
	studentList, _ := service.GetStudent()
	expected := []Student{
		{FullName: "Ihsan Arif", Grade: "B", Class: 1},
		{FullName: "Tono", Grade: "A", Class: 2},
		{FullName: "Andi", Grade: "C", Class: 3},
	}
	for i := range studentList {
		assert.Equal(t, studentList[i].FullName, expected[i].FullName, "fullname must be"+expected[i].FullName)
		assert.Equal(t, studentList[i].Grade, expected[i].Grade, "fullname must be"+expected[i].Grade)
		assert.Equal(t, studentList[i].Class, expected[i].Class, "fullname must be"+strconv.Itoa(expected[i].Class))
	}
	fmt.Println(studentList)
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
			name: "case when get data student",
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
			name: "case when data student not found",
			fields: fields{
				StudentRepositoryInterface: &StudentRepositoryInterfaceMock{
					GetAllStudentsFunc: func() ([]Student, error) {
						return nil, errors.New("student not found")
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
			got, err := s.GetStudent()
			if (err != nil) != tt.wantErr {
				t.Errorf("StudentService.GetUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StudentService.GetUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserRepository_GetAllUsers(t *testing.T) {
	tests := []struct {
		name    string
		r       StudentRepository
		want    []Student
		wantErr bool
	}{
		{
			name: "get all data student",
			want: []Student{
				{FullName: "Ihsan Arif", Grade: "B", Class: 1},
				{FullName: "Tono", Grade: "A", Class: 2},
				{FullName: "Andi", Grade: "C", Class: 3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := StudentRepository{}
			got, err := r.GetAllStudents()
			if (err != nil) != tt.wantErr {
				t.Errorf("UserRepository.GetAllUsers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserRepository.GetAllUsers() = %v, want %v", got, tt.want)
			}
		})
	}
}
