package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jnanendraveer/transactions-golang-app/api/models"
	"github.com/jnanendraveer/transactions-golang-app/api/responses"
)

func (server *Server) WebEngageEventController(c *gin.Context) {
	var RequestData map[string]interface{}
	c.BindJSON(&RequestData)
	json.Unmarshal(models.WebEngageEvents(RequestData), &RequestData)
	responses.JSON(c, http.StatusOK, RequestData)
	return
}

// func (server *Server) WebEngageTransactionSuccessController(c *gin.Context) {
// 	var RequestData map[string]interface{}
// 	c.BindJSON(&RequestData)
// 	json.Unmarshal(CommonFunction.WebEngageTransaction(RequestData), &RequestData)
// 	responses.JSON(c, http.StatusOK, RequestData)
// 	return
// }

// func (server *Server) WebEngageTransactionFailController(c *gin.Context) {
// 	var RequestData map[string]interface{}
// 	c.BindJSON(&RequestData)
// 	json.Unmarshal(CommonFunction.WebEngageTransaction(RequestData), &RequestData)
// 	responses.JSON(c, http.StatusOK, RequestData)
// 	return
// }

func (server *Server) WebEngageController(c *gin.Context) {
	var (
		RequestData map[string]interface{}
		// db          = server.DB
	)

	c.BindJSON(&RequestData)

}
