package main

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func TestMain_GetAllStudents(t *testing.T) {
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
			name: "Case ambil seluruh data student",
			fields: fields{
				StudentRepositoryInterface: &StudentRepositoryInterfaceMock{
					GetAllStudentsFunc: func() ([]Student, error) {
						Students := []Student{
							{FullName: "Ihsan Arif", Grade: "B", Class: 1},
							{FullName: "Tono", Grade: "A", Class: 2},
							{FullName: "Andi", Grade: "C", Class: 3},
						}
						return Students, nil
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
			name: "Case ambil seluruh data student tapi data kosong",
			fields: fields{
				StudentRepositoryInterface: &StudentRepositoryInterfaceMock{
					GetAllStudentsFunc: func() ([]Student, error) {
						Students := []Student{}
						if len(Students) < 1 {
							return nil, errors.New("Data Kosong")
						} else {
							return Students, nil
						}
					},
				},
			},
			want: []Student{
				{FullName: "Ihsan Arif", Grade: "B", Class: 1},
				{FullName: "Tono", Grade: "A", Class: 2},
				{FullName: "Andi", Grade: "C", Class: 3},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := StudentService{
				StudentRepositoryInterface: tt.fields.StudentRepositoryInterface,
			}
			got, err := s.GetAllStudents()
			if err != nil {
				t.Errorf("StudentService.getAllStudents() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StudentService.getAllStudents() = %v, wantErr %v", got, tt.want)
				return
			}

			fmt.Printf("Test dengan nama %v berhasil.\n", tt.name)

		})
	}

}
