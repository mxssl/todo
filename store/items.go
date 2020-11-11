package store

import (
	"errors"

	"github.com/mxssl/todo/db"
)

// ItemStore ...
type ItemStore struct {
	db *db.DB
}

// NewItemStore ...
func NewItemStore(db *db.DB) *ItemStore {
	return &ItemStore{db}
}

// Item ...
type Item struct {
	Completed   bool   `db:"completed"`
	Description string `db:"description"`
	ID          int64  `db:"id"`
}

// Items ...
type Items []Item

// ErrNothingToDelete ...
var ErrNothingToDelete = errors.New("nothing to delete")

// ErrNothingToUpdate ...
var ErrNothingToUpdate = errors.New("nothing to update")

// AddItem adds item to db
func (is *ItemStore) AddItem(description string) (err error) {
	_, err = is.db.Exec(`INSERT INTO items (description) values ($1)`, description)
	return
}

// UpdateItemByID updates item in db
func (is *ItemStore) UpdateItemByID(id int64, description string, completed bool) (err error) {
	res, err := is.db.Exec(`UPDATE items set description=$1, completed=$2 where id=$3`, description, completed, id)
	if err != nil {
		return
	}

	// Check that id exists
	if ra, _ := res.RowsAffected(); ra < 1 {
		return ErrNothingToUpdate
	}

	return
}

// DeleteItemByID removes item from db
func (is *ItemStore) DeleteItemByID(id int64) (err error) {
	res, err := is.db.Exec(`DELETE from items where id=$1`, id)
	if err != nil {
		return
	}

	// Check that id exists
	if ra, _ := res.RowsAffected(); ra < 1 {
		return ErrNothingToDelete
	}

	return
}

// GetAllItems returns all items from db
func (is *ItemStore) GetAllItems() (it Items, err error) {
	err = is.db.Select(&it, `SELECT id, description, completed FROM items`)
	return
}
