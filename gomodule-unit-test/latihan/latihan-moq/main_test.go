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
			name: "success get students data",
			fields: fields{
				StudentRepositoryInterface: &StudentRepositoryInterfaceMock{
					GetAllStudentsFunc: func() ([]Student, error) {
						students := []Student{
							{FullName: "lele", Grade: "X", Class: 0},
						}
						return students, nil
					},
				},
			},
			want: []Student{
				{FullName: "lele", Grade: "X", Class: 0},
			},
			wantErr: false,
		},
		{
			name: "error get students data",
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
		r       StudentRepository
		want    []Student
		wantErr bool
	}{
		{
			name: "success get all students",
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
				t.Errorf("StudentRepository.GetAllStudents() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StudentRepository.GetAllStudents() = %v, want %v", got, tt.want)
			}
		})
	}
}
