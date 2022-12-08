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
		{
			name: "succes expected get all students",
			s: StudentService{
				StudentRepositoryInterface: &StudentRepositoryInterfaceMock{
					GetAllStudentsFunc: func() ([]Student, error) {
						return []Student{
							{
								FullName: "Zenandi",
								Grade:    "B",
								Class:    2,
							},
							{
								FullName: "Barkah",
								Grade:    "C",
								Class:    1,
							},
							{
								FullName: "Tariadi",
								Grade:    "A",
								Class:    1,
							},
						}, nil
					},
				},
			},
			want: []Student{
				{
					FullName: "Zenandi",
					Grade:    "B",
					Class:    2,
				},
				{
					FullName: "Barkah",
					Grade:    "C",
					Class:    1,
				},
				{
					FullName: "Tariadi",
					Grade:    "A",
					Class:    1,
				},
			},
			wantErr: false,
		},
		{
			name: "expected error when getting all student",
			s: StudentService{
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
