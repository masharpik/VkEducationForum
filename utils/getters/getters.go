package getters

import (
	"net/http"

	"github.com/gorilla/mux"
)

func GetRequestVar(r *http.Request, varName string) string {
	vars := mux.Vars(r)

	return vars[varName]
}
