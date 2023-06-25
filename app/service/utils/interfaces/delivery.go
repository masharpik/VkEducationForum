package interfaces

import "net/http"

type IRouter interface {
	ClearDB(http.ResponseWriter, *http.Request)
	GetInfoDB(http.ResponseWriter, *http.Request)
}
