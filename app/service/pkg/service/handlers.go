package serviceService

import serviceStructs "github.com/masharpik/ForumVKEducation/app/service/utils/structs"

func (service service) ClearDB() (err error) {
	err = service.repo.ClearDB()

	return
}

func (service service) GetInfoDB() (status serviceStructs.Status, err error) {
	status, err = service.repo.GetInfoDB()

	return
}
