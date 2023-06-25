package userService

import (
	userStructs "github.com/masharpik/ForumVKEducation/app/users/utils/structs"
)

func (service service) CreateUser(user userStructs.User) (createdUser userStructs.User, err error) {
	createdUser, err = service.repo.CreateUser(user)
	if err != nil {
		return
	}

	return
}

func (service service) GetUsersByNickOrMail(user userStructs.User) (createdUsers userStructs.Users, err error) {
	createdUsers, err = service.repo.GetUsersByNickOrMail(user.Nickname, user.Email)
	if err != nil {
		return
	}

	return
}

func (service service) GetUserInfo(nickname string) (user userStructs.User, err error) {
	user, err = service.repo.GetUserInfo(nickname)
	if err != nil {
		return
	}

	return
}

func (service service) UpdateUserInfo(user userStructs.User) (userStructs.User, error) {
	gotUser, err := service.repo.GetUserInfo(user.Nickname)
	if err != nil {
		return userStructs.User{}, err
	}

	if user.Email == "" {
		user.Email = gotUser.Email
	}
	if user.Fullname == "" {
		user.Fullname = gotUser.Fullname
	}
	if user.About == "" {
		user.About = gotUser.About
	}

	err = service.repo.UpdateUserInfo(user)
	if err != nil {
		return userStructs.User{}, err
	}

	return user, nil
}

func (service service) GetNickByNick(nickname string) (originNick string, err error) {
	originNick, err = service.repo.GetNickByNick(nickname)
	if err != nil {
		return
	}

	return
}
