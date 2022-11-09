package main

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"reflect"
	"testing"
)

type StudentRepositoryMock struct {
	mock.Mock
}

func (s *StudentRepositoryMock) GetAllStudents() ([]Student, error) {
	args := s.Called()
	student := []Student{
		{FullName: "Ihsan Arif", Grade: "B", Class: 1},
		{FullName: "Tono", Grade: "A", Class: 2},
		{FullName: "Andi", Grade: "C", Class: 3},
		{FullName: "Rendy", Grade: "A", Class: 4},
	}
	return student, args.Error(1)
}

func TestService_GetStudent(t *testing.T) {
	repository := StudentRepositoryMock{}
	repository.On("GetAllStudents").Return(Student{}, nil)

	service := StudentService{&repository}
	student, err := service.GetStudent()
	expected := []Student{
		{FullName: "Ihsan Arif", Grade: "B", Class: 1},
		{FullName: "Tono", Grade: "A", Class: 2},
		{FullName: "Andi", Grade: "C", Class: 3},
		{FullName: "Rendy", Grade: "A", Class: 4},
	}
	assert.Equal(t, expected, student, fmt.Sprintf("result must be %v", expected))
	assert.Nil(t, err, "error must be nil")
}

func TestStudentRepository_GetAllStudents(t *testing.T) {
	tests := []struct {
		name    string
		want    []Student
		wantErr bool
	}{
		{
			name: "case get all data student",
			want: []Student{
				{FullName: "Ihsan Arif", Grade: "B", Class: 1},
				{FullName: "Tono", Grade: "A", Class: 2},
				{FullName: "Andi", Grade: "C", Class: 3},
				{FullName: "Rendy", Grade: "A", Class: 4},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &StudentRepository{}
			got, err := r.GetAllStudents()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllStudents() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllStudents() got = %v, want %v", got, tt.want)
			}
		})
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
		{
			name: "case get data student",
			fields: fields{StudentRepositoryInterface: &StudentRepositoryInterfaceMock{
				GetAllStudentsFunc: func() ([]Student, error) {
					return []Student{
						{FullName: "Ihsan Arif", Grade: "B", Class: 1},
						{FullName: "Tono", Grade: "A", Class: 2},
						{FullName: "Andi", Grade: "C", Class: 3},
						{FullName: "Rendy", Grade: "A", Class: 4},
					}, nil
				},
			}},
			want: []Student{
				{FullName: "Ihsan Arif", Grade: "B", Class: 1},
				{FullName: "Tono", Grade: "A", Class: 2},
				{FullName: "Andi", Grade: "C", Class: 3},
				{FullName: "Rendy", Grade: "A", Class: 4},
			},
			wantErr: false,
		},
		{
			name: "case gagal get data student",
			fields: fields{StudentRepositoryInterface: &StudentRepositoryInterfaceMock{
				GetAllStudentsFunc: func() ([]Student, error) {
					return nil, errors.New("students not found")
				},
			}},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &StudentService{
				StudentRepositoryInterface: tt.fields.StudentRepositoryInterface,
			}
			got, err := s.GetStudent()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetStudent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetStudent() got = %v, want %v", got, tt.want)
			}
			assert.Equal(t, tt.want, got, fmt.Sprintf("result must be %v", tt.want))
		})
	}
}

func Test_Main(t *testing.T) {
	main()
}
