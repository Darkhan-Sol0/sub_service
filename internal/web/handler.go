package web

import (
	"net/http"
	"service/internal/dto"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type Response struct {
	Data any `json:"data"`
}

// @Summary Test endpoint
// @Description Returns Hello World message
// @Tags Test
// @Accept  json
// @Produce  json
// @Success 200 {object} Response "Success response"
// @Router / [get]
func (r *routing) Hello(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, Response{Data: "Hello, World!"})
}

// @Summary Add subscription
// @Description Add a new subscription
// @Tags Subscriptions
// @Accept  json
// @Produce  json
// @Param   request body dto.AddSubFromWeb true "Subscription data"
// @Success 200 {object} Response "Success response"
// @Failure 400 {object} Response "Bad request"
// @Router /add_sub [post]
func (r *routing) AddSub(ctx echo.Context) error {
	logger := ctx.Get("logger").(*logrus.Logger)
	var data dto.AddSubFromWeb
	if err := ctx.Bind(&data); err != nil {
		logger.Info("Not OK")
		return ctx.JSON(http.StatusBadRequest, Response{Data: err.Error()})
	}
	logger.Info(data)
	if err := r.service.AddNewSubs(ctx, data); err != nil {
		logger.Info("add:Not OK ", data)
		return ctx.JSON(http.StatusBadRequest, Response{Data: err.Error()})
	}
	logger.Info("add:OK ", data)
	return ctx.JSON(http.StatusOK, Response{Data: "OK"})
}

// @Summary Get subscription by ID
// @Description Get subscription details by ID
// @Tags Subscriptions
// @Accept  json
// @Produce  json
// @Param   id path int true "Subscription ID"
// @Success 200 {object} Response "Success response"
// @Failure 400 {object} Response "Bad request"
// @Router /get_sub_by_id/{id} [get]с
func (r *routing) GetSubById(ctx echo.Context) (err error) {
	logger := ctx.Get("logger").(*logrus.Logger)
	var data dto.GetSubFromWeb
	if data.Id, err = strconv.Atoi(ctx.Param("id")); err != nil {
		logger.Info("Not OK")
		return ctx.JSON(http.StatusBadRequest, Response{Data: err.Error()})
	}
	dataOut, err := r.service.GetSubById(ctx, data)
	if err != nil {
		logger.Info("Not OK")
		return ctx.JSON(http.StatusBadRequest, Response{Data: err.Error()})
	}
	logger.Info("OK")
	return ctx.JSON(http.StatusOK, Response{Data: dataOut})
}

// @Summary Get user subscriptions
// @Description Get list of subscriptions by user
// @Tags Subscriptions
// @Accept  json
// @Produce  json
// @Param   uuid path string true "User UUID"
// @Success 200 {object} Response "Success response"
// @Failure 400 {object} Response "Bad request"
// @Router /get_list_by_user/{uuid} [get]
func (r *routing) GetListSubByUser(ctx echo.Context) (err error) {
	logger := ctx.Get("logger").(*logrus.Logger)
	var data dto.GetSubByUserFromWeb
	data.UserId = ctx.Param("uuid")
	if data.UserId == "" {
		logger.Info("Not OK")
		return ctx.JSON(http.StatusBadRequest, Response{Data: "incorrect uuid"})
	}
	dataOut, err := r.service.GetListSubByUser(ctx, data)
	if err != nil {
		logger.Info("Not OK")
		return ctx.JSON(http.StatusBadRequest, Response{Data: err.Error()})
	}
	logger.Info("OK")
	return ctx.JSON(http.StatusOK, Response{Data: dataOut})
}

// @Summary Get all subscriptions
// @Description Get list of all subscriptions
// @Tags Subscriptions
// @Accept  json
// @Produce  json
// @Success 200 {object} Response "Success response"
// @Failure 400 {object} Response "Bad request"
// @Router /get_list [get]
func (r *routing) GetListSub(ctx echo.Context) (err error) {
	logger := ctx.Get("logger").(*logrus.Logger)
	dataOut, err := r.service.GetListSub(ctx)
	if err != nil {
		logger.Info("Not OK")
		return ctx.JSON(http.StatusBadRequest, Response{Data: err.Error()})
	}
	logger.Info("OK")
	return ctx.JSON(http.StatusOK, Response{Data: dataOut})
}

// @Summary Get subscription price by filter
// @Description Get subscription price based on filter criteria
// @Tags Subscriptions
// @Accept  json
// @Produce  json
// @Param   serv query string true "Service name"
// @Param   uuid query string true "User UUID"
// @Param   sdate query string true "Start date (MM-YYYY)"
// @Param   edate query string true "End date (MM-YYYY)"
// @Success 200 {object} Response "Success response"
// @Failure 400 {object} Response "Bad request"
// @Router /get_price_subs [get]
func (r *routing) GetPriceSubByFilter(ctx echo.Context) (err error) {
	logger := ctx.Get("logger").(*logrus.Logger)
	var data dto.GetSubPriceByFilterFromWeb
	data.ServiceName = ctx.QueryParam("serv")
	data.UserId = ctx.QueryParam("uuid")
	data.StartDate = ctx.QueryParam("sdate")
	data.EndDate = ctx.QueryParam("edate")
	if data.ServiceName == "" || data.UserId == "" || data.StartDate == "" || data.EndDate == "" {
		logger.Info("Not OK")
		return ctx.JSON(http.StatusBadRequest, Response{Data: "All parameters (serv, uuid, sdate, edate) are required"})
	}
	dataOut, err := r.service.GetPriceSubByFilter(ctx, data)
	if err != nil {
		logger.Info("Not OK")
		return ctx.JSON(http.StatusBadRequest, Response{Data: err.Error()})
	}
	logger.Info(dataOut)
	return ctx.JSON(http.StatusOK, Response{Data: dataOut})
}

// @Summary Update subscription
// @Description Update existing subscription
// @Tags Subscriptions
// @Accept  json
// @Produce  json
// @Param   request body dto.UpdateSubFromWeb true "Subscription data to update"
// @Success 200 {object} Response "Success response"
// @Failure 400 {object} Response "Bad request"
// @Router /update_sub [patch]
func (r *routing) UpdateSub(ctx echo.Context) (err error) {
	logger := ctx.Get("logger").(*logrus.Logger)
	var data dto.UpdateSubFromWeb
	if err := ctx.Bind(&data); err != nil {
		logger.Info("Not OK")
		return ctx.JSON(http.StatusBadRequest, Response{Data: err.Error()})
	}
	if err := r.service.UpdateSubById(ctx, data); err != nil {
		logger.Info("Not OK")
		return ctx.JSON(http.StatusBadRequest, Response{Data: err.Error()})
	}
	logger.Info("OK")
	return ctx.JSON(http.StatusOK, Response{Data: "OK"})
}

// @Summary Delete subscription
// @Description Delete subscription by ID
// @Tags Subscriptions
// @Accept  json
// @Produce  json
// @Param   id path int true "Subscription ID"
// @Success 200 {object} Response "Success response"
// @Failure 400 {object} Response "Bad request"
// @Router /delete_sub/{id} [delete]
func (r *routing) Delete(ctx echo.Context) (err error) {
	logger := ctx.Get("logger").(*logrus.Logger)
	var data dto.GetSubFromWeb
	if data.Id, err = strconv.Atoi(ctx.Param("id")); err != nil {
		logger.Info("Not OK")
		return ctx.JSON(http.StatusBadRequest, Response{Data: err.Error()})
	}
	if err := r.service.DeleteSub(ctx, data); err != nil {
		logger.Info("Not OK")
		return ctx.JSON(http.StatusBadRequest, Response{Data: err.Error()})
	}
	logger.Info("OK")
	return ctx.JSON(http.StatusOK, Response{Data: "OK"})
}
