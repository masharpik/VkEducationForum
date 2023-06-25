package interfaces

import userStructs "github.com/masharpik/ForumVKEducation/app/users/utils/structs"

type IRepository interface {
	CreateUser(userStructs.User) (userStructs.User, error)
	GetUsersByNickOrMail(string, string) (userStructs.Users, error)
	GetUserInfo(string) (userStructs.User, error)
	UpdateUserInfo(userStructs.User) error
	GetNickByNick(string) (string, error)
}
