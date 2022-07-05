package main

import (
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	mock.Mock
}

func (r UserRepositoryMock) GetAllUsers() ([]User, error) {
	args := r.Called()
	users := []User{
		{"mock", "*****"},
	}
	return users, args.Error(1)
}

func TestService_GetUser(t *testing.T) {
	repository := UserRepositoryMock{}
	repository.On("GetAllUsers").Return([]User{}, nil)

	service := UserService{repository}
	users, _ := service.GetUser()
	for i := range users {
		assert.Equal(t, users[i].Password, "*****", "user password must be encrypted")
	}
	fmt.Println(users)
}

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
			name: "case ambil data user",
			fields: fields{
				UserRepositoryInterface: &UserRepositoryInterfaceMock{
					GetAllUsersFunc: func() ([]User, error) {
						return []User{
							{
								Username: "ihsan",
								Password: "*****",
							},
						}, nil
					},
				},
			},
			want: []User{
				{
					Username: "ihsan",
					Password: "*****",
				},
			},
			wantErr: false,
		},
		{
			name: "case saat ambil data error",
			fields: fields{
				UserRepositoryInterface: &UserRepositoryInterfaceMock{
					GetAllUsersFunc: func() ([]User, error) {
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

func TestUserRepository_GetAllUsers(t *testing.T) {
	tests := []struct {
		name    string
		r       UserRepository
		want    []User
		wantErr bool
	}{
		{
			name: "get all data user",
			want: []User{
				{"real", "real"},
				{"real2", "real2"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := UserRepository{}
			got, err := r.GetAllUsers()
			if (err != nil) != tt.wantErr {
				t.Errorf("UserRepository.GetAllUsers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserRepository.GetAllUsers() = %v, want %v", got, tt.want)
			}
		})
	}
}
