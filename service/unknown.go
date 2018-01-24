package service

import (
	"net/http"
)

func NotImplemented(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "开发中", http.StatusNotImplemented)
}

func NotImplementedHandler() http.Handler {
	return http.HandlerFunc(NotImplemented)
}
