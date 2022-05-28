package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/test_server/internal/domain/event"
)

type EventController struct {
	service *event.Service
}

func NewEventController(s *event.Service) *EventController {
	return &EventController{
		service: s,
	}
}

func (c *EventController) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var eventData map[string]string
		err := json.NewDecoder(r.Body).Decode(&eventData)
		if err != nil {
			fmt.Println(err)
		}

		err = (*c.service).Create(eventData)
		if err != nil {
			fmt.Printf("EventController.Create(): %s", err)
			err = internalServerError(w, err)
			if err != nil {
				fmt.Printf("EventController.Create(): %s", err)
			}
			return
		}

		err = success(w, 201)
		if err != nil {
			fmt.Printf("EventController.Create(): %s", err)
		}
	}
}

func (c *EventController) FindAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		events, err := (*c.service).FindAll()
		if err != nil {
			fmt.Printf("EventController.FindAll(): %s", err)
			err = internalServerError(w, err)
			if err != nil {
				fmt.Printf("EventController.FindAll(): %s", err)
			}
			return
		}

		err = success(w, events)
		if err != nil {
			fmt.Printf("EventController.FindAll(): %s", err)
		}
	}
}

func (c *EventController) FindOne() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
		if err != nil {
			fmt.Printf("EventController.FindOne(): %s", err)
			err = internalServerError(w, err)
			if err != nil {
				fmt.Printf("EventController.FindOne(): %s", err)
			}
			return
		}
		event, err := (*c.service).FindOne(id)
		if err != nil {
			fmt.Printf("EventController.FindOne(): %s", err)
			err = internalServerError(w, err)
			if err != nil {
				fmt.Printf("EventController.FindOne(): %s", err)
			}
			return
		}

		err = success(w, event)
		if err != nil {
			fmt.Printf("EventController.FindOne(): %s", err)
		}
	}
}

func (c *EventController) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
		if err != nil {
			fmt.Printf("EventController.FindOne(): %s", err)
			err = internalServerError(w, err)
			if err != nil {
				fmt.Printf("EventController.FindOne(): %s", err)
			}
			return
		}

		var eventData map[string]string
		err = json.NewDecoder(r.Body).Decode(&eventData)
		if err != nil {
			fmt.Println(err)
		}

		err = (*c.service).Update(id, eventData)
		if err != nil {
			fmt.Printf("EventController.Create(): %s", err)
			err = internalServerError(w, err)
			if err != nil {
				fmt.Printf("EventController.Create(): %s", err)
			}
			return
		}

		err = success(w, 204)
		if err != nil {
			fmt.Printf("EventController.Create(): %s", err)
		}
	}
}
