package event

import (
	"errors"
	"strings"

	"github.com/upper/db/v4"
)

type Repository interface {
	Create(eventData map[string]string) error
	FindAll() ([]Event, error)
	FindOne(id int64) (*Event, error)
	Update(id int64, eventData map[string]string) error
}

type repository struct {
	eventCollection db.Collection // add Collection field
}

func NewRepository(col db.Collection) Repository {
	return &repository{eventCollection: col}
}

var (
	events []Event
	event  Event
)

func (r *repository) Create(eventData map[string]string) error {
	event.setValueFromMap(eventData)

	eventExists, _ := r.eventCollection.Find().And(
		"name = ? AND latitude = ? AND Longitude = ? AND DATE(date_time) = ?",
		event.Name, event.Latitude, event.Longitude, strings.Split(event.DateTime, " ")[0],
	).Exists()
	if eventExists {
		err := errors.New("An event with such name, coords and date already exists")
		return err
	}

	_, err := r.eventCollection.Insert(event)
	return err
}

func (r *repository) FindAll() ([]Event, error) {
	res := r.eventCollection.Find()
	err := res.OrderBy("-date_time", "-name").All(&events)
	return events, err
}

func (r *repository) FindOne(id int64) (*Event, error) {
	res := r.eventCollection.Find(id)
	err := res.One(&event)
	return &event, err
}

func (r *repository) Update(id int64, eventData map[string]string) error {
	event.setValueFromMap(eventData)

	eventExists, _ := r.eventCollection.Find().And(
		"id != ? AND name = ? AND latitude = ? AND Longitude = ? AND DATE(date_time) = ?",
		id, event.Name, event.Latitude, event.Longitude, strings.Split(event.DateTime, " ")[0],
	).Exists()
	if eventExists {
		err := errors.New("An event with a different id, but with the same name, coords and date, already exists")
		return err
	}

	res := r.eventCollection.Find(id)
	err := res.Update(event)
	return err
}
