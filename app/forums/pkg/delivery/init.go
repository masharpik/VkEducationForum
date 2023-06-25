package forumDelivery

import (
	"errors"
	"fmt"

	"github.com/gorilla/mux"

	"github.com/masharpik/ForumVKEducation/app/forums/utils/interfaces"
	threadInterfaces "github.com/masharpik/ForumVKEducation/app/threads/utils/interfaces"
	userInterfaces "github.com/masharpik/ForumVKEducation/app/users/utils/interfaces"
	mainLiterals "github.com/masharpik/ForumVKEducation/utils/literals"
	"github.com/masharpik/ForumVKEducation/utils/logger"
)

type router struct {
	router        *mux.Router
	service       interfaces.IService
	threadService threadInterfaces.IService
	userService   userInterfaces.IService
}

func RegisterHandlers(r *mux.Router, service interfaces.IService, threadService threadInterfaces.IService, userService userInterfaces.IService) {
	router := router{
		router:        r,
		service:       service,
		threadService: threadService,
		userService:   userService,
	}

	_, ok := interface{}(router).(interfaces.IRouter)
	if !ok {
		logger.LogOperationFatal(errors.New(fmt.Sprintf(mainLiterals.LogStructNotSatisfyInterface, "forumRouter")))
	}

	router.router.HandleFunc("/create", router.CreateForum).Methods("POST")
	router.router.HandleFunc("/{slug}/details", router.GetForumInfo).Methods("GET")
	router.router.HandleFunc("/{slug}/create", router.CreateForumThread).Methods("POST")
	router.router.HandleFunc("/{slug}/users", router.GetForumUsers).Methods("GET")
	router.router.HandleFunc("/{slug}/threads", router.GetForumThreads).Methods("GET")
}
