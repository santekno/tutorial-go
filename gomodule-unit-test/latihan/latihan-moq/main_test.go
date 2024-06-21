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
			name: "get student success",
			s: StudentService{
				StudentRepositoryInterface: &StudentRepositoryInterfaceMock{
					GetAllStudentsFunc: func() ([]Student, error) {
						return []Student{
							{FullName: "Ihsan Arif", Grade: "B", Class: 1},
							{FullName: "Tono", Grade: "A", Class: 2},
							{FullName: "Andi", Grade: "C", Class: 3},
						}, nil
					},
				},
			},
			wantErr: false,
			want: []Student{
				{FullName: "Ihsan Arif", Grade: "B", Class: 1},
				{FullName: "Tono", Grade: "A", Class: 2},
				{FullName: "Andi", Grade: "C", Class: 3},
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
