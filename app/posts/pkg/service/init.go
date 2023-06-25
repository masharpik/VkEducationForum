package postService

import (
	forumInterfaces "github.com/masharpik/ForumVKEducation/app/forums/utils/interfaces"
	"github.com/masharpik/ForumVKEducation/app/posts/utils/interfaces"
	threadInterfaces "github.com/masharpik/ForumVKEducation/app/threads/utils/interfaces"
	userInterfaces "github.com/masharpik/ForumVKEducation/app/users/utils/interfaces"
)

type service struct {
	repo       interfaces.IRepository
	userRepo   userInterfaces.IRepository
	forumRepo  forumInterfaces.IRepository
	threadRepo threadInterfaces.IRepository
}

func NewService(repo interfaces.IRepository, userRepo userInterfaces.IRepository, forumRepo forumInterfaces.IRepository, threadRepo threadInterfaces.IRepository) interfaces.IService {
	service := service{
		repo:       repo,
		userRepo:   userRepo,
		forumRepo:  forumRepo,
		threadRepo: threadRepo,
	}

	return service
}
