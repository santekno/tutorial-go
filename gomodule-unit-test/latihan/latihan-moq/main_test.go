package main

import (
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
			name: "sukses",
			s: StudentService{
				StudentRepositoryInterface: &StudentRepositoryInterfaceMock{
					GetAllStudentsFunc: func() ([]Student, error) {
						return []Student{
							{
								FullName: "Ihsan Arief",
								Grade:    "B",
								Class:    1,
							},
							{
								FullName: "Tono",
								Grade:    "A",
								Class:    2,
							},
						}, nil
					},
				},
			},
			wantErr: false,
			want: []Student{
				{
					FullName: "Ihsan Arief",
					Grade:    "B",
					Class:    1,
				},
				{
					FullName: "Tono",
					Grade:    "A",
					Class:    2,
				},
			},
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
