package repository

import (
	"github.com/jinzhu/gorm"
	"user/domain/model"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type IUserRepository interface {
	InitTable() error

	FindUserByName(string) (*model.User, error)
	FindUserByID(int64) (*model.User, error)
	CreateUser(user *model.User) (int64, error)
	DeleteUserByID(int642 int64) error
	UpdateUser(user *model.User) error
	FindAll() ([]model.User, error)
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{mysqlDb: db}
}

type UserRepository struct {
	mysqlDb *gorm.DB
}

func (u *UserRepository) InitTable() error {
	return u.mysqlDb.CreateTable(&model.User{}).Error
}

func (u *UserRepository) FindUserByName(name string) (*model.User, error) {
	user := &model.User{}
	return user, u.mysqlDb.Where("user_name=?", name).Find(user).Error
}

func (u *UserRepository) FindUserByID(userId int64) (*model.User, error) {
	user := &model.User{}
	return user, u.mysqlDb.First(user, userId).Error
}

func (u *UserRepository) CreateUser(user *model.User) (int64, error) {
	return user.ID, u.mysqlDb.Create(user).Error
}

func (u *UserRepository) DeleteUserByID(userId int64) error {
	return u.mysqlDb.Where("ID=?", userId).Delete(&model.User{}).Error
}

func (u *UserRepository) UpdateUser(user *model.User) error {
	return u.mysqlDb.Model(user).Update(&user).Error
}

func (u *UserRepository) FindAll() ([]model.User, error) {
	var userAll []model.User
	return userAll, u.mysqlDb.Find(&userAll).Error
}
