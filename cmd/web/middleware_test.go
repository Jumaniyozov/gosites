package main

import (
	"fmt"
	"github.com/jumaniyozov/gosites/pkg/middlewares"
	"net/http"
	"testing"
)


func TestNoSurf(t *testing.T) {
	var myH myHandler

	h := middlewares.MiddlewareRepo.NoSurf(&myH)

	switch v := h.(type) {
	case http.Handler:
	default:
		t.Error(fmt.Sprintf("type is not http.Handler, but is %T", v))
	}
}

func TestSessionLoad(t *testing.T) {
	var myH myHandler

	h := middlewares.MiddlewareRepo.SessionLoad(&myH)

	switch v := h.(type) {
	case http.Handler:
	default:
		t.Error(fmt.Sprintf("type is not http.Handler, but is %T", v))
	}
}

