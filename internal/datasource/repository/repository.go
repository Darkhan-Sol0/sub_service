package repository

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"service/internal/dto"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/labstack/echo/v4"
)

type (
	Client interface {
		Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
		QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
		Exec(ctx context.Context, sql string, args ...interface{}) (pgconn.CommandTag, error)
		Close()
	}

	Repository struct {
		Client Client
	}

	Storage interface {
		AddNewSubs(ctx echo.Context, data dto.AddSubToDb) error
		GetSubById(ctx echo.Context, data dto.GetSubFromWeb) (dto.GetSubFromDb, error)
		GetListSub(ctx echo.Context) ([]dto.GetSubFromDb, error)
		GetListSubByUser(ctx echo.Context, data dto.GetSubByUserFromWeb) ([]dto.GetSubFromDb, error)
		GetPriceSubByFilter(ctx echo.Context, data dto.GetSubPriceByFilterToDb) (dto.GetSubPriceByFilterFromDb, error)
		UpdateSubById(ctx echo.Context, data dto.UpdateSubToDb) error
		DeleteSub(ctx echo.Context, data dto.GetSubFromWeb) error
	}
)

func NewDatabase(client Client) Storage {
	return &Repository{
		Client: client,
	}
}

func StructToNamedArgs(s any) (pgx.NamedArgs, error) {
	args := pgx.NamedArgs{}
	v := reflect.ValueOf(s)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	t := v.Type()
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		key := field.Tag.Get("db")
		if key == "" {
			key = strings.ToLower(field.Name)
		}
		args[key] = v.Field(i).Interface()
	}
	return args, nil
}

func (r *Repository) AddNewSubs(ctx echo.Context, data dto.AddSubToDb) error {
	query := `INSERT INTO subs (
	service_name,
	price,
	user_id,
	start_date,
	end_date) VALUES (
	@service_name,
	@price,
	@user_id,
	@start_date,
	@end_date)`

	args, err := StructToNamedArgs(data)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	_, err = r.Client.Exec(ctx.Request().Context(), query, args)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	return nil
}

// func (r *Repository) AddNewSubs(ctx echo.Context, data dto.AddSubToDb) error {
// 	query := `INSERT INTO subs (
// 	service_name,
// 	price,
// 	user_id,
// 	start_date,
// 	end_date) VALUES ($1, $2, $3, $4, $5)`
// 	_, err := r.Client.Exec(ctx.Request().Context(), query,
// 		data.ServiceName,
// 		data.Price,
// 		data.UserId,
// 		data.StartDate,
// 		data.EndDate,
// 	)
// 	if err != nil {
// 		return fmt.Errorf("%w", err)
// 	}
// 	return nil
// }

func (r *Repository) GetSubById(ctx echo.Context, data dto.GetSubFromWeb) (dto.GetSubFromDb, error) {
	if data.Id <= 0 {
		return dto.GetSubFromDb{}, fmt.Errorf("invalid sub ID: %d", data.Id)
	}
	query := `SELECT
	id,
	service_name,
	price,
	user_id,
	start_date,
	end_date 
	FROM subs
	WHERE id = $1`
	row := r.Client.QueryRow(ctx.Request().Context(), query, data.Id)
	var out dto.GetSubFromDb
	if err := row.Scan(
		&out.Id,
		&out.ServiceName,
		&out.Price,
		&out.UserId,
		&out.StartDate,
		&out.EndDate); err != nil {
		return dto.GetSubFromDb{}, fmt.Errorf("%w", err)
	}
	return out, nil
}

func (r *Repository) GetListSub(ctx echo.Context) ([]dto.GetSubFromDb, error) {
	query := `SELECT
	id,
	service_name,
	price,
	user_id,
	start_date,
	end_date 
	FROM subs`
	rows, err := r.Client.Query(ctx.Request().Context(), query)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	var out []dto.GetSubFromDb
	for rows.Next() {
		var data dto.GetSubFromDb
		if err := rows.Scan(
			&data.Id,
			&data.ServiceName,
			&data.Price,
			&data.UserId,
			&data.StartDate,
			&data.EndDate); err != nil {
			return nil, fmt.Errorf("%w", err)
		}
		out = append(out, data)
	}
	return out, nil
}

func (r *Repository) GetListSubByUser(ctx echo.Context, data dto.GetSubByUserFromWeb) ([]dto.GetSubFromDb, error) {
	if data.UserId == "" {
		return nil, fmt.Errorf("invalid user ID: %s", data.UserId)
	}
	query := `SELECT
	id,
	service_name,
	price,
	user_id,
	start_date,
	end_date 
	FROM subs
	WHERE user_id = $1`
	rows, err := r.Client.Query(ctx.Request().Context(), query, data.UserId)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	var out []dto.GetSubFromDb
	for rows.Next() {
		var data dto.GetSubFromDb
		if err := rows.Scan(
			&data.Id,
			&data.ServiceName,
			&data.Price,
			&data.UserId,
			&data.StartDate,
			&data.EndDate); err != nil {
			return nil, fmt.Errorf("%w", err)
		}
		out = append(out, data)
	}
	return out, nil
}

func (r *Repository) GetPriceSubByFilter(ctx echo.Context, data dto.GetSubPriceByFilterToDb) (dto.GetSubPriceByFilterFromDb, error) {
	query := `SELECT
	SUM(price)
	FROM subs
	WHERE service_name = $1 AND user_id = $2 AND start_date BETWEEN $3 AND $4`
	row := r.Client.QueryRow(ctx.Request().Context(), query, data.ServiceName, data.UserId, data.StartDate, data.EndDate)
	var out dto.GetSubPriceByFilterFromDb
	if err := row.Scan(&out.Price); err != nil {
		return dto.GetSubPriceByFilterFromDb{}, fmt.Errorf("%w", err)
	}
	return out, nil
}

func (r *Repository) UpdateSubById(ctx echo.Context, data dto.UpdateSubToDb) error {
	if data.Id <= 0 {
		return fmt.Errorf("invalid sub ID: %d", data.Id)
	}
	var setClauses []string
	var args []interface{}
	argID := 1
	if data.ServiceName != "" {
		setClauses = append(setClauses, fmt.Sprintf("service_name = $%d", argID))
		args = append(args, data.ServiceName)
		argID++
	}
	if data.Price != 0 {
		setClauses = append(setClauses, fmt.Sprintf("price = $%d", argID))
		args = append(args, data.Price)
		argID++
	}
	if data.UserId != "" {
		setClauses = append(setClauses, fmt.Sprintf("user_id = $%d", argID))
		args = append(args, data.UserId)
		argID++
	}
	if !data.StartDate.IsZero() {
		setClauses = append(setClauses, fmt.Sprintf("start_date = $%d", argID))
		args = append(args, data.StartDate)
		argID++
	}
	if !data.EndDate.IsZero() {
		setClauses = append(setClauses, fmt.Sprintf("end_date = $%d", argID))
		args = append(args, data.EndDate)
		argID++
	}
	if len(setClauses) == 0 {
		return fmt.Errorf("no fields to update for sub with id %d", data.Id)
	}
	query := fmt.Sprintf("UPDATE subs SET %s WHERE id = $%d RETURNING id", strings.Join(setClauses, ", "), argID)
	args = append(args, data.Id)
	var updatedID int
	err := r.Client.QueryRow(ctx.Request().Context(), query, args...).Scan(&updatedID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return fmt.Errorf("sub with id %d not found", data.Id)
		}
		return fmt.Errorf("failed to update sub: %w", err)
	}
	return nil
}

func (r *Repository) DeleteSub(ctx echo.Context, data dto.GetSubFromWeb) error {
	if data.Id <= 0 {
		return fmt.Errorf("invalid sub ID: %d", data.Id)
	}
	query := `DELETE FROM subs WHERE id = $1`
	res, err := r.Client.Exec(ctx.Request().Context(), query, data.Id)
	if err != nil {
		return fmt.Errorf("failed to delete sub: %w", err)
	}
	if res.RowsAffected() == 0 {
		return fmt.Errorf("sub with id %d not found", data.Id)
	}
	return nil
}
