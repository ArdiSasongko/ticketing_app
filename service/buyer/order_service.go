package buyer_service

import (
	"github.com/ArdiSasongko/ticketing_app/model/entity/buyer"
	"github.com/ArdiSasongko/ticketing_app/model/web/buyer"
	"gorm.io/gorm"
)

type OrderService interface {
	WithTx(tx *gorm.DB) OrderService

	ListOrder(filters map[string]string, sort string, limit int, page int) ([]buyer_entity.HistoryLiteEntity, error)
	ViewOrder(historyId int) (buyer_entity.HistoryEntity, error)
	CreateOrder(request buyer_web.CreateOrderRequest, buyerId int) (buyer_entity.HistoryEntity, error)
	PayOrder(orderId int) (buyer_entity.HistoryEntity, error)
	DeleteActiveOrder(buyerId int) error
	DeleteOrder(historyId int) error

	generateOrderNumber(buyerId int) (string, error)
}
