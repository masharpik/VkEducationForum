package interfaces

import (
	forumStructs "github.com/masharpik/ForumVKEducation/app/forums/utils/structs"
	userStructs "github.com/masharpik/ForumVKEducation/app/users/utils/structs"
)

type IService interface {
	CreateForum(forumStructs.Forum) (forumStructs.Forum, error)
	GetForumInfo(string) (forumStructs.Forum, error)
	GetForumUsers(string, string, int, bool) (userStructs.Users, error)
	CheckExistenceForum(string) (bool, error)
	GetForumSlugBySlug(string) (string, error)
}
