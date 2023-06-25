package interfaces

import userStructs "github.com/masharpik/ForumVKEducation/app/users/utils/structs"

type IService interface {
	CreateUser(userStructs.User) (userStructs.User, error)
	GetUsersByNickOrMail(userStructs.User) (userStructs.Users, error)
	GetUserInfo(string) (userStructs.User, error)
	UpdateUserInfo(userStructs.User) (userStructs.User, error)
	GetNickByNick(string) (string, error)
}
