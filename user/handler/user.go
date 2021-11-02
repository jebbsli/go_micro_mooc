package handler

import (
	"context"
	"user/domain/model"
	"user/domain/service"

	user "user/proto/user"
)

type User struct{
	UserDataService service.IUserDataService
}

//注册
func (u *User) Register(ctx context.Context, userRegisterRequest *user.UserRegisterRequest,
	userRegisterResponse *user.UserRegisterResponse) error {
	userRegister := &model.User{
		Username: userRegisterRequest.UserName,
		Firstname: userRegisterRequest.FirstName,
		HashPassword: userRegisterRequest.Pwd,
	}
	_, err := u.UserDataService.AddUser(userRegister)
	if err != nil {
		return err
	}

	userRegisterResponse.Message = "add user success"
	return nil
}

//登录
func (u *User) Login(ctx context.Context, userLoginRequest *user.UserLoginRequest,
	userLoginResponse *user.UserLoginResponse) error {
	isOk, err := u.UserDataService.CheckPwd(userLoginRequest.UserName, userLoginRequest.Pwd)
	if err != nil {
		return err
	}

	userLoginResponse.IsSuccess = isOk
	return nil
}

//查询用户信息
func (u *User) GetUserInfo(ctx context.Context, userInfoRequest *user.UserInfoRequest,
	userInfoResponse *user.UserInfoResponse) error {
	userInfo, err := u.UserDataService.FindUserByName(userInfoRequest.UserName)
	if err != nil {
		return err
	}

	userInfoResponse = UserForResponse(userInfo)

	return nil
}

func UserForResponse(userModel *model.User) *user.UserInfoResponse {
	response := &user.UserInfoResponse{}
	response.UserName = userModel.Username
	response.FirstName = userModel.Firstname
	response.UserId = userModel.ID

	return response
}
