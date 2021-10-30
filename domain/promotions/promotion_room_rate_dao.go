package promotions

import (
	"context"
	"time"

	"github.com/psinthorn/go_smallsite/datasources/drivers"
)

const (
	queryGetAllPromotionRoomRate = `select * from promotions_room_rate where status = $1`
	insertPromotionerate         = `insert into room_rates (title, promotion_id, room_type_id, rate_type_id, rate, status, start_date, end_date) value ($1,$2,$3,$4,$5,$6,$7,$8,$9) returning id`
)

var PromotionRoomRateService promotionRoomRateInterface = &PromotionRoomRate{}

type PromotionRoomRate promotionRoomRate
type promotionRoomRateInterface interface {
	Create(PromotionRoomRate) (int, error)
	Get(string) ([]PromotionRoomRate, error)
	GetById(int) (PromotionRoomRate, error)
}

// Create insert new room rate to database
func (p *PromotionRoomRate) Create(pmr PromotionRoomRate) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	dbConn, err := drivers.ConnectDB("pgx", drivers.PgDsn)
	if err != nil {
		return 0, err
	}

	var newId int
	err = dbConn.SQL.QueryRowContext(ctx, queryInsertPromotion, pmr.Title, pmr.PromotionId, pmr.RoomTypeId, pmr.Rate, pmr.Status, pmr.CreatedAt, pmr.UpdatedAt).Scan(&newId)
	if err != nil {
		return 0, err
	}
	defer dbConn.SQL.Close()

	return newId, err
}

// Get select and return all data from promotion room rate
func (p *PromotionRoomRate) Get(status string) ([]PromotionRoomRate, error) {
	var pmrs []PromotionRoomRate
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	dbConn, err := drivers.ConnectDB("pgx", drivers.PgDsn)
	if err != nil {
		panic(err)
	}

	rows, err := dbConn.SQL.QueryContext(ctx, queryGetAllPromotionRoomRate, status)
	defer rows.Close()

	for rows.Next() {
		var pmr PromotionRoomRate
		err := rows.Scan(
			&pmr.Id,
			&pmr.Title,
			&pmr.RoomTypeId,
			&pmr.PromotionId,
			&pmr.Rate,
			&pmr.Status,
			&pmr.CreatedAt,
			&pmr.UpdatedAt,
		)
		if err != nil {
			return pmrs, err
		}
		pmrs = append(pmrs, pmr)
	}

	if err = rows.Err(); err != nil {
		return pmrs, err
	}

	return pmrs, nil
}

// GetRoomByID return room details
func (s *PromotionRoomRate) GetById(id int) (PromotionRoomRate, error) {
	var pmr PromotionRoomRate
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	dbConn, err := drivers.ConnectDB("pgx", drivers.PgDsn)
	if err != nil {
		return pmr, err
	}

	err = dbConn.SQL.QueryRowContext(ctx, queryGetPromotionById, id).Scan(
		&pmr.Id,
		&pmr.Title,
		&pmr.RoomTypeId,
		&pmr.PromotionId,
		&pmr.Rate,
		&pmr.Status,
		&pmr.CreatedAt,
		&pmr.UpdatedAt,
	)
	if err != nil {
		return pmr, err
	}
	defer dbConn.SQL.Close()

	return pmr, nil

}

// Update
func (s *PromotionRoomRate) Update(pmr promotionRoomRate) error {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	dbConn, err := drivers.ConnectDB("pgx", drivers.PgDsn)
	if err != nil {
		return err
	}

	_, err = dbConn.SQL.QueryContext(ctx, queryUpdateById,
		pmr.Title,
		pmr.RoomTypeId,
		pmr.PromotionId,
		pmr.Rate,
		pmr.Status,
		pmr.CreatedAt,
		pmr.UpdatedAt,
		time.Now(),
	)

	return nil
}
