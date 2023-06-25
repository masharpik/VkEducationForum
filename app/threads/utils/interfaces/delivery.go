package interfaces

import "net/http"

type IRouter interface {
	GetThreadInfo(http.ResponseWriter, *http.Request)
	UpdateThreadInfo(http.ResponseWriter, *http.Request)
	CreateThreadPosts(http.ResponseWriter, *http.Request)
	VoteThread(http.ResponseWriter, *http.Request)
	GetPosts(http.ResponseWriter, *http.Request)
}
