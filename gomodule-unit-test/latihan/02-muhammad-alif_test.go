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
		// TODO: Add test cases.
		{
			name: "success get data all student",
			fields: fields{
				StudentRepositoryInterface: &StudentRepositoryInterfaceMock{
					GetAllStudentsFunc: func() ([]Student, error) {
						return []Student{
							{
								FullName: "Ihsan Arif",
								Grade:    "B",
								Class:    1,
							},
							{
								FullName: "Tono",
								Grade:    "A",
								Class:    2,
							},
							{
								FullName: "Andi",
								Grade:    "C",
								Class:    3,
							},
						}, nil
					},
				},
			},
			want: []Student{
				{
					FullName: "Ihsan Arif",
					Grade:    "B",
					Class:    1,
				},
				{
					FullName: "Tono",
					Grade:    "A",
					Class:    2,
				},
				{
					FullName: "Andi",
					Grade:    "C",
					Class:    3,
				},
			},
		},
		{
			name: "failed when get data student",
			fields: fields{
				StudentRepositoryInterface: &StudentRepositoryInterfaceMock{
					GetAllStudentsFunc: func() ([]Student, error) {
						return nil, errors.New("got error")
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
