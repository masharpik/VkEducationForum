package interfaces

import "net/http"

type IRouter interface {
	GetPostInfo(http.ResponseWriter, *http.Request)
	UpdatePostInfo(http.ResponseWriter, *http.Request)
}
