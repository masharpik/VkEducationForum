package forum

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	mainLiterals "github.com/masharpik/ForumVKEducation/utils/literals"
	"github.com/masharpik/ForumVKEducation/utils/logger"
)

func StartServer(r *mux.Router) error {
	addr := fmt.Sprintf("%s:%s", os.Getenv("SERVER_APP_HOST"), os.Getenv("SERVER_APP_PORT"))

	logger.LogOperationSuccess(fmt.Sprintf(mainLiterals.LogServerWasStarted, addr))
	return http.ListenAndServe(addr, r)
}
