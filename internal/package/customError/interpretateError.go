package customError

import (
	"net/http"

	"github.com/Kostikans/avitoTest/internal/package/clientError"
	"github.com/Kostikans/avitoTest/internal/package/serverError"
)

var convertStatusToHTTP = map[int]int{
	clientError.BadRequest:           http.StatusBadRequest,
	clientError.PaymentReq:           http.StatusPaymentRequired,
	clientError.Locked:               http.StatusLocked,
	clientError.Unauthorizied:        http.StatusUnauthorized,
	clientError.Conflict:             http.StatusConflict,
	clientError.Forbidden:            http.StatusForbidden,
	clientError.Gone:                 http.StatusGone,
	clientError.UnsupportedMediaType: http.StatusUnsupportedMediaType,
	serverError.ServerInternalError:  http.StatusInternalServerError,
	clientError.NotFound:             http.StatusNotFound,
	clientError.NotAccespteble:       http.StatusNotAcceptable,
}

func StatusCode(code int) int {
	return convertStatusToHTTP[code]
}