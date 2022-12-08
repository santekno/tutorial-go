package main

import (
	"errors"
	"reflect"
	"testing"
)

func Test_StudentService_GetStudent(t *testing.T) {
	type fields struct {
		repo StudentRepositoryInterface
	}

	tests := []struct {
		name    string
		fields  fields
		want    []Student
		wantErr bool
	}{
		{
			name: "Success",
			fields: fields{
				repo: &StudentRepositoryInterfaceMock{
					GetAllStudentsFunc: func() ([]Student, error) {
						return []Student{
							{FullName: "Ihsan Arif", Grade: "B", Class: 1},
							{FullName: "Tono", Grade: "A", Class: 2},
							{FullName: "Andi", Grade: "C", Class: 3},
						}, nil
					},
				},
			},
			want: []Student{
				{FullName: "Ihsan Arif", Grade: "B", Class: 1},
				{FullName: "Tono", Grade: "A", Class: 2},
				{FullName: "Andi", Grade: "C", Class: 3},
			},
			wantErr: false,
		},
		{
			name: "Getting Error",
			fields: fields{
				repo: &StudentRepositoryInterfaceMock{
					GetAllStudentsFunc: func() ([]Student, error) {
						return nil, errors.New("terjadi kesalahan")
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
				tt.fields.repo,
			}
			got, err := s.GetAllStudents()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllStudents() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllStudents() got = %v, want %v", got, tt.want)
			}
		})
	}

}
