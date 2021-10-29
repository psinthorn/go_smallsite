package rates

import (
	"context"
	"time"

	"github.com/psinthorn/go_smallsite/datasources/drivers"
)

const (
	queryInsertRateType           = `insert into rate_types (title, acronym, description, status, created_at, updated_at) values ($1,$2,$3,$4,$5,$6) returning id`
	queryGetAllRateType           = `select * from rate_types order by id asc`
	queryGetAllRateTypeWithStatus = `select * from rate_types where status = $1 order by id asc`
	queryGetRateTypeById          = `select * from rate_types where id = $1`
	queryUpdateRateTypeById       = `update rate_types set title= $1, acronym = $2, description = $3, status = $4, updated_at = $5 where id = $6`

	queryDeleteRateTypeById = `delete from rate_types where id = $1`
)

type RateType rateType

var RateTypeService RateTypeInterface = &RateType{}

type RateTypeInterface interface {
	Create(RateType) (int, error)
	Get(string) ([]RateType, error)
	GetById(int) (RateType, error)
	Update(RateType) error
	Delete(int) error

	AdminGet() ([]RateType, error)
}

func (r *RateType) Create(rt RateType) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	dbConn, err := drivers.ConnectDB("pgx", drivers.PgDsn)
	if err != nil {
		panic(err)
	}

	var id int
	err = dbConn.SQL.QueryRowContext(ctx, queryInsertRateType, rt.Title, rt.Acronym, rt.Description, rt.Status, rt.CreatedAt, rt.UpdatedAt).Scan(&id)
	if err != nil {
		return 0, err
	}
	defer dbConn.SQL.Close()
	return id, nil
}

// Get get all data from database with specific status and return to request
func (r *RateType) Get(status string) ([]RateType, error) {
	var rts []RateType

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	dbConn, err := drivers.ConnectDB("pgx", drivers.PgDsn)
	if err != nil {
		panic(err)
	}

	rows, err := dbConn.SQL.QueryContext(ctx, queryGetAllRateTypeWithStatus, status)
	if err != nil {
		return rts, err
	}
	defer rows.Close()

	for rows.Next() {
		var rt RateType
		err := rows.Scan(
			&rt.Id,
			&rt.Title,
			&rt.Acronym,
			&rt.Description,
			&rt.Status,
			&rt.CreatedAt,
			&rt.UpdatedAt,
		)
		if err != nil {
			return rts, err
		}
		rts = append(rts, rt)
	}

	if err = rows.Err(); err != nil {
		return rts, err
	}

	return rts, nil
}

func (r *RateType) GetById(id int) (RateType, error) {
	var rt RateType

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	dbConn, err := drivers.ConnectDB("pgx", drivers.PgDsn)
	if err != nil {
		panic(err)
	}

	err = dbConn.SQL.QueryRowContext(ctx, queryGetRateTypeById, id).Scan(
		&rt.Id,
		&rt.Title,
		&rt.Acronym,
		&rt.Description,
		&rt.Status,
		&rt.CreatedAt,
		&rt.UpdatedAt,
	)
	if err != nil {
		return rt, err
	}
	defer dbConn.SQL.Close()
	return rt, nil
}

func (r *RateType) Update(rt RateType) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	dbConn, err := drivers.ConnectDB("pgx", drivers.PgDsn)
	if err != nil {
		panic(err)
	}

	_, err = dbConn.SQL.QueryContext(ctx, queryUpdateRateTypeById,
		rt.Title,
		rt.Acronym,
		rt.Description,
		rt.Status,
		time.Now(),
		rt.Id,
	)
	dbConn.SQL.Close()

	if err != nil {
		return err
	}

	return nil
}

// Delete remove data from database by id
func (r *RateType) Delete(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	dbConn, err := drivers.ConnectDB("pgx", drivers.PgDsn)
	if err != nil {
		panic(err)
	}

	_, err = dbConn.SQL.ExecContext(ctx, queryDeleteRateTypeById, id)
	if err != nil {
		return err
	}
	dbConn.SQL.Close()

	return nil
}

// AdminGet get all data all status from database and return to request
func (r *RateType) AdminGet() ([]RateType, error) {
	var rts []RateType

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	dbConn, err := drivers.ConnectDB("pgx", drivers.PgDsn)
	if err != nil {
		panic(err)
	}

	rows, err := dbConn.SQL.QueryContext(ctx, queryGetAllRateType)

	for rows.Next() {
		var rt RateType
		err := rows.Scan(
			&rt.Id,
			&rt.Title,
			&rt.Acronym,
			&rt.Description,
			&rt.Status,
			&rt.CreatedAt,
			&rt.UpdatedAt,
		)
		if err != nil {
			return rts, err
		}
		rts = append(rts, rt)
	}

	if err = rows.Err(); err != nil {
		return rts, err
	}

	return rts, nil
}
