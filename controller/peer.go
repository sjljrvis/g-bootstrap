package controller

import (
	"fmt"
	"net/http"

	Helper "github.com/sjljrvis/g-bootstrap/helpers"
	tools "github.com/sjljrvis/g-bootstrap/tools"
)

// GetAll controller
func GetAll(w http.ResponseWriter, r *http.Request) {
	val, err := tools.Redis.Get("key").Result()
	fmt.Println("key", val)
	if err != nil {
		Helper.RespondWithError(w, 204, err.Error()) // TODO: Do something about the error
	} else {
		Helper.RespondWithJSON(w, 200, nil)
	}
}
