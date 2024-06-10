package buyer_entity

import (
	"time"

	"github.com/ArdiSasongko/ticketing_app/model/domain"
)

type EventEntity struct {
	ID        int       `json:"id"`
	SellerID  int       `json:"seller_id"`
	Name      string    `json:"name"`
	Date      time.Time `json:"date"`
	Location  string    `json:"location"`
	Qty       int       `json:"qty"`
	Category  string    `json:"category"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func ToEventEntity(event domain.Event) EventEntity {
	return EventEntity{
		ID:        event.ID,
		SellerID:  event.SellerID,
		Name:      event.Name,
		Date:      event.Date,
		Location:  event.Location,
		Qty:       event.Qty,
		Category:  event.Category,
		Price:     event.Price,
		CreatedAt: event.CreatedAt,
		UpdatedAt: event.UpdatedAt,
	}
}

func ToEventListEntity(events []domain.Event) []EventEntity {
	eventList := []EventEntity{}

	for _, event := range events {
		eventList = append(eventList, ToEventEntity(event))
	}

	return eventList
}
