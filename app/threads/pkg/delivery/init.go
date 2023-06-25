package threadDelivery

import (
	"errors"
	"fmt"

	"github.com/gorilla/mux"

	postInterfaces "github.com/masharpik/ForumVKEducation/app/posts/utils/interfaces"
	"github.com/masharpik/ForumVKEducation/app/threads/utils/interfaces"
	mainLiterals "github.com/masharpik/ForumVKEducation/utils/literals"
	"github.com/masharpik/ForumVKEducation/utils/logger"
)

type router struct {
	router      *mux.Router
	service     interfaces.IService
	postService postInterfaces.IService
}

func RegisterHandlers(r *mux.Router, service interfaces.IService, postService postInterfaces.IService) {
	router := router{
		router:      r,
		service:     service,
		postService: postService,
	}

	_, ok := interface{}(router).(interfaces.IRouter)
	if !ok {
		logger.LogOperationFatal(errors.New(fmt.Sprintf(mainLiterals.LogStructNotSatisfyInterface, "threadRouter")))
	}

	router.router.HandleFunc("/{slug_or_id}/create", router.CreateThreadPosts).Methods("POST")
	router.router.HandleFunc("/{slug_or_id}/details", router.GetThreadInfo).Methods("GET")
	router.router.HandleFunc("/{slug_or_id}/details", router.UpdateThreadInfo).Methods("POST")
	router.router.HandleFunc("/{slug_or_id}/posts", router.GetPosts).Methods("GET")
	router.router.HandleFunc("/{slug_or_id}/vote", router.VoteThread).Methods("POST")
}
