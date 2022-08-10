package main

import (
	"errors"
	"reflect"
	"testing"
)

func TestStudentRepository_GetStudents(t *testing.T) {
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
							{
								FullName: "Tono",
								Grade:    "B",
								Class:    1,
							},
							{
								FullName: "Arif",
								Grade:    "B",
								Class:    1,
							},
						}, nil
					},
				},
			},
			want: []Student{
				{
					FullName: "Tono",
					Grade:    "B",
					Class:    1,
				},
				{
					FullName: "Arif",
					Grade:    "B",
					Class:    1,
				},
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
		r       StudentRepository
		want    []Student
		wantErr bool
	}{
		{
			name: "get all data student",
			want: []Student{
				{"Ihsan Arif", "B", 1},
				{"Tono", "A", 2},
				{"Andi", "C", 3},
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
