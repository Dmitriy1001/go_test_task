package event

import (
	"github.com/upper/db/v4"
)

type Repository interface {
	FindAll() ([]Event, error)
	FindOne(id int64) (*Event, error)
}

type repository struct {
	collection db.Collection // add Collection field
}

func NewRepository(col db.Collection) Repository {
	return &repository{collection: col}
}

var (
	events []Event
	event  Event
)

func (r *repository) FindAll() ([]Event, error) {
	res := r.collection.Find()
	err := res.OrderBy("-date_time", "-name").All(&events)
	return events, err
}

func (r *repository) FindOne(id int64) (*Event, error) {
	res := r.collection.Find(id)
	err := res.One(&event)
	return &event, err
}
