package interfaces

import (
	forumStructs "github.com/masharpik/ForumVKEducation/app/forums/utils/structs"
	userStructs "github.com/masharpik/ForumVKEducation/app/users/utils/structs"
)

type IRepository interface {
	CreateForum(forumStructs.Forum) (forumStructs.Forum, error)
	GetForumInfoBySlug(string) (forumStructs.Forum, error)
	GetUsers(string, string, int, bool) (userStructs.Users, error)
	CheckExistenceForum(string) (bool, error)
	GetForumSlugBySlug(string) (string, error)
}
