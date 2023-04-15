package request

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetJsonBody(res http.ResponseWriter, req *http.Request) map[string]any {
	b, err3 := io.ReadAll(req.Body)

	if err3 != nil {
		fmt.Println(err3)
	}

	var result map[string]any
	json.Unmarshal([]byte(b), &result)

	return result
}
