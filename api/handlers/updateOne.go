package handlers

import (
	log "github.com/sirupsen/logrus"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"github.com/mxssl/todo/api/models"
	"github.com/mxssl/todo/api/restapi/operations/todos"
	"github.com/mxssl/todo/store"
)

func (h *handlers) UpdateOneHandler(params todos.UpdateOneParams) middleware.Responder {
	err := h.itemStore.UpdateItemByID(params.ID, *params.Body.Description, params.Body.Completed)
	if err == store.ErrNothingToUpdate {
		log.Printf("item id %d not found: %v", params.ID, err)
		return todos.NewFindTodosDefault(404).WithPayload(&models.Error{
			Code:    404,
			Message: swag.String("item not found"),
		})
	} else if err != nil {
		log.Printf("cannot update item %d: %v", params.ID, err)
		return todos.NewFindTodosDefault(500).WithPayload(&models.Error{
			Code:    500,
			Message: swag.String("cannot update item"),
		})
	}

	return todos.NewUpdateOneOK()
}
