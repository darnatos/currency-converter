package objects

import (
	"encoding/json"
	"net/http"
)

type ResponseWrapper struct {
	Code   int    `json:"code"`
	Result string `json:"result"`
}

// Json convert ResponseWrapper in json
func (e *ResponseWrapper) Json() []byte {
	if e == nil {
		return []byte("{}")
	}
	res, _ := json.Marshal(e)
	return res
}

// StatusCode return status code
func (e *ResponseWrapper) StatusCode() int {
	if e == nil || e.Code == 0 {
		return http.StatusOK
	}
	return e.Code
}
