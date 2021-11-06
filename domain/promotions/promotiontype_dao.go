package promotions

import (
	"context"
	"time"

	"github.com/psinthorn/go_smallsite/datasources/drivers"
)

const (
	queryGetAllPromotionTypes      = `select id, title, description, start_date, end_date, status,created_at, updated_at from promotion_types where status = $1 order by id desc`
	queryAdminGetAllPromotionTypes = `select id, title, description, start_date, end_date, status,created_at, updated_at from promotion_types order by id desc`
	queryInsertPromotionType       = `INSERT INTO promotion_types (title, description, start_date, end_date, status, created_at, updated_at) values ($1,$2,$3,$4,$5, $6, $7) returning id`
	queryGetPromotionTypeById      = `SELECT pmt.id, pmt.title, pmt.description, pmt.start_date, pmt.end_date, pmt.status, pmt.created_at, pmt.updated_at from promotion_types pmt where pmt.id = $1`
	queryUpdatePromotionType       = `update promotion_types set title= $1, description = $2, start_date = $3, end_date = $4, status = $5, updated_at = $6 where id = $7`
	queryDeletePromotionType       = `delete from promotion_types where id = $1`
)

var PromotionTypeService promotionTypeInterface = &PromotionType{}

type PromotionType promotionType
type promotionTypeInterface interface {
	Create(PromotionType) (int, error)
	Get(string) ([]PromotionType, error)
	AdminGet() ([]PromotionType, error)
	GetById(int) (PromotionType, error)
	Update(PromotionType) error
	Delete(int) error
}

func (pm *PromotionType) Create(p PromotionType) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	dbConn, err := drivers.ConnectDB("pgx", drivers.PgDsn)
	if err != nil {
		panic(err)
	}

	var newPromotionTypeId int
	err = dbConn.SQL.QueryRowContext(ctx, queryInsertPromotionType, p.Title, p.Description, p.StartDate, p.EndDate, p.Status, p.CreatedAt, p.UpdatedAt).Scan(&newPromotionTypeId)
	if err != nil {
		return 0, nil
	}

	defer dbConn.SQL.Close()

	return newPromotionTypeId, nil
}

// Get return all current promotion types list
func (pm *PromotionType) Get(status string) ([]PromotionType, error) {
	var pts []PromotionType
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	dbConn, err := drivers.ConnectDB("pgx", drivers.PgDsn)
	if err != nil {
		return pts, err
	}

	rows, err := dbConn.SQL.QueryContext(ctx, queryGetAllPromotionTypes, status)
	if err != nil {
		return pts, err
	}
	defer rows.Close()

	for rows.Next() {
		var pt PromotionType
		err := rows.Scan(
			&pt.Id,
			&pt.Title,
			&pt.Description,
			&pt.StartDate,
			&pt.EndDate,
			&pt.Status,
			&pt.CreatedAt,
			&pt.UpdatedAt,
		)
		if err != nil {
			return pts, err
		}
		pts = append(pts, pt)
	}

	if err = rows.Err(); err != nil {
		return pts, err
	}

	return pts, nil
}

// AdminGet return all current promotion types list
func (pm *PromotionType) AdminGet() ([]PromotionType, error) {
	var pts []PromotionType
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	dbConn, err := drivers.ConnectDB("pgx", drivers.PgDsn)
	if err != nil {
		return pts, err
	}

	rows, err := dbConn.SQL.QueryContext(ctx, queryAdminGetAllPromotionTypes)
	if err != nil {
		return pts, err
	}
	defer rows.Close()

	for rows.Next() {
		var pt PromotionType
		err := rows.Scan(
			&pt.Id,
			&pt.Title,
			&pt.Description,
			&pt.StartDate,
			&pt.EndDate,
			&pt.Status,
			&pt.CreatedAt,
			&pt.UpdatedAt,
		)
		if err != nil {
			return pts, err
		}
		pts = append(pts, pt)
	}

	if err = rows.Err(); err != nil {
		return pts, err
	}

	return pts, nil
}

// GetByID select promotion type by id and return to request
func (s *PromotionType) GetById(id int) (PromotionType, error) {
	var pmt PromotionType
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	dbConn, err := drivers.ConnectDB("pgx", drivers.PgDsn)
	if err != nil {
		return pmt, err
	}

	err = dbConn.SQL.QueryRowContext(ctx, queryGetPromotionTypeById, id).Scan(
		&pmt.Id,
		&pmt.Title,
		&pmt.Description,
		&pmt.StartDate,
		&pmt.EndDate,
		&pmt.Status,
		&pmt.CreatedAt,
		&pmt.UpdatedAt,
	)
	if err != nil {
		return pmt, err
	}
	defer dbConn.SQL.Close()

	return pmt, nil

}

// Update edit and update data by id
func (s *PromotionType) Update(pmt PromotionType) error {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	dbConn, err := drivers.ConnectDB("pgx", drivers.PgDsn)
	if err != nil {
		return err
	}
	_, err = dbConn.SQL.QueryContext(ctx, queryUpdatePromotionType,
		pmt.Title,
		pmt.Description,
		pmt.StartDate,
		pmt.EndDate,
		pmt.Status,
		time.Now(),
		pmt.Id,
	)

	return nil
}

// Delete is delete data by id
func (s *PromotionType) Delete(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	dbConn, err := drivers.ConnectDB("pgx", drivers.PgDsn)
	if err != nil {
		return err
	}

	_, err = dbConn.SQL.ExecContext(ctx, queryDeletePromotionType, id)
	if err != nil {
		return err
	}

	return nil
}
