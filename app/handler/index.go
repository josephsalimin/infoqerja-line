package handler

import (
	iqc "infoqerja-line/app/config"
	"net/http"
)

// IndexHandler will handle all request match the URL
type IndexHandler struct {
	Config iqc.Config
}

// Welcome is index page
func (h IndexHandler) Welcome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome!"))
}
