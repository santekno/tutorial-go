package main

import (
	"errors"
	"reflect"
	"testing"
)

func TestStudentRepository_GetAllStudents(t *testing.T) {
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
