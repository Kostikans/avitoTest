package responses

import (
	"encoding/json"
	"net/http"

	"github.com/Kostikans/avitoTest/internal/package/clientError"

	"github.com/mailru/easyjson"
)

func SendDataResponse(w http.ResponseWriter, data interface{}) {
	response := HttpResponse{Data: data, Code: clientError.OkResponse}
	_, _, err := easyjson.MarshalToHTTPResponseWriter(response, w)
	if err != nil {
		SendErrorResponse(w, http.StatusInternalServerError)
	}
}

func SendOkResponse(w http.ResponseWriter) {
	err := json.NewEncoder(w).Encode(HttpResponse{Code: clientError.OkResponse})
	if err != nil {
		SendErrorResponse(w, http.StatusInternalServerError)
	}
}
