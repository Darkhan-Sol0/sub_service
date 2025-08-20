package service

import (
	"fmt"
	"service/internal/datasource/repository"
	"service/internal/dto"
	"time"

	"github.com/labstack/echo/v4"
)

type (
	ServiceSubs struct {
		Storage repository.Storage
	}

	Service interface {
		AddNewSubs(ctx echo.Context, data dto.AddSubFromWeb) error
		GetSubById(ctx echo.Context, data dto.GetSubFromWeb) (dto.GetSubFromDb, error)
		GetListSub(ctx echo.Context) ([]dto.GetSubFromDb, error)
		GetListSubByUser(ctx echo.Context, data dto.GetSubByUserFromWeb) ([]dto.GetSubFromDb, error)
		GetPriceSubByFilter(ctx echo.Context, data dto.GetSubPriceByFilterFromWeb) (dto.GetSubPriceByFilterFromDb, error)
		UpdateSubById(ctx echo.Context, data dto.UpdateSubFromWeb) error
		DeleteSub(ctx echo.Context, data dto.GetSubFromWeb) error
	}
)

func NewService(storage repository.Storage) Service {
	return &ServiceSubs{
		Storage: storage,
	}
}

func (s *ServiceSubs) AddNewSubs(ctx echo.Context, data dto.AddSubFromWeb) error {
	sdate, err := time.Parse("01-2006", data.StartDate)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	month := data.Month
	if month <= 0 {
		month = 1
	}
	edate := sdate.AddDate(0, month, 0)
	dataOut := dto.AddSubToDb{
		ServiceName: data.ServiceName,
		Price:       data.Price,
		UserId:      data.UserId,
		StartDate:   sdate,
		EndDate:     edate,
	}
	if err := s.Storage.AddNewSubs(ctx, dataOut); err != nil {
		return fmt.Errorf("%w", err)
	}
	return nil
}

func (s *ServiceSubs) GetSubById(ctx echo.Context, data dto.GetSubFromWeb) (dto.GetSubFromDb, error) {
	dataOut, err := s.Storage.GetSubById(ctx, data)
	if err != nil {
		return dto.GetSubFromDb{}, fmt.Errorf("%w", err)
	}
	return dataOut, nil
}

func (s *ServiceSubs) GetListSub(ctx echo.Context) ([]dto.GetSubFromDb, error) {
	data, err := s.Storage.GetListSub(ctx)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (s *ServiceSubs) GetListSubByUser(ctx echo.Context, data dto.GetSubByUserFromWeb) ([]dto.GetSubFromDb, error) {
	dataOut, err := s.Storage.GetListSubByUser(ctx, data)
	if err != nil {
		return nil, err
	}
	return dataOut, nil
}

func (s *ServiceSubs) GetPriceSubByFilter(ctx echo.Context, dataIn dto.GetSubPriceByFilterFromWeb) (dto.GetSubPriceByFilterFromDb, error) {
	sdate, err := time.Parse("01-2006", dataIn.StartDate)
	if err != nil {
		return dto.GetSubPriceByFilterFromDb{}, fmt.Errorf("%w", err)
	}
	edate, err := time.Parse("01-2006", dataIn.EndDate)
	if err != nil {
		return dto.GetSubPriceByFilterFromDb{}, fmt.Errorf("%w", err)
	}
	data := dto.GetSubPriceByFilterToDb{
		ServiceName: dataIn.ServiceName,
		UserId:      dataIn.UserId,
		StartDate:   sdate,
		EndDate:     edate,
	}

	dataOut, err := s.Storage.GetPriceSubByFilter(ctx, data)
	if err != nil {
		return dto.GetSubPriceByFilterFromDb{}, err
	}
	return dataOut, nil
}

func (s *ServiceSubs) UpdateSubById(ctx echo.Context, data dto.UpdateSubFromWeb) error {
	sdate, edate := time.Time{}, time.Time{}
	if data.StartDate != "" {
		var err error
		sdate, err = time.Parse("01-2006", data.StartDate)
		if err != nil {
			return fmt.Errorf("%w", err)
		}
		month := data.Month
		if month <= 0 {
			month = 1
		}
		edate = sdate.AddDate(0, month, 0)
	}
	dataOut := dto.UpdateSubToDb{
		Id:          data.Id,
		ServiceName: data.ServiceName,
		Price:       data.Price,
		UserId:      data.UserId,
		StartDate:   sdate,
		EndDate:     edate,
	}
	if err := s.Storage.UpdateSubById(ctx, dataOut); err != nil {
		return fmt.Errorf("%w", err)
	}
	return nil
}

func (s *ServiceSubs) DeleteSub(ctx echo.Context, data dto.GetSubFromWeb) error {
	err := s.Storage.DeleteSub(ctx, data)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	return nil
}
