package threadDelivery

import (
	"net/http"
	"strconv"

	"github.com/mailru/easyjson"
	forumLiterals "github.com/masharpik/ForumVKEducation/app/forums/utils/literals"
	postLiterals "github.com/masharpik/ForumVKEducation/app/posts/utils/literals"
	postStructs "github.com/masharpik/ForumVKEducation/app/posts/utils/structs"
	threadLiterals "github.com/masharpik/ForumVKEducation/app/threads/utils/literals"
	threadStructs "github.com/masharpik/ForumVKEducation/app/threads/utils/structs"
	"github.com/masharpik/ForumVKEducation/utils/getters"
	"github.com/masharpik/ForumVKEducation/utils/writer"
)

func (router router) GetThreadInfo(w http.ResponseWriter, r *http.Request) {
	slugOrId := getters.GetRequestVar(r, "slug_or_id")
	var thread threadStructs.Thread

	idInt64, err := strconv.ParseInt(slugOrId, 10, 64)
	if err != nil {
		thread, err = router.service.GetThreadInfoBySlug(slugOrId)
	} else {
		thread, err = router.service.GetThreadInfoById(int(idInt64))
	}

	if err != nil {
		errStr := err.Error()
		if errStr == threadLiterals.ThreadNotFound || errStr == forumLiterals.ForumNotFound {
			writer.WriteErrorMessageRespond(w, r, http.StatusNotFound, errStr)
			return
		}

		writer.WriteErrorMessageRespond(w, r, http.StatusInternalServerError, errStr)
		return
	}

	writer.WriteJSONResponse(w, r, http.StatusOK, thread)
}

func (router router) UpdateThreadInfo(w http.ResponseWriter, r *http.Request) {
	slugOrId := getters.GetRequestVar(r, "slug_or_id")

	var thread threadStructs.Thread
	defer r.Body.Close()
	err := easyjson.UnmarshalFromReader(r.Body, &thread)

	var updatedThread threadStructs.Thread
	idInt64, err := strconv.ParseInt(slugOrId, 10, 64)
	if err != nil {
		updatedThread, err = router.service.UpdateThreadInfoBySlug(slugOrId, thread.Title, thread.Message)
	} else {
		updatedThread, err = router.service.UpdateThreadInfoById(int(idInt64), thread.Title, thread.Message)
	}

	if err != nil {
		errStr := err.Error()
		if errStr == threadLiterals.ThreadNotFound {
			writer.WriteErrorMessageRespond(w, r, http.StatusNotFound, errStr)
			return
		}

		writer.WriteErrorMessageRespond(w, r, http.StatusInternalServerError, errStr)
		return
	}

	writer.WriteJSONResponse(w, r, http.StatusOK, updatedThread)
}

func (router router) CreateThreadPosts(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var posts postStructs.Posts
	err := easyjson.UnmarshalFromReader(r.Body, &posts)
	if err != nil {
		writer.WriteErrorMessageRespond(w, r, http.StatusBadRequest, err.Error())
		return
	}

	slugOrId := getters.GetRequestVar(r, "slug_or_id")
	var id int

	idInt64, err := strconv.ParseInt(slugOrId, 10, 64)
	if err != nil {
		id, err = router.service.GetThreadIdBySlug(slugOrId)
		if err != nil {
			errStr := err.Error()
			if errStr == threadLiterals.ThreadNotFound {
				writer.WriteErrorMessageRespond(w, r, http.StatusNotFound, errStr)
				return
			}

			writer.WriteErrorMessageRespond(w, r, http.StatusInternalServerError, errStr)
			return
		}
	} else {
		id = int(idInt64)
	}

	forumSlug, err := router.service.GetForumSlugByThreadId(id)
	if err != nil {
		errStr := err.Error()
		if errStr == threadLiterals.ThreadNotFound {
			writer.WriteErrorMessageRespond(w, r, http.StatusNotFound, errStr)
			return
		}

		writer.WriteErrorMessageRespond(w, r, http.StatusInternalServerError, errStr)
		return
	}

	createdPosts, err := router.postService.CreatePostsByThreadId(id, forumSlug, posts)
	if err != nil {
		errStr := err.Error()
		if errStr == postLiterals.ParentNotInThread {
			writer.WriteErrorMessageRespond(w, r, http.StatusConflict, errStr)
			return
		}
		if errStr == postLiterals.ParentOrAuthorFKError {
			writer.WriteErrorMessageRespond(w, r, http.StatusNotFound, errStr)
			return
		}

		writer.WriteErrorMessageRespond(w, r, http.StatusInternalServerError, errStr)
		return
	}

	writer.WriteJSONResponse(w, r, http.StatusCreated, createdPosts)
}

func (router router) VoteThread(w http.ResponseWriter, r *http.Request) {
	slugOrId := getters.GetRequestVar(r, "slug_or_id")
	var id int

	idInt64, err := strconv.ParseInt(slugOrId, 10, 64)
	if err != nil {
		id, err = router.service.GetThreadIdBySlug(slugOrId)
		if err != nil {
			errStr := err.Error()
			if errStr == threadLiterals.ThreadNotFound {
				writer.WriteErrorMessageRespond(w, r, http.StatusNotFound, errStr)
				return
			}

			writer.WriteErrorMessageRespond(w, r, http.StatusInternalServerError, errStr)
			return
		}
	} else {
		id = int(idInt64)
	}

	var vote threadStructs.Vote
	defer r.Body.Close()
	err = easyjson.UnmarshalFromReader(r.Body, &vote)
	if err != nil {
		writer.WriteErrorMessageRespond(w, r, http.StatusBadRequest, err.Error())
		return
	}

	thread, err := router.service.VoteByThreadId(id, vote.Nickname, int(vote.Voice))
	if err != nil {
		errStr := err.Error()
		if errStr == threadLiterals.ForumHasNotThread {
			writer.WriteErrorMessageRespond(w, r, http.StatusNotFound, errStr)
			return
		}

		writer.WriteErrorMessageRespond(w, r, http.StatusInternalServerError, errStr)
		return
	}

	writer.WriteJSONResponse(w, r, http.StatusOK, thread)
}

func (router router) GetPosts(w http.ResponseWriter, r *http.Request) {
	slugOrId := getters.GetRequestVar(r, "slug_or_id")
	var id int

	idInt64, err := strconv.ParseInt(slugOrId, 10, 64)
	if err != nil {
		id, err = router.service.GetThreadIdBySlug(slugOrId)
		if err != nil {
			errStr := err.Error()
			if errStr == threadLiterals.ThreadNotFound {
				writer.WriteErrorMessageRespond(w, r, http.StatusNotFound, errStr)
				return
			}

			writer.WriteErrorMessageRespond(w, r, http.StatusInternalServerError, errStr)
			return
		}
	} else {
		id = int(idInt64)

		exists, err := router.service.CheckExistenceThreadById(id)
		if err != nil {
			writer.WriteErrorMessageRespond(w, r, http.StatusInternalServerError, err.Error())
			return
		}

		if !exists {
			writer.WriteErrorMessageRespond(w, r, http.StatusNotFound, threadLiterals.ThreadNotFound)
			return
		}
	}

	params := r.URL.Query()

	var (
		limit int
		since int
		sort  string
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

	if sinceStr := params.Get("since"); sinceStr != "" {
		since, err = strconv.Atoi(sinceStr)
		if err != nil {
			writer.WriteErrorMessageRespond(w, r, http.StatusBadRequest, err.Error())
			return
		}
	} else {
		since = 0
	}

	sort = params.Get("sort")
	if sort == "" {
		sort = "flat"
	}

	if descStr := params.Get("desc"); descStr != "" {
		desc, err = strconv.ParseBool(descStr)
		if err != nil {
			writer.WriteErrorMessageRespond(w, r, http.StatusBadRequest, err.Error())
			return
		}
	} else {
		desc = false
	}

	posts, err := router.postService.GetPostsByThreadId(id, since, limit, sort, desc)
	if err != nil {
		writer.WriteErrorMessageRespond(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	writer.WriteJSONResponse(w, r, http.StatusOK, posts)
}
