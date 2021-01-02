package customError

import (
	"errors"
	"net/http"

	"github.com/Kostikans/avitoTest/internal/package/responses"

	"github.com/Kostikans/avitoTest/internal/package/logger"
)

func PostError(w http.ResponseWriter, req *http.Request, log *logger.CustomLogger, err error, code interface{}) {
	if code != nil {
		err = NewCustomError(err, code.(int), 2)
	}

	log.LogError(req.Context(), err)
	responses.SendErrorResponse(w, StatusCode(ParseCode(err)), errors.Unwrap(err).Error())
}
