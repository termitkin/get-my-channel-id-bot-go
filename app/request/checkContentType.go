package request

import (
	"fmt"
	"net/http"
)

func CheckContentType(res http.ResponseWriter, req *http.Request) bool {
	contentType := req.Header.Get("Content-Type")

	if contentType != "application/json" {
		res.WriteHeader(http.StatusUnsupportedMediaType)
		_, err := res.Write([]byte("415 - Unsupported Media Type. Please send JSON"))

		if err != nil {
			fmt.Println(err)
		}

		return false
	}

	return true
}
