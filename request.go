package aefire

import (
	"net/http"
	"strings"
)

func PathSegments(r *http.Request) []string {
	return strings.Split(r.URL.Path, "/")
}

func LastPathSegments(r *http.Request) string {
	seg := PathSegments(r)
	return seg[len(seg)-1]
}
