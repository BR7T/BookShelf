package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ParseJsonBody(w http.ResponseWriter, r *http.Request , dst interface{})error{
	content := r.Header.Get("Content-Type")
	// Verify type json
	if content != "application/json"{
		return fmt.Errorf("content-type must be application/json, got %s", content)
	}
	//  Limit json size
	r.Body = http.MaxBytesReader(w , r.Body , 1048576)

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(dst)
	if err != nil {
        return err
    }
	return nil
}