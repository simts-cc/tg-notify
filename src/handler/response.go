package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// APIResponse 回傳成功結果
type APIResponse struct {
	Ok bool
}

// Response 送出json
func (j APIResponse) Response(w *http.ResponseWriter) []byte {
	res, _ := json.Marshal(j)
	(*w).Header().Set("Content-Type", "application/json")
	fmt.Fprint(*w, string(res))

	return res
}

// APIErrorResponse 回傳失敗結果
type APIErrorResponse struct {
	Ok      bool
	Message string
}

// Response 送出json
func (j APIErrorResponse) Response(w *http.ResponseWriter) []byte {
	res, _ := json.Marshal(j)
	(*w).Header().Set("Content-Type", "application/json")
	fmt.Fprint(*w, string(res))

	return res
}
