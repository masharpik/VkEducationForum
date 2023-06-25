package forumDelivery

import (
	"net/http"
	"strconv"

	"github.com/mailru/easyjson"
	forumLiterals "github.com/masharpik/ForumVKEducation/app/forums/utils/literals"
	forumStructs "github.com/masharpik/ForumVKEducation/app/forums/utils/structs"
	threadLiterals "github.com/masharpik/ForumVKEducation/app/threads/utils/literals"
	threadStructs "github.com/masharpik/ForumVKEducation/app/threads/utils/structs"
	userLiterals "github.com/masharpik/ForumVKEducation/app/users/utils/literals"
	"github.com/masharpik/ForumVKEducation/utils/getters"
	"github.com/masharpik/ForumVKEducation/utils/writer"
)

func (router router) CreateForum(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var forum forumStructs.Forum
	err := easyjson.UnmarshalFromReader(r.Body, &forum)
	if err != nil {
		writer.WriteErrorMessageRespond(w, r, http.StatusBadRequest, err.Error())
		return
	}

	createdForum, err := router.service.CreateForum(forum)
	if err != nil {
		var errStr string = err.Error()
		switch errStr {
		case forumLiterals.ForumAlreadyExists:
			writer.WriteJSONResponse(w, r, http.StatusConflict, createdForum)
			return
		case forumLiterals.UserFKNotFound:
			writer.WriteErrorMessageRespond(w, r, http.StatusNotFound, errStr)
			return
		}

		writer.WriteErrorMessageRespond(w, r, http.StatusInternalServerError, errStr)
		return
	}

	writer.WriteJSONResponse(w, r, http.StatusCreated, createdForum)
}

func (router router) GetForumInfo(w http.ResponseWriter, r *http.Request) {
	slug := getters.GetRequestVar(r, "slug")

	forum, err := router.service.GetForumInfo(slug)
	if err != nil {
		errStr := err.Error()

		switch errStr {
		case forumLiterals.ForumNotFound:
			writer.WriteErrorMessageRespond(w, r, http.StatusNotFound, errStr)
			return
		}

		writer.WriteErrorMessageRespond(w, r, http.StatusInternalServerError, errStr)
		return
	}

	writer.WriteJSONResponse(w, r, http.StatusOK, forum)
}

func (router router) CreateForumThread(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	slug := getters.GetRequestVar(r, "slug")

	var thread threadStructs.Thread
	err := easyjson.UnmarshalFromReader(r.Body, &thread)
	if err != nil {
		writer.WriteErrorMessageRespond(w, r, http.StatusBadRequest, err.Error())
		return
	}

	originSlug, err := router.service.GetForumSlugBySlug(slug)
	if err != nil {
		errStr := err.Error()

		if errStr == forumLiterals.ForumNotFound {
			writer.WriteErrorMessageRespond(w, r, http.StatusNotFound, errStr)
			return
		}
		writer.WriteErrorMessageRespond(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	thread.Forum = originSlug

	originNickname, err := router.userService.GetNickByNick(thread.Author)
	if err != nil {
		errStr := err.Error()
		if errStr == userLiterals.UserNotFound {
			writer.WriteErrorMessageRespond(w, r, http.StatusNotFound, errStr)
			return
		}

		writer.WriteErrorMessageRespond(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	thread.Author = originNickname

	createdThread, err := router.threadService.CreateThread(thread)
	if err != nil {
		errStr := err.Error()
		switch errStr {
		case threadLiterals.ThreadAlreadyExists:
			writer.WriteJSONResponse(w, r, http.StatusConflict, createdThread)
			return
		}

		writer.WriteErrorMessageRespond(w, r, http.StatusInternalServerError, errStr)
		return
	}

	writer.WriteJSONResponse(w, r, http.StatusCreated, createdThread)
}

func (router router) GetForumUsers(w http.ResponseWriter, r *http.Request) {
	slug := getters.GetRequestVar(r, "slug")

	params := r.URL.Query()

	var (
		limit int
		since string
		desc  bool
		err   error
	)

	if limitStr := params.Get("limit"); limitStr != "" {
		limit, err = strconv.Atoi(limitStr)
		if err != nil {
			writer.WriteErrorMessageRespond(w, r, http.StatusBadRequest, err.Error())
			return
		}
	} else {
		limit = 100
	}

	since = params.Get("since")

	if descStr := params.Get("desc"); descStr != "" {
		desc, err = strconv.ParseBool(descStr)
		if err != nil {
			writer.WriteErrorMessageRespond(w, r, http.StatusBadRequest, err.Error())
			return
		}
	} else {
		desc = false
	}

	users, err := router.service.GetForumUsers(slug, since, limit, desc)
	if err != nil {
		errStr := err.Error()
		if errStr == forumLiterals.ForumNotFound {
			writer.WriteErrorMessageRespond(w, r, http.StatusNotFound, errStr)
			return
		}

		writer.WriteErrorMessageRespond(w, r, http.StatusInternalServerError, errStr)
		return
	}

	writer.WriteJSONResponse(w, r, http.StatusOK, users)
}

func (router router) GetForumThreads(w http.ResponseWriter, r *http.Request) {
	slug := getters.GetRequestVar(r, "slug")

	forumExists, err := router.service.CheckExistenceForum(slug)
	if err != nil {
		writer.WriteErrorMessageRespond(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	if !forumExists {
		writer.WriteErrorMessageRespond(w, r, http.StatusNotFound, forumLiterals.ForumNotFound)
		return
	}

	params := r.URL.Query()

	var (
		limit int
		since string
		desc  bool
	)

	if limitStr := params.Get("limit"); limitStr != "" {
		limit, err = strconv.Atoi(limitStr)
		if err != nil {
			writer.WriteErrorMessageRespond(w, r, http.StatusBadRequest, err.Error())
			return
		}
	} else {
		limit = 100
	}

	since = params.Get("since")

	if descStr := params.Get("desc"); descStr != "" {
		desc, err = strconv.ParseBool(descStr)
		if err != nil {
			writer.WriteErrorMessageRespond(w, r, http.StatusBadRequest, err.Error())
			return
		}
	} else {
		desc = false
	}

	threads, err := router.threadService.GetForumThreads(slug, since, limit, desc)
	if err != nil {
		writer.WriteErrorMessageRespond(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	writer.WriteJSONResponse(w, r, http.StatusOK, threads)
}
