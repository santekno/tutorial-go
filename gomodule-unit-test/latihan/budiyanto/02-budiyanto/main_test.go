package main

import (
	"errors"
	"reflect"
	"testing"
)

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
				StudentRepositoryInterface: &StudentRepositoryInterfaceMock{
					GetAllStudentsFunc: func() ([]Student, error) {
						return []Student{
							{FullName: "Budiyanto", Grade: "X", Class: 6},
							{FullName: "Yor", Grade: "Y", Class: 5},
							{FullName: "Anya", Grade: "Z", Class: 4},
						}, nil
					},
				},
			},
			want: []Student{
				{FullName: "Ihsan Arif", Grade: "X", Class: 6},
				{FullName: "Ihsan Arif", Grade: "Y", Class: 5},
				{FullName: "Ihsan Arif", Grade: "Z", Class: 4},
			},
			wantErr: false,
		},
		{
			name: "case saat ambil data error",
			fields: fields{
				StudentRepositoryInterface: &StudentRepositoryInterfaceMock{
					GetAllStudentsFunc: func() ([]Student, error) {
						return nil, errors.New("error")
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
		r       StudentRepositoryInterface
		want    []Student
		wantErr bool
	}{
		{
			name: "case ambil data user",
			r: &StudentRepositoryInterfaceMock{
				GetAllStudentsFunc: func() ([]Student, error) {
					return []Student{
						{FullName: "Budiyanto", Grade: "X", Class: 6},
						{FullName: "Yor", Grade: "Y", Class: 5},
						{FullName: "Anya", Grade: "Z", Class: 4},
					}, nil
				},
			},
			want: []Student{
				{FullName: "Budiyanto", Grade: "X", Class: 6},
				{FullName: "Yor", Grade: "Y", Class: 5},
				{FullName: "Anya", Grade: "Z", Class: 4},
			},
			wantErr: false,
		},
		{
			name: "case saat ambil data error",
			r: &StudentRepositoryInterfaceMock{
				GetAllStudentsFunc: func() ([]Student, error) {
					return nil, errors.New("error")
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := StudentService{tt.r}
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
