package buyer

import "github.com/ArdiSasongko/ticketing_app/model/domain"

type EventRepository interface {
	ListEvents(filters map[string]string, sort string, limit int, page int) ([]domain.Events, error)
}