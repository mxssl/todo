package handlers

import (
	"github.com/mxssl/todo/api/restapi/operations"
	"github.com/mxssl/todo/api/restapi/operations/todos"
	"github.com/mxssl/todo/store"
)

type handlers struct {
	itemStore *store.ItemStore
}

// Init handlers for API
func Init(itemStore *store.ItemStore, api *operations.TodoAPI) {
	h := &handlers{itemStore}
	api.TodosFindTodosHandler = todos.FindTodosHandlerFunc(h.FindTodosHandler)
	api.TodosAddOneHandler = todos.AddOneHandlerFunc(h.AddOneHandler)
	api.TodosUpdateOneHandler = todos.UpdateOneHandlerFunc(h.UpdateOneHandler)
	api.TodosDestroyOneHandler = todos.DestroyOneHandlerFunc(h.DestroyOneHandler)
}
