package main

import (
	"errors"
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
	students := []Student{
		{FullName: "Yogi Is Ariyanto", Grade: "B", Class: 1},
		{FullName: "Budi", Grade: "A", Class: 2},
		{FullName: "Agus", Grade: "C", Class: 3},
	}

	return students, args.Error(1)
}

func TestService_GetStudent(t *testing.T) {
	repository := StudentRepositoryMock{}
	repository.On("GetAllStudents").Return([]Student{}, nil)

	service := StudentService{repository}
	students, _ := service.GetStudent()
	studentData := []Student{
		{FullName: "Yogi Is Ariyanto", Grade: "B", Class: 1},
		{FullName: "Budi", Grade: "A", Class: 2},
		{FullName: "Agus", Grade: "C", Class: 3},
	}
	for i := range students {
		assert.Equal(t, students[i].FullName, studentData[i].FullName, "Fullname should be"+studentData[i].FullName)
		assert.Equal(t, students[i].Grade, studentData[i].Grade, "Grade student should be"+studentData[i].Grade)
		assert.Equal(t, students[i].Class, studentData[i].Class, "Student Class should be"+strconv.Itoa(studentData[i].Class))
	}
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
			name: "Success to get Student Data",
			fields: fields{
				StudentRepositoryInterface: &StudentRepositoryInterfaceMock{
					GetAllStudentsFunc: func() ([]Student, error) {
						return []Student{
							{FullName: "Yogi Is Ariyanto", Grade: "B", Class: 1},
						}, nil
					},
				},
			},
			want: []Student{
				{FullName: "Yogi Is Ariyanto", Grade: "B", Class: 1},
			},
			wantErr: false,
		},
		{
			name: "Failed to get Student Data",
			fields: fields{
				StudentRepositoryInterface: &StudentRepositoryInterfaceMock{
					GetAllStudentsFunc: func() ([]Student, error) {
						return nil, errors.New("Cant get the data, Student not found")
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
		r       StudentRepository
		want    []Student
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Get all Student Data",
			want: []Student{
				{FullName: "Yogi Is Ariyanto", Grade: "B", Class: 1},
				{FullName: "Budi", Grade: "A", Class: 2},
				{FullName: "Agus", Grade: "C", Class: 3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := StudentRepository{}
			got, err := r.GetAllStudents()
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
