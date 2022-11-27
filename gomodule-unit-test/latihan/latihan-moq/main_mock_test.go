package main

import (
	"sync"
)

var _ StudentRepositoryInterface = &StudentRepositoryInterfaceMock{}

type StudentRepositoryInterfaceMock struct {
	GetAllStudentsFunc func() ([]Student, error)

	calls struct {
		GetAllStudents []struct {
		}
	}
	lockGetAllStudents sync.RWMutex
}

func (mock *StudentRepositoryInterfaceMock) GetAllStudents() ([]Student, error) {
	if mock.GetAllStudentsFunc == nil {
		panic("StudentRepositoryInterfaceMock.GetAllStudentsFunc: method is nil but StudentRepositoryInterface.GetAllStudents was just called")
	}
	callInfo := struct {
	}{}
	mock.lockGetAllStudents.Lock()
	mock.calls.GetAllStudents = append(mock.calls.GetAllStudents, callInfo)
	mock.lockGetAllStudents.Unlock()
	return mock.GetAllStudentsFunc()
}

func (mock *StudentRepositoryInterfaceMock) GetAllStudentsCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockGetAllStudents.RLock()
	calls = mock.calls.GetAllStudents
	mock.lockGetAllStudents.RUnlock()
	return calls
}
