package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (server *Server) IndexController(c *gin.Context) {
	AddForm := []byte(`
	<html>
	<h3 style="top: 50%;
		text-align: center;
		margin-top: 10%;
		font-size: 90px;"><a href="https://fitpass.co.in" style="text-decoration: none"><span style="    color: #ef4c4f;">FITPASS</span></a></h3>
		</html>`)

	ContentTypeHTML := "text/html; charset=utf-8"
	c.Data(http.StatusOK, ContentTypeHTML, AddForm)
	return
}
