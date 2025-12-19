package handler

import (
	"session-19/service"
)

type Handler struct {
	HandlerAuth       AuthHandler
	HandlerMenu       MenuHandler
	AssignmentHandler AssignmentHandler
}

func NewHandler(service service.Service) Handler {
	return Handler{
		// HandlerAuth:       NewAuthHandler(service.AuthService),
		// HandlerMenu:       NewMenuHandler(),
		AssignmentHandler: NewAssignmentHandler(service.AssignmentService),
	}
}
