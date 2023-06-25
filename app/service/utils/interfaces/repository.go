package interfaces

import serviceStructs "github.com/masharpik/ForumVKEducation/app/service/utils/structs"

type IRepository interface {
	ClearDB() error
	GetInfoDB() (serviceStructs.Status, error)
}
