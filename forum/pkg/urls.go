package forum

import (
	"github.com/gorilla/mux"

	initDelivery "github.com/masharpik/ForumVKEducation/app/app/delivery"
)

func RegisterUrls() (r *mux.Router, err error) {
	r = mux.NewRouter()

	apiRouter := r.PathPrefix("/api").Subrouter()
	err = initDelivery.RegisterHandlers(apiRouter)
	if err != nil {
		return
	}

	return
}
