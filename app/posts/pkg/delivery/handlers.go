package postDelivery

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/mailru/easyjson"
	postLiterals "github.com/masharpik/ForumVKEducation/app/posts/utils/literals"
	postStructs "github.com/masharpik/ForumVKEducation/app/posts/utils/structs"
	"github.com/masharpik/ForumVKEducation/utils/getters"
	"github.com/masharpik/ForumVKEducation/utils/writer"
)

func (router router) GetPostInfo(w http.ResponseWriter, r *http.Request) {
	idStr := getters.GetRequestVar(r, "id")

	idInt64, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		writer.WriteErrorMessageRespond(w, r, http.StatusBadRequest, err.Error())
		return
	}

	id := int(idInt64)

	relatedSlice := strings.Split(r.URL.Query().Get("related"), ",")
	var (
		relatedUser   bool
		relatedForum  bool
		relatedThread bool
	)
	for _, related := range relatedSlice {
		switch related {
		case "forum":
			relatedForum = true
		case "thread":
			relatedThread = true
		case "user":
			relatedUser = true
		}
	}

	post, err := router.service.GetPostInfo(id, relatedUser, relatedForum, relatedThread)
	if err != nil {
		errStr := err.Error()
		if errStr == postLiterals.PostNotFound {
			writer.WriteErrorMessageRespond(w, r, http.StatusNotFound, errStr)
			return
		}

		writer.WriteErrorMessageRespond(w, r, http.StatusInternalServerError, errStr)
		return
	}

	writer.WriteJSONResponse(w, r, http.StatusOK, post)
}

func (router router) UpdatePostInfo(w http.ResponseWriter, r *http.Request) {
	idStr := getters.GetRequestVar(r, "id")

	idInt64, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		writer.WriteErrorMessageRespond(w, r, http.StatusBadRequest, err.Error())
		return
	}

	id := int(idInt64)

	var post postStructs.Post
	defer r.Body.Close()
	err = easyjson.UnmarshalFromReader(r.Body, &post)
	if err != nil {
		writer.WriteErrorMessageRespond(w, r, http.StatusBadRequest, err.Error())
		return
	}

	createdPost, err := router.service.UpdatePostInfo(id, post.Message)
	if err != nil {
		errStr := err.Error()
		if errStr == postLiterals.PostNotFound {
			writer.WriteErrorMessageRespond(w, r, http.StatusNotFound, errStr)
			return
		}

		writer.WriteErrorMessageRespond(w, r, http.StatusInternalServerError, errStr)
		return
	}

	writer.WriteJSONResponse(w, r, http.StatusOK, createdPost)
}
