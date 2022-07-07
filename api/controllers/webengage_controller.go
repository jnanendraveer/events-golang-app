package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jnanendraveer/events-golang-app/api/models"
	"github.com/jnanendraveer/events-golang-app/api/responses"
	"github.com/jnanendraveer/events-golang-app/api/utils/CommonFunction"
)

func (server *Server) WebEngageTransactionEventController(c *gin.Context) {
	var RequestData map[string]interface{}
	c.BindJSON(&RequestData)
	json.Unmarshal(models.WebEngageEvents(RequestData), &RequestData)
	responses.JSON(c, http.StatusOK, RequestData)
	return
}

func (server *Server) WebEngageTransactionCreateController(c *gin.Context) {
	var (
		err                       error
		RequestData               map[string]interface{}
		db                        = server.DB
		CustomerPendingOrderModel models.CustomerPendingOrders
		ResponseCode              int64 = http.StatusOK
	)

	if err = c.BindJSON(&RequestData); err != nil {
		ResponseCode = http.StatusPreconditionFailed
		Attatchment := CommonFunction.Attachments(ResponseCode, nil, "")
		c.JSON(http.StatusPreconditionFailed, Attatchment)
		return
	}

	CustomerPendingOrderModel.FillRemainningCustomerPendingOrder(RequestData)
	CustomerPendingOrderModel.SaveCustomerPendingOrders(db)
	Attatchment := CommonFunction.Attachments(ResponseCode, nil, "")
	c.JSON(http.StatusOK, Attatchment)
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
