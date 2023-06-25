package postStructs

//go:generate easyjson -all responses.go

import (
	forumStructs "github.com/masharpik/ForumVKEducation/app/forums/utils/structs"
	threadStructs "github.com/masharpik/ForumVKEducation/app/threads/utils/structs"
	userStructs "github.com/masharpik/ForumVKEducation/app/users/utils/structs"
)

// Сообщение внутри ветки обсуждения на форуме
// swagger:model
type Post struct {
	// Идентификатор данного сообщения.
	// readOnly: true
	Id int64 `json:"id"`

	// Идентификатор родительского сообщения (0 - корневое сообщение обсуждения).
	Parent int64 `json:"parent"`

	// Автор, написавший данное сообщение.
	// example: j.sparrow
	// required: true
	// swagger:strfmt identity
	Author string `json:"author"`

	// Собственно сообщение форума.
	// example: We should be afraid of the Kraken.
	// swagger:strfmt text
	// required: true
	Message string `json:"message"`

	// Истина, если данное сообщение было изменено.
	// readOnly: true
	IsEdited bool `json:"isEdited"`

	// Идентификатор форума (slug) данного сообещния.
	// readOnly: true
	// swagger:strfmt identity
	Forum string `json:"forum"`

	// Идентификатор ветви (id) обсуждения данного сообещния.
	// readOnly: true
	Thread int32 `json:"thread"`

	// Дата создания сообщения на форуме.
	// example: 2017-01-01T00:00:00.000Z
	// swagger:strfmt date-time
	// readOnly: true
	Created string `json:"created"`
}

// Список сообщений в ветке обсуждения
// swagger:model
//
//easyjson:json
type Posts []Post

// Полная информация о сообщении, включая связанные объекты
// swagger:model
type PostFull struct {
	Post   Post                  `json:"post"`
	Author *userStructs.User     `json:"author,omitempty"`
	Thread *threadStructs.Thread `json:"thread,omitempty"`
	Forum  *forumStructs.Forum   `json:"forum,omitempty"`
}
