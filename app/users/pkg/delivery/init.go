package userDelivery

import (
	"errors"
	"fmt"

	"github.com/gorilla/mux"

	"github.com/masharpik/ForumVKEducation/app/users/utils/interfaces"
	mainLiterals "github.com/masharpik/ForumVKEducation/utils/literals"
	"github.com/masharpik/ForumVKEducation/utils/logger"
)

type router struct {
	router  *mux.Router
	service interfaces.IService
}

func RegisterHandlers(r *mux.Router, service interfaces.IService) {
	router := router{
		router:  r,
		service: service,
	}

	_, ok := interface{}(&router).(interfaces.IRouter)
	if !ok {
		logger.LogOperationFatal(errors.New(fmt.Sprintf(mainLiterals.LogStructNotSatisfyInterface, "userRouter")))
	}

	router.router.HandleFunc("/{nickname}/create", router.CreateUser).Methods("POST")
	router.router.HandleFunc("/{nickname}/profile", router.GetUserInfo).Methods("GET")
	router.router.HandleFunc("/{nickname}/profile", router.UpdateUserInfo).Methods("POST")
}
