package main

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func TestStudentService_GetStudent(t *testing.T) {
	type fields struct {
		StudentRepository StudentRepositoryInterface
	}
	tests := []struct {
		testName  string
		fields    fields
		expect    []Student
		expectErr bool
	}{
		{
			testName: "it should return all students",
			fields: fields{
				StudentRepository: &StudentRepositoryInterfaceMock{
					GetAllStudentsFunc: func() ([]Student, error) {
						return []Student{
							{FullName: "Yogi Pristiawan", Grade: "B", Class: 12},
						}, nil
					},
				},
			},
			expect: []Student{
				{FullName: "Yogi Pristiawan", Grade: "B", Class: 12},
			},
			expectErr: false,
		},
		{
			testName: "it should return error if any",
			fields: fields{
				StudentRepository: &StudentRepositoryInterfaceMock{
					GetAllStudentsFunc: func() ([]Student, error) {
						return nil, errors.New("error while fetching data")
					},
				},
			},
			expect:    nil,
			expectErr: true,
		},
	}

	for _, tt := range tests {
		service := StudentService{
			StudentRepositoryInterface: tt.fields.StudentRepository,
		}

		got, err := service.GetStudent()
		if (err != nil) != tt.expectErr {
			t.Errorf("StudentService.GetStudent expected error %v to be PASSED but got %v instead", tt.expectErr, err)
		}

		if !reflect.DeepEqual(tt.expect, got) {
			t.Errorf("StudentService.GetStudent expected %v to be PASSED but got %v instead", tt.expect, got)
		}
	}
}

func TestStudentRepository_GetAllStudents(t *testing.T) {
	tests := []struct {
		testName  string
		expect    []Student
		expectErr bool
	}{
		{
			testName: "it should return all students",
			expect: []Student{
				{FullName: "Ihsan Arif", Grade: "B", Class: 1},
				{FullName: "Tono", Grade: "A", Class: 2},
				{FullName: "Andi", Grade: "C", Class: 3},
			},
			expectErr: false,
		},
	}

	for _, tt := range tests {
		repository := StudentRepository{}
		got, err := repository.GetAllStudents()

		if (err != nil) != tt.expectErr {
			t.Errorf("StudentRepository.GetAllStudents expect error %v to be PASSED but got %v instead", tt.expectErr, err)
		}

		if !reflect.DeepEqual(tt.expect, got) {
			t.Errorf("StudentRepository.GetAllStudents expect %v to be PASSED but got %v instead", tt.expect, got)
		}
	}
}
