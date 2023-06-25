package interfaces

import serviceStructs "github.com/masharpik/ForumVKEducation/app/service/utils/structs"

type IService interface {
	ClearDB() error
	GetInfoDB() (serviceStructs.Status, error)
}
