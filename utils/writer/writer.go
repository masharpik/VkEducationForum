package writer

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/mailru/easyjson"
	mainErrors "github.com/masharpik/ForumVKEducation/utils/errors"
	mainLiterals "github.com/masharpik/ForumVKEducation/utils/literals"
	"github.com/masharpik/ForumVKEducation/utils/logger"
)

func WriteErrorMessageRespond(w http.ResponseWriter, r *http.Request, statusCode int, message string) {
	responseErr := mainErrors.New(message)

	w.WriteHeader(statusCode)
	started, _, err := easyjson.MarshalToHTTPResponseWriter(responseErr, w)

	if !logger.DebugOn {
		return
	}

	if !started {
		errorMsg := fmt.Errorf(mainLiterals.LogErrorOccurredBeforeResponseWriterMethods, err).Error()
		logger.LogRequestError(r, http.StatusInternalServerError, errors.New(errorMsg))
		return
	} else if err != nil {
		logger.LogRequestError(r, http.StatusInternalServerError, err)
		return
	}

	logger.LogRequestError(r, statusCode, errors.New(message))
}

func WriteErrorJSONRespond(w http.ResponseWriter, r *http.Request, statusCode int, responseJSON easyjson.Marshaler, message string) {
	w.WriteHeader(statusCode)
	started, _, err := easyjson.MarshalToHTTPResponseWriter(responseJSON, w)

	if !logger.DebugOn {
		return
	}

	if !started {
		errorMsg := fmt.Errorf(mainLiterals.LogErrorOccurredBeforeResponseWriterMethods, err).Error()
		logger.LogRequestError(r, http.StatusInternalServerError, errors.New(errorMsg))
		return
	} else if err != nil {
		logger.LogRequestError(r, http.StatusInternalServerError, err)
		return
	}

	logger.LogRequestError(r, statusCode, errors.New(message))
}

func WriteJSONResponse(w http.ResponseWriter, r *http.Request, statusCode int, responseJSON easyjson.Marshaler) {
	w.WriteHeader(statusCode)
	started, _, err := easyjson.MarshalToHTTPResponseWriter(responseJSON, w)
	if !logger.DebugOn {
		return
	}

	if !started {
		errorMsg := fmt.Errorf(mainLiterals.LogErrorOccurredBeforeResponseWriterMethods, err).Error()
		logger.LogRequestError(r, http.StatusInternalServerError, errors.New(errorMsg))
		return
	} else if err != nil {
		logger.LogRequestError(r, http.StatusInternalServerError, err)
		return
	}

	logger.LogRequestSuccess(r, statusCode)
}
