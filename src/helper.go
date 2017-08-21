package http_response_helper

import (
	"encoding/json"
	"net/http"
)

// ResponseHelper TODO
type ResponseHelper struct {
	ResultCode        int         `json:"resultCode,omitempty"`
	ResultDescription string      `json:"resultDescription,omitempty"`
	ResultData        interface{} `json:"resultData,omitempty"`
	Error             error       `json:"resultError,omitempty"`
}

// JSONResponse TODO
func JSONResponse(data interface{}, w http.ResponseWriter, e error) {
	httpCode, resultCode := handlerError(e)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(httpCode)
	resp := &ResponseHelper{
		ResultCode:        resultCode,
		ResultDescription: ResultDescription[resultCode],
		ResultData:        data,
	}
	json.NewEncoder(w).Encode(resp)
	return
}
