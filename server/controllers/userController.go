package controllers

import (
	"net/http"
)

type UserController struct {
	path map[string]map[string]func(http.ResponseWriter, *http.Request)
}

func (u *UserController) add(method, pattern string, fn func(http.ResponseWriter, *http.Request)) {
	if u.path == nil {
		u.path = make(map[string]map[string]func(http.ResponseWriter, *http.Request), 10)
	}

	if u.path[method] == nil {
		u.path[method] = make(map[string]func(http.ResponseWriter, *http.Request))
	}
	u.path[method][pattern] = fn

}

func (u *UserController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	uri := r.URL.Path
	method := r.Method
	if f, ok := u.path[method][uri]; ok {
		f(w, r)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}
