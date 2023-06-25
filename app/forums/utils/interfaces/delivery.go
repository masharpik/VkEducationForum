package interfaces

import "net/http"

type IRouter interface {
	CreateForum(http.ResponseWriter, *http.Request)
	GetForumInfo(http.ResponseWriter, *http.Request)
	CreateForumThread(http.ResponseWriter, *http.Request)
	GetForumUsers(http.ResponseWriter, *http.Request)
	GetForumThreads(http.ResponseWriter, *http.Request)
}
