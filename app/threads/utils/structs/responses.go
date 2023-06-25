package threadStructs

import threadLiterals "github.com/masharpik/ForumVKEducation/app/threads/utils/literals"

//go:generate easyjson -all responses.go

// Ветка обсуждения на форуме
// swagger:model
type Thread struct {
	// Идентификатор ветки обсуждения.
	// example: 42
	// readOnly: true
	Id int32 `json:"id"`

	// Заголовок ветки обсуждения.
	// example: Davy Jones cache
	// required: true
	Title string `json:"title"`

	// Пользователь, создавший данную ветку обсуждения.
	// example: j.sparrow
	// required: true
	// swagger:strfmt identity
	Author string `json:"author"`

	// Форум, в котором расположена данная ветка обсуждения.
	// example: pirate-stories
	// required: true
	// swagger:strfmt identity
	Forum string `json:"forum"`

	// Описание ветки обсуждения.
	// example: An urgent need to reveal the hiding place of Davy Jones. Who is willing to help in this matter?
	// swagger:strfmt text
	// required: true
	Message string `json:"message"`

	// Кол-во голосов непосредственно за данное сообщение форума.
	// example: 34
	// readOnly: true
	Votes int32 `json:"votes"`

	// Человекопонятный URL (https://ru.wikipedia.org/wiki/%D0%A1%D0%B5%D0%BC%D0%B0%D0%BD%D1%82%D0%B8%D1%87%D0%B5%D1%81%D0%BA%D0%B8%D0%B9_URL).
	// В данной структуре slug опционален.
	// example: jones-cache
	// pattern: ^(\d|\w|-|_)*(\w|-|_)(\d|\w|-|_)*$
	// swagger:strfmt identity
	// readOnly: true
	Slug string `json:"slug"`

	// Дата создания ветки на форуме.
	// example: 2017-01-01T00:00:00.000Z
	// swagger:strfmt date-time
	// required: true
	Created string `json:"created"`
}

// Список веток обсуждения на форуме
// swagger:model
//
//easyjson:json
type Threads []Thread

// Данные для обновления ветки обсуждения на форуме
// Пустые параметры остаются без изменений
// swagger:model
type ThreadUpdate struct {
	// Заголовок ветки обсуждения.
	// example: Davy Jones cache
	Title string `json:"title"`

	// Описание ветки обсуждения.
	// example: An urgent need to reveal the hiding place of Davy Jones. Who is willing to help in this matter?
	// swagger:strfmt text
	Message string `json:"message"`
}

// Информация о голосовании пользователя
// swagger:model
type Vote struct {
	// Идентификатор пользователя.
	// required: true
	// swagger:strfmt identity
	Nickname string `json:"nickname"`

	// Отданный голос.
	// required: true
	Voice threadLiterals.Voice `json:"voice"`
}
