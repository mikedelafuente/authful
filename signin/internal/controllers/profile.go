package controllers

import (
	"net/http"

	"github.com/mikedelafuente/authful-servertools/pkg/httptools"
)

func GetProfile(w http.ResponseWriter, r *http.Request) {
	httptools.HandleResponse(w, []byte{}, http.StatusOK)
}
