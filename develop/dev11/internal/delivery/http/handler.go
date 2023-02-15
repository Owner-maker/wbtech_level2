package http

import (
	"dev11/internal/service"
	"net/http"
)

type Handler struct {
	service service.User
}

func NewHandler(s service.User) *Handler {
	return &Handler{s}
}
func (h *Handler) InitRoutes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/create_event", h.createEvent)
	mux.HandleFunc("/update_event", h.updateEvent)
	mux.HandleFunc("/delete_event", h.deleteEvent)
	mux.HandleFunc("/events_for_day", h.getEventsForDay)
	mux.HandleFunc("/events_for_week", h.getEventsForWeek)
	mux.HandleFunc("/events_for_month", h.getEventsForMonth)
	handler := Log(mux)
	return handler
}
