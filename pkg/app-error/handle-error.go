package apperror

import (
	"link-shortener/pkg/response"
	"net/http"
)

func HandleError(err error, w http.ResponseWriter) {
	log.Error("Error occured: ", err)

	w.Header().Set("Content-Type", "application/json")

	errRes := struct {
		Error string `json:"error"`
		Code  int    `json:"code"`
	}{
		Code:  http.StatusInternalServerError,
		Error: err.Error(),
	}

	if appErr, ok := err.(*AppError); ok {
		errRes.Error = appErr.Message
		errRes.Code = appErr.Code
	}

	response.Json(w, errRes.Code, errRes)
}
