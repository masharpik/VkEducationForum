package userDelivery

import (
	"net/http"

	"github.com/mailru/easyjson"
	userLiterals "github.com/masharpik/ForumVKEducation/app/users/utils/literals"
	userStructs "github.com/masharpik/ForumVKEducation/app/users/utils/structs"
	"github.com/masharpik/ForumVKEducation/utils/getters"
	"github.com/masharpik/ForumVKEducation/utils/writer"
)

func (router *router) CreateUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var user userStructs.User
	err := easyjson.UnmarshalFromReader(r.Body, &user)
	if err != nil {
		writer.WriteErrorMessageRespond(w, r, http.StatusBadRequest, err.Error())
		return
	}

	user.Nickname = getters.GetRequestVar(r, "nickname")

	createdUser, err := router.service.CreateUser(user)
	if err == nil {
		writer.WriteJSONResponse(w, r, http.StatusCreated, createdUser)
		return
	}

	if err.Error() == userLiterals.LogUserAlreadyExists {
		createdUsers, err := router.service.GetUsersByNickOrMail(user)
		if err != nil {
			writer.WriteErrorMessageRespond(w, r, http.StatusInternalServerError, err.Error())
			return
		}

		writer.WriteJSONResponse(w, r, http.StatusConflict, createdUsers)
		return
	}

	writer.WriteErrorMessageRespond(w, r, http.StatusInternalServerError, err.Error())
}

func (router *router) GetUserInfo(w http.ResponseWriter, r *http.Request) {
	nickname := getters.GetRequestVar(r, "nickname")

	user, err := router.service.GetUserInfo(nickname)
	if err != nil {
		errStr := err.Error()

		switch errStr {
		case userLiterals.UserNotFound:
			writer.WriteErrorMessageRespond(w, r, http.StatusNotFound, errStr)
		default:
			writer.WriteErrorMessageRespond(w, r, http.StatusInternalServerError, errStr)
		}

		return
	}

	writer.WriteJSONResponse(w, r, http.StatusOK, user)
}

func (router *router) UpdateUserInfo(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var user userStructs.User
	err := easyjson.UnmarshalFromReader(r.Body, &user)
	if err != nil {
		writer.WriteErrorMessageRespond(w, r, http.StatusBadRequest, err.Error())
		return
	}

	user.Nickname = getters.GetRequestVar(r, "nickname")

	updatedUser, err := router.service.UpdateUserInfo(user)
	if err != nil {
		errStr := err.Error()

		switch errStr {
		case userLiterals.UserNotFound:
			writer.WriteErrorMessageRespond(w, r, http.StatusNotFound, errStr)
		case userLiterals.LogUserAlreadyExists:
			writer.WriteErrorMessageRespond(w, r, http.StatusConflict, errStr)
		default:
			writer.WriteErrorMessageRespond(w, r, http.StatusInternalServerError, errStr)
		}

		return
	}

	writer.WriteJSONResponse(w, r, http.StatusOK, updatedUser)
}
