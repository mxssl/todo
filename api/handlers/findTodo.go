package handlers

import (
	log "github.com/sirupsen/logrus"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"github.com/mxssl/todo/api/models"
	"github.com/mxssl/todo/api/restapi/operations/todos"
)

func (h *handlers) FindTodosHandler(params todos.FindTodosParams) middleware.Responder {
	it, err := h.itemStore.GetAllItems()
	if err != nil {
		log.Printf("cannot get items from db: %v", err)
		return todos.NewFindTodosDefault(500).WithPayload(&models.Error{
			Code:    500,
			Message: swag.String("cannot get items from db"),
		})
	}

	var items []*models.Item

	for _, item := range it {
		items = append(items, &models.Item{
			ID:          item.ID,
			Description: &item.Description,
			Completed:   item.Completed,
		})
	}

	return todos.NewFindTodosOK().WithPayload(items)
}
