package web

import (
	"net/http"

	"github.com/mikedelafuente/authful/servertools"
)

func GetProfile(w http.ResponseWriter, r *http.Request) {
	servertools.HandleResponse(w, []byte{}, http.StatusOK)
}
