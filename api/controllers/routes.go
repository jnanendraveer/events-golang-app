package controllers

import "github.com/jnanendraveer/transactions-golang-app/api/middlewares"

func (s *Server) initializeRoutes() {
	s.Router.POST("transaction/event", middlewares.SetMiddlewareAuthentication(), s.WebEngageEventController)
	// s.Router.POST("transaction/success", middlewares.SetMiddlewareAuthentication(), s.WebEngageTransactionSucessController)
	// s.Router.POST("transaction/failed", middlewares.SetMiddlewareAuthentication(), s.WebEngageTransactionFailController)
}
