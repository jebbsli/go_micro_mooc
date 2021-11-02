package service

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"user/domain/model"
	"user/domain/repository"
)

type IUserDataService interface {
	AddUser(user *model.User) (int64, error)
	DeleteUser(int642 int64) error
	UpdateUser(user *model.User, isChangePwd bool) (err error)
	FindUserByName(string) (*model.User, error)
	CheckPwd(username string, pwd string) (isOk bool, err error)
}

func NewUserDataService(userRepository repository.IUserRepository) IUserDataService {
	return &UserDataService{UserRepository: userRepository}
}

type UserDataService struct {
	UserRepository repository.IUserRepository
}

func GeneratePassword(userPassword string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(userPassword), bcrypt.DefaultCost)
}

func ValidatePassword(userPassword string, hashed string) (isOk bool, err error) {
	if err = bcrypt.CompareHashAndPassword([]byte(hashed), []byte(userPassword)); err != nil {
		return false, errors.New("password error")
	}

	return true, nil
}

func (u *UserDataService) AddUser(user *model.User) (int64, error) {
	pwdBytes, err := GeneratePassword(user.HashPassword)
	if err != nil {
		return user.ID, err
	}

	user.HashPassword = string(pwdBytes)
	return u.UserRepository.CreateUser(user)
}

func (u *UserDataService) DeleteUser(userId int64) error {
	return u.UserRepository.DeleteUserByID(userId)
}

func (u *UserDataService) UpdateUser(user *model.User, isChangePwd bool) (err error) {
	if isChangePwd {
		pwdBytes, err := GeneratePassword(user.HashPassword)
		if err != nil {
			return err
		}
		user.HashPassword = string(pwdBytes)
	}
	return u.UserRepository.UpdateUser(user)
}

func (u *UserDataService) FindUserByName(username string) (*model.User, error) {
	return u.UserRepository.FindUserByName(username)
}

func (u *UserDataService) CheckPwd(username string, pwd string) (isOk bool, err error) {
	user, err := u.UserRepository.FindUserByName(username)
	if err != nil {
		return false, err
	}
	return ValidatePassword(pwd, user.HashPassword)
}
