package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jnanendraveer/events-golang-app/api/models"
	"github.com/jnanendraveer/events-golang-app/api/utils/CommonFunction"
)

func (server *Server) CreatePendingOrderController(c *gin.Context) {
	var (
		err                error
		ResponseCode       int64
		pendingOrderModels models.PendingOrder
	)
	db := server.DB
	c.BindJSON(&pendingOrderModels)

	if err = pendingOrderModels.Validate(""); err != nil {
		ResponseCode = http.StatusPreconditionFailed
		c.JSON(int(ResponseCode), CommonFunction.Attachments(ResponseCode, err, ""))
		return
	}
	pendingOrderModels.FillRemainningPendingOrder()

	if _, err = pendingOrderModels.SavePendingOrders(db); err != nil {
		ResponseCode = http.StatusPreconditionFailed
		c.JSON(int(ResponseCode), CommonFunction.Attachments(ResponseCode, err, ""))
		return
	}

	c.JSON(http.StatusOK, pendingOrderModels)
	return
}
