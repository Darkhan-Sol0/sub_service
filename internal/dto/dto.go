package dto

import "time"

type (
	AddSubFromWeb struct {
		ServiceName string `json:"service_name" db:"service_name" example:"YandexGold"`
		Price       int    `json:"price" db:"price" example:"500"`
		UserId      string `json:"user_id" db:"user_id" example:"60601fee-2bf1-4721-ae6f-7636e79a0cba"`
		StartDate   string `json:"start_date" db:"start_date" example:"02-2022"`
		Month       int    `json:"month" db:"month" example:"5"`
	}

	AddSubToDb struct {
		ServiceName string    `json:"service_name" db:"service_name" example:"YandexGold"`
		Price       int       `json:"price" db:"price" example:"500"`
		UserId      string    `json:"user_id" db:"user_id" example:"60601fee-2bf1-4721-ae6f-7636e79a0cba"`
		StartDate   time.Time `json:"start_date" db:"start_date" example:"02-2022"`
		EndDate     time.Time `json:"end_date" db:"end_date" example:"03-2022"`
	}

	GetSubFromWeb struct {
		Id int `json:"id" db:"id" example:"1"`
	}

	GetSubFromDb struct {
		Id          int       `json:"id" db:"id" example:"1"`
		ServiceName string    `json:"service_name" db:"service_name" example:"YandexGold"`
		Price       int       `json:"price" db:"price" example:"500"`
		UserId      string    `json:"user_id" db:"user_id" example:"60601fee-2bf1-4721-ae6f-7636e79a0cba"`
		StartDate   time.Time `json:"start_date" db:"start_date" example:"02-2022"`
		EndDate     time.Time `json:"end_date" db:"end_date" example:"03-2022"`
	}

	GetSubByUserFromWeb struct {
		UserId string `json:"user_id" db:"user_id" example:"60601fee-2bf1-4721-ae6f-7636e79a0cba"`
	}

	GetSubPriceByFilterFromWeb struct {
		ServiceName string `json:"service_name" db:"service_name" example:"YandexGold"`
		UserId      string `json:"user_id" db:"user_id" example:"60601fee-2bf1-4721-ae6f-7636e79a0cba"`
		StartDate   string `json:"start_date" db:"start_date" example:"02-2022"`
		EndDate     string `json:"end_date" db:"end_date" example:"03-2022"`
	}

	GetSubPriceByFilterToDb struct {
		ServiceName string    `json:"service_name" db:"service_name" example:"YandexGold"`
		UserId      string    `json:"user_id" db:"user_id" example:"60601fee-2bf1-4721-ae6f-7636e79a0cba"`
		StartDate   time.Time `json:"start_date" db:"start_date" example:"02-2022"`
		EndDate     time.Time `json:"end_date" db:"end_date" example:"03-2022"`
	}

	GetSubPriceByFilterFromDb struct {
		Price int `json:"price" db:"price" example:"500"`
	}

	UpdateSubFromWeb struct {
		Id          int    `json:"id" db:"id" example:"1"`
		ServiceName string `json:"service_name" db:"service_name" example:"YandexGold"`
		Price       int    `json:"price" db:"price" example:"500"`
		UserId      string `json:"user_id" db:"user_id" example:"60601fee-2bf1-4721-ae6f-7636e79a0cba"`
		StartDate   string `json:"start_date" db:"start_date" example:"02-2022"`
		Month       int    `json:"month" db:"month" example:"5"`
	}

	UpdateSubToDb struct {
		Id          int       `json:"id" db:"id" example:"1"`
		ServiceName string    `json:"service_name" db:"service_name" example:"YandexGold"`
		Price       int       `json:"price" db:"price" example:"500"`
		UserId      string    `json:"user_id" db:"user_id" example:"60601fee-2bf1-4721-ae6f-7636e79a0cba"`
		StartDate   time.Time `json:"start_date" db:"start_date" example:"02-2022"`
		EndDate     time.Time `json:"end_date" db:"end_date" example:"03-2022"`
	}
)
