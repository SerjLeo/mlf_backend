package mocks

import (
	"github.com/SerjLeo/mlf_backend/internal/models"
)

func (o *ServiceMock) Create(user models.User) (string, error) {
	args := o.Called(user)
	return args.String(0), args.Error(1)
}

func (o *ServiceMock) CreateUserByEmail(email string) (string, error) {
	args := o.Called(email)
	return args.String(0), args.Error(1)
}

func (o *ServiceMock) SignIn(email, password string) (string, error) {
	args := o.Called(email, password)
	return args.String(0), args.Error(1)
}

func (o *ServiceMock) CheckUserToken(token string) (int, error) {
	args := o.Called(token)
	return args.Int(0), args.Error(1)
}

func (o *ServiceMock) GetUserProfile(userId int) (models.User, error) {
	args := o.Called(userId)
	return models.User{}, args.Error(1)
}

func (o *ServiceMock) SendTestEmail() error {
	args := o.Called()
	return args.Error(1)
}
