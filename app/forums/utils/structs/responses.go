package forumStructs

//go:generate easyjson -all responses.go

// Информация о форуме
// swagger:model
type Forum struct {
	// Название форума.
	// example: Pirate stories
	// required: true
	Title string `json:"title"`

	// Nickname пользователя, который отвечает за форум.
	// example: j.sparrow
	// required: true
	// swagger:strfmt identity
	User string `json:"user"`

	// Человекопонятный URL (https://ru.wikipedia.org/wiki/%D0%A1%D0%B5%D0%BC%D0%B0%D0%BD%D1%82%D0%B8%D1%87%D0%B5%D1%81%D0%BA%D0%B8%D0%B9_URL), уникальное поле.
	// example: pirate-stories
	// required: true
	// pattern: ^(\d|\w|-|_)*(\w|-|_)(\d|\w|-|_)*$
	// swagger:strfmt identity
	Slug string `json:"slug"`

	// Общее кол-во сообщений в данном форуме.
	// example: 200000
	// readOnly: true
	Posts int `json:"posts,omitempty"`

	// Общее кол-во ветвей обсуждения в данном форуме.
	// example: 200
	// readOnly: true
	Threads int `json:"threads,omitempty"`
}
