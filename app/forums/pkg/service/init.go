package forumService

import (
	"github.com/masharpik/ForumVKEducation/app/forums/utils/interfaces"
	userInterfaces "github.com/masharpik/ForumVKEducation/app/users/utils/interfaces"
)

type service struct {
	repo     interfaces.IRepository
	userRepo userInterfaces.IRepository
}

func NewService(repo interfaces.IRepository, userRepo userInterfaces.IRepository) interfaces.IService {
	service := service{
		repo:     repo,
		userRepo: userRepo,
	}

	return service
}
