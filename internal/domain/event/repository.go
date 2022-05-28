package event

import (
	"errors"
	"strconv"
	"strings"

	"github.com/upper/db/v4"
)

type Repository interface {
	Create(map[string]string) error
	FindAll() ([]Event, error)
	FindOne(id int64) (*Event, error)
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
	event.Name = eventData["name"]
	event.Image = eventData["image"]
	event.Description = eventData["description"]
	event.Gallery = strings.Split(eventData["gallery"], ", ")
	event.Latitude, _ = strconv.ParseFloat(eventData["latitude"], 64)
	event.Longitude, _ = strconv.ParseFloat(eventData["longitude"], 64)
	event.DateTime = eventData["date_time"]

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
