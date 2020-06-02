package handler

import (
	iqc "infoqerja-line/app/config"
	"net/http"
)

// IndexHandler will handle all request match the URL
type IndexHandler struct {
	config iqc.Config
}

// BuildIndexHandler return IndexHandler struct
func BuildIndexHandler(config iqc.Config) *IndexHandler {
	return &IndexHandler{config: config}
}

// Welcome is index page
func (h IndexHandler) Welcome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome!"))
}
