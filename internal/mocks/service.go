package mocks

import "github.com/stretchr/testify/mock"

type ServiceMock struct {
	mock.Mock
}

func NewServiceMock() *ServiceMock {
	return &ServiceMock{}
}
