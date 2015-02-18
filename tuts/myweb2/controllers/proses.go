package controllers

import (
	"amelia/gomvc"
	"net/http"
)

type ProsesController struct {
	gomvc.Controller
}

func (c *ProsesController) Project1(rw http.ResponseWriter, req *http.Request) {
	req.ParseForm()
}
