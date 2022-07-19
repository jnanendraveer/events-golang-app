package controllers

import "github.com/jnanendraveer/events-golang-app/api/middlewares"

func (s *Server) initializeRoutes() {
	s.Router.GET("/", s.IndexController)
	s.Router.POST("transaction/event", middlewares.SetMiddlewareAuthentication(), s.WebEngageTransactionEventController)
	s.Router.POST("transaction/create", middlewares.SetMiddlewareAuthentication(), s.WebEngageTransactionCreateController)
	s.Router.POST("/pending-order/create", middlewares.SetMiddlewareAuthentication(), s.CreatePendingOrderController)
	s.Router.POST("/billdesk/order-create", middlewares.SetMiddlewareAuthentication(), s.BillDeskController)

	// s.Router.POST("transaction/success", middlewares.SetMiddlewareAuthentication(), s.WebEngageTransactionSucessController)
	// s.Router.POST("transaction/failed", middlewares.SetMiddlewareAuthentication(), s.WebEngageTransactionFailController)
}
