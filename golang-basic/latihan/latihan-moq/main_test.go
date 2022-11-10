package main

import (
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
				UserRepositoryInterface: UserRepositoryMock{},
			},
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
