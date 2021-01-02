package customError

import (
	"net/http"

	"github.com/Kostikans/avitoTest/internal/package/clientError"
	"github.com/Kostikans/avitoTest/internal/package/serverError"
)

var convertStatusToHTTP = map[int]int{
	clientError.BadRequest:          http.StatusBadRequest,
	serverError.ServerInternalError: http.StatusInternalServerError,
	clientError.NotFound:            http.StatusNotFound,
}

func StatusCode(code int) int {
	return convertStatusToHTTP[code]
}
