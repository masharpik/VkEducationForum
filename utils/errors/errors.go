package mainErrors

//go:generate easyjson -all errors.go

// Кастомная ошибка
// swagger:model
type Error struct {
	// Текстовое описание ошибки.
	// readOnly: true
	// example: Can't find user with id #42
	Message string `json:"message"`
}

func (err *Error) Error() string {
	return err.Message
}

func New(msg string) *Error {
	return &Error{
		Message: msg,
	}
}
