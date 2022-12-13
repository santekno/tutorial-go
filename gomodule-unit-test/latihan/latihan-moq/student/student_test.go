package student

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
		// TODO: Add test cases.
		{
			name: "Success GetAllStudent",
			s: StudentService{
				StudentRepositoryInterface: &StudentRepositoryInterfaceMock{
					GetAllStudentsFunc: func() ([]Student, error) {
						return []Student{
							{
								FullName: "Hendrawan",
								Grade:    "A",
								Class:    1,
							},
							{
								FullName: "Beni",
								Grade:    "B",
								Class:    2,
							},
							{
								FullName: "Yudi",
								Grade:    "A",
								Class:    2,
							},
						}, nil
					},
				},
			},
			want: []Student{
				{
					FullName: "Hendrawan",
					Grade:    "A",
					Class:    1,
				},
				{
					FullName: "Beni",
					Grade:    "B",
					Class:    2,
				},
				{
					FullName: "Yudi",
					Grade:    "A",
					Class:    2,
				},
			},
			wantErr: false,
		},
		{
			name: "Error GetAllStudent",
			s: StudentService{
				StudentRepositoryInterface: &StudentRepositoryInterfaceMock{
					GetAllStudentsFunc: func() ([]Student, error) {
						return nil, errors.New("Error GetAllStudent")
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
