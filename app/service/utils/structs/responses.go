package serviceStructs
//go:generate easyjson -all responses.go

// Описывает данные о базе данных
// swagger:model
type Status struct {
	// Кол-во пользователей в базе данных.
    // example: 1000
	// required: true
	User   int32 `json:"user"`

	// Кол-во разделов в базе данных.
    // example: 100
	// required: true
	Forum  int32 `json:"forum"`

	// Кол-во веток обсуждения в базе данных.
    // example: 1000
	// required: true
	Thread int32 `json:"thread"`

	// Кол-во сообщений в базе данных.
    // example: 1000000
	// required: true
	Post   int64 `json:"post"`
}
