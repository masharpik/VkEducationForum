package userStructs
//go:generate easyjson -all responses.go

// Информация о пользователе
// swagger:model
type User struct {
	// Имя пользователя (уникальное поле).
	// Данное поле допускает только латиницу, цифры и знак подчеркивания.
	// Сравнение имени регистронезависимо.
	// example: j.sparrow
	// readOnly: true
	// swagger:strfmt identity
	Nickname string `json:"nickname"`

	// Полное имя пользователя.
	// example: Captain Jack Sparrow
	// required: true
	Fullname string `json:"fullname"`

	// Описание пользователя.
	// example: This is the day you will always remember as the day that you almost caught Captain Jack Sparrow!
	// swagger:strfmt text
	About string `json:"about"`

	// Почтовый адрес пользователя (уникальное поле).
	// example: captaina@blackpearl.sea
	// swagger:strfmt email
	// required: true
	Email string `json:"email"`
}

//easyjson:json
// Список пользователей
// swagger:model
type Users []User

// Информация о пользователе для обновления
// swagger:model
type UserUpdate struct {
	// Полное имя пользователя.
	// example: Captain Jack Sparrow
	// required: true
	Fullname string `json:"fullname"`

	// Описание пользователя.
	// example: This is the day you will always remember as the day that you almost caught Captain Jack Sparrow!
	// swagger:strfmt text
	About string `json:"about"`

	// Почтовый адрес пользователя (уникальное поле).
	// example: captaina@blackpearl.sea
	// swagger:strfmt email
	// required: true
	Email string `json:"email"`
}
