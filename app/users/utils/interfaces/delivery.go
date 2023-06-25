package interfaces

import "net/http"

type IRouter interface {
	// swagger:route POST /user/{nickname}/create userCreate
	// Создание нового пользователя
	//
	// Создание нового пользователя в базе данных.
	//
	// responses:
	//   201: User Пользователь успешно создан.
	//        Возвращает данные созданного пользователя.
	//   409: Users Пользователь уже присутсвует в базе данных.
	//        Возвращает данные ранее созданных пользователей с тем же nickname-ом иои email-ом.
	// parameters:
	// + name: nickname
	//   in: path
	//   description: Идентификатор пользователя.
	//   required: true
	//   type: string
	// + name: profile
	//   in: body
	//   description: Данные пользовательского профиля.
	//   required: true
	//   type: User
	CreateUser(http.ResponseWriter, *http.Request)

	// swagger:route GET /user/{nickname}/profile userGetOne
	// Получение информации о пользователе
	//
	// Получение информации о пользователе форума по его имени.
	//
	// responses:
	//   200: User Информация о пользователе.
	//   404: Error Пользователь отсутсвует в системе.
	// parameters:
	// + name: nickname
	//	 in: path
	//   description: Идентификатор пользователя.
	//   required: true
	//   type: string
	GetUserInfo(http.ResponseWriter, *http.Request)

	// swagger:route POST /user/{nickname}/profile userUpdate
	// Изменение данных о пользователе
	//
	// Изменение информации в профиле пользователя.
	//
	// responses:
	//   200: User Актуальная информация о пользователе после изменения профиля.
	//   404: Error Пользователь отсутсвует в системе.
	//   409: Error Новые данные профиля пользователя конфликтуют с имеющимися пользователями.
	// parameters:
	// + name: nickname
	//	 in: path
	//   description: Идентификатор пользователя.
	//   required: true
	//   type: string
	// + name: profile
	//	 in: body
	//   description: Изменения профиля пользователя.
	//   required: true
	//   type: UserUpdate
	UpdateUserInfo(http.ResponseWriter, *http.Request)
}
