package forumService

import (
	"errors"

	forumLiterals "github.com/masharpik/ForumVKEducation/app/forums/utils/literals"
	forumStructs "github.com/masharpik/ForumVKEducation/app/forums/utils/structs"
	userLiterals "github.com/masharpik/ForumVKEducation/app/users/utils/literals"
	userStructs "github.com/masharpik/ForumVKEducation/app/users/utils/structs"
)

func (service service) CreateForum(forum forumStructs.Forum) (createdForum forumStructs.Forum, err error) {
	nickname, err := service.userRepo.GetNickByNick(forum.User)
	if err != nil {
		if err.Error() == userLiterals.UserNotFound {
			err = errors.New(forumLiterals.UserFKNotFound)
		}

		return
	}

	forum.User = nickname

	createdForum, err = service.repo.CreateForum(forum)
	if err != nil {
		if err.Error() == forumLiterals.ForumAlreadyExists {
			var errInfo error

			createdForum, errInfo = service.repo.GetForumInfoBySlug(forum.Slug)
			if errInfo != nil {
				return forumStructs.Forum{}, errInfo
			}
		}

		return
	}

	return
}

func (service service) GetForumInfo(slug string) (forum forumStructs.Forum, err error) {
	forum, err = service.repo.GetForumInfoBySlug(slug)
	if err != nil {
		return
	}

	return
}

func (service service) GetForumUsers(slug, sinceNickname string, limit int, desc bool) (users userStructs.Users, err error) {
	exists, err := service.repo.CheckExistenceForum(slug)
	if err != nil {
		return
	}
	if !exists {
		err = errors.New(forumLiterals.ForumNotFound)
		return
	}

	users, err = service.repo.GetUsers(slug, sinceNickname, limit, desc)
	if err != nil {
		return
	}

	return
}

func (service service) CheckExistenceForum(slug string) (exists bool, err error) {
	exists, err = service.repo.CheckExistenceForum(slug)
	if err != nil {
		return
	}

	return
}

func (service service) GetForumSlugBySlug(slug string) (originSlug string, err error) {
	originSlug, err = service.repo.GetForumSlugBySlug(slug)
	if err != nil {
		return
	}

	return
}
