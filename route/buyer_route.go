package route

import (
	"github.com/ArdiSasongko/ticketing_app/app"
	buyer_controller "github.com/ArdiSasongko/ticketing_app/controller/buyer"
	helper "github.com/ArdiSasongko/ticketing_app/helper"
	"github.com/ArdiSasongko/ticketing_app/middleware"
	buyer_query_builder "github.com/ArdiSasongko/ticketing_app/query_builder/buyer"
	buyer_repository "github.com/ArdiSasongko/ticketing_app/repository/buyer"
	buyer_service "github.com/ArdiSasongko/ticketing_app/service/buyer"
	"github.com/labstack/echo/v4"
)

func RegisterBuyerRoutes(prefix string, e *echo.Echo) {
	db := app.DBConnection()
	token := helper.NewTokenUseCase()
	buyerAuthRepo := buyer_repository.NewBuyerRepository(db)
	buyerAuthService := buyer_service.NewBuyerService(buyerAuthRepo, token)
	buyerAuthController := buyer_controller.NewBuyerController(buyerAuthService)
	buyerEventQB := buyer_query_builder.NewEventQueryBuilder(db)
	buyerEventRepo := buyer_repository.NewEventRepository(buyerEventQB)
	buyerEventService := buyer_service.NewEventService(buyerEventRepo)
	buyerEventController := buyer_controller.NewEventController(buyerEventService)

	g := e.Group(prefix)

	authRoute := g.Group("/auth")
	authRoute.POST("/register", buyerAuthController.Register)
	authRoute.POST("/login", buyerAuthController.Login)

	meRoute := g.Group("/me", middleware.JWTProtection())
	meRoute.POST("/update", buyerAuthController.Update)

	eventRoute := g.Group("/events")
	eventRoute.GET("/", buyerEventController.GetEventList)
}
