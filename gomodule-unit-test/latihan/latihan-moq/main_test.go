package main

import (
	"errors"
	"reflect"
	"testing"
)

func TestStudentService_GetStudent(t *testing.T) {
	tests := []struct {
		name    string
		s       StudentService
		want    []Student
		wantErr bool
	}{
		{
			name: "success get data student all",
			s: StudentService{
				StudentRepositoryInterface: &StudentRepositoryInterfaceMock{
					GetAllStudentsFunc: func() ([]Student, error) {
						return []Student{
							{
								FullName: "Syamsul",
								Grade:    "A",
								Class:    1,
							},
							{
								FullName: "Bahari",
								Grade:    "B",
								Class:    1,
							},
						}, nil
					},
				},
			},
			want: []Student{
				{
					FullName: "Syamsul",
					Grade:    "A",
					Class:    1,
				},
				{
					FullName: "Bahari",
					Grade:    "B",
					Class:    1,
				},
			},
		},
		{
			name: "Failed get data",
			s: StudentService{
				StudentRepositoryInterface: &StudentRepositoryInterfaceMock{
					GetAllStudentsFunc: func() ([]Student, error) {
						return nil, errors.New("get error")
					},
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.GetStudent()
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
