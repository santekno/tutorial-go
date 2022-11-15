package main

import (
	"errors"
	"reflect"
	"testing"
)

func TestUserService_GetUser(t *testing.T) {
	type fields struct {
		UserRepositoryInterface UserRepositoryInterface
	}
	tests := []struct {
		name    string
		fields  fields
		want    []User
		wantErr bool
	}{
		{
			name: "case 1 : error get all users",
			fields: fields{
				UserRepositoryInterface: &UserRepositoryInterfaceMock{
					GetAllUsersFunc: func() ([]User, error) {
						return nil, errors.New("got error")
					},
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "case 2 : success get user",
			fields: fields{
				UserRepositoryInterface: &UserRepositoryInterfaceMock{
					GetAllUsersFunc: func() ([]User, error) {
						return []User{
							{
								Username: "faiz",
								Password: "Baraja",
							},
							{
								Username: "faiz",
								Password: "Baraja",
							},
						}, nil
					},
				},
			},
			want: []User{
				{
					Username: "faiz",
					Password: "*****",
				},
				{
					Username: "faiz",
					Password: "*****",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := UserService{
				UserRepositoryInterface: tt.fields.UserRepositoryInterface,
			}
			got, err := s.GetUser()
			if (err != nil) != tt.wantErr {
				t.Errorf("UserService.GetUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserService.GetUser() = %v, want %v", got, tt.want)
			}
		})
	}
}
