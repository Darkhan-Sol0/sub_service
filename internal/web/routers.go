package web

import (
	"service/internal/service"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type (
	routing struct {
		service service.Service
		log     *logrus.Logger
	}

	Routing interface {
		RegisterRoutes(e *echo.Echo)
	}
)

func NewRouting(service service.Service, log *logrus.Logger) Routing {
	return &routing{
		service: service,
		log:     log,
	}
}

func (r *routing) RegisterRoutes(e *echo.Echo) {
	e.Use(r.logger)

	e.GET("/", r.Hello)
	e.POST("/add_sub", r.AddSub)
	e.GET("/get_sub_by_id/:id", r.GetSubById)
	e.GET("/get_list", r.GetListSub)
	e.GET("/get_list_by_user/:uuid", r.GetListSubByUser)
	e.GET("/get_price_subs", r.GetPriceSubByFilter)
	e.PATCH("/update_sub", r.UpdateSub)
	e.DELETE("/delete_sub/:id", r.Delete)

	e.GET("/swagger/*", echoSwagger.WrapHandler)
}

func (r *routing) logger(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		ctx.Set("logger", r.log)
		return next(ctx)
	}
}
