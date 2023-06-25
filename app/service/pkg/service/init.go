package serviceService

import (
	"github.com/masharpik/ForumVKEducation/app/service/utils/interfaces"
)

type service struct {
	repo interfaces.IRepository
}

func NewService(repo interfaces.IRepository) interfaces.IService {
	service := service{
		repo: repo,
	}

	return service
}
