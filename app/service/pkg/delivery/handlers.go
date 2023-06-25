package serviceDelivery

import (
	"net/http"

	"github.com/masharpik/ForumVKEducation/utils/writer"
)

func (router router) ClearDB(w http.ResponseWriter, r *http.Request) {
	err := router.service.ClearDB()
	if err != nil {
		writer.WriteErrorMessageRespond(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	writer.WriteJSONResponse(w, r, http.StatusOK, nil)
}

func (router router) GetInfoDB(w http.ResponseWriter, r *http.Request) {
	status, err := router.service.GetInfoDB()
	if err != nil {
		writer.WriteErrorMessageRespond(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	writer.WriteJSONResponse(w, r, http.StatusOK, status)
}
