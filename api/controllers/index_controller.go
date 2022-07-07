package controllers

import (
	"fmt"
	"net/http"
)

func (server *Server) IndexController(w http.ResponseWriter, r *http.Request) {
	const AddForm = `
	<html>
	<h3 style="top: 50%;
		text-align: center;
		margin-top: 10%;
		font-size: 90px;"><a href="https://fitpass.co.in" style="text-decoration: none"><span style="    color: #ef4c4f;">FITPASS</span></a></h3>
		</html>`

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, AddForm)
	return
}
