package handlers

import (
	log "github.com/sirupsen/logrus"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"github.com/mxssl/todo/api/models"
	"github.com/mxssl/todo/api/restapi/operations/todos"
)

func (h *handlers) AddOneHandler(params todos.AddOneParams) middleware.Responder {

	err := h.itemStore.AddItem(*params.Body.Description)
	if err != nil {
		log.Printf("cannot add item to db: %v", err)
		return todos.NewFindTodosDefault(500).WithPayload(&models.Error{
			Code:    500,
			Message: swag.String("cannot add item to db"),
		})
	}

	return todos.NewAddOneCreated()
}
