package event

import (
	"fmt"

	"github.com/upper/db/v4"
)

type Repository interface {
	FindAll() ([]Event, error)
	FindOne(id int64) (*Event, error)
}

const EventsCount int64 = 10

type repository struct {
	collection db.Collection // add Collection field
}

func NewRepository(col db.Collection) Repository {
	return &repository{collection: col}
}

func (r *repository) FindAll() ([]Event, error) {
	var events []Event
	eventsCol := r.collection

	res := eventsCol.Find()
	err := res.OrderBy("-date_time", "-name").All(&events)
	return events, err
}

func (r *repository) FindOne(id int64) (*Event, error) {
	if id <= EventsCount {
		return &Event{
			Id:   id,
			Name: fmt.Sprintf("Event #%d", id),
		}, nil
	} else {
		return nil, nil
	}
}
