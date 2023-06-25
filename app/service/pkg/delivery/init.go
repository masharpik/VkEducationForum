package serviceDelivery

import (
	"errors"
	"fmt"

	"github.com/gorilla/mux"

	"github.com/masharpik/ForumVKEducation/app/service/utils/interfaces"
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

	_, ok := interface{}(router).(interfaces.IRouter)
	if !ok {
		logger.LogOperationFatal(errors.New(fmt.Sprintf(mainLiterals.LogStructNotSatisfyInterface, "serviceRouter")))
	}

	router.router.HandleFunc("/clear", router.ClearDB).Methods("POST")
	router.router.HandleFunc("/status", router.GetInfoDB).Methods("GET")
}
