package rates

import (
	"context"
	"time"

	"github.com/psinthorn/go_smallsite/datasources/drivers"
)

const (
	queryInsertPromotionRate = `insert into promotions_room_rate (title, room_type_id, rate_type_id, promotion_type_id, rate, start_date, end_date, status, created_at, updated_at) values ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10) returning id`

	queryGetAllPromotionRates = `select pmr.id, pmr.title, pmr.description, pmr.price, pmr.promotion_type_id, pmr.start_date, pmr.end_date, pmr.status, pmr.created_at, pmr.updated_at, pt.id, pt.title
							from promotions pmr
							left join promotion_types pt
							on (pmr.promotion_type_id = pt.id)
							where pmr.status = $1
							order by pmr.id desc`

	queryAdminGetAllPromotionRates = `select pmr.id, pmr.title, pmr.promotion_id, pmr.room_type_id, pmr.rate_type_id, pmr.rate, pmr.start_date, pmr.end_date, pmr.status, pmr.created_at, pmr.updated_at
							from promotions_room_rate pmr
							order by pmr.id desc`

	queryGetPromotionRateById = `SELECT pm.id, pm.title, pm.description, pm.price, pm.promotion_type_id, pm.start_date, pm.end_date, pm.status, pm.created_at, pm.updated_at, pt.id, pt.title
							from promotions pm
							left join promotion_types pt
							on (pm.promotion_type_id = pt.id)
							where pm.id = $1`

	queryUpdatePromotionRateById = `update promotions set title= $1, description = $2, promotion_type_id = $3, start_date = $4, end_date = $5, price = $6, status = $7, updated_at = $8 where id = $9`

	queryDeletePromotionRateById = `delete from promotions_room_rate where id = $1`
)

var PromotionRateService promotionRateInterface = &PromotionRate{}

type PromotionRate promotionRate
type promotionRateInterface interface {
	Create(PromotionRate) (int, error)
	Get(string) ([]PromotionRate, error)
	GetById(int) (PromotionRate, error)
	Update(PromotionRate) error
	Delete(int) error

	AdminGet() ([]PromotionRate, error)
}

// Create insert and return room data
// เพิ่มข้อมมูลห้องพักเก็บในดาต้าเบสและคืนข้อมูลที่เพิ่มสำเร็จแล้วกลับให้ผู้ใช้งาน
func (s *PromotionRate) Create(p PromotionRate) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	dbConn, err := drivers.ConnectDB("pgx", drivers.PgDsn)
	if err != nil {
		panic(err)
	}

	var newProId int
	err = dbConn.SQL.QueryRowContext(ctx, queryInsertPromotionRate, p.Title, p.RoomTypeId, p.RateTypeId, p.PromotionId, p.Rate, p.StartDate, p.EndDate, p.Status, p.CreatedAt, p.UpdatedAt).Scan(&newProId)
	if err != nil {
		return 0, err
	}
	defer dbConn.SQL.Close()
	return newProId, err
}

// Get select all rooms  data from table and return all rooms slice to request
func (s *PromotionRate) Get(st string) ([]PromotionRate, error) {
	var promotionRates []PromotionRate
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	dbConn, err := drivers.ConnectDB("pgx", drivers.PgDsn)
	if err != nil {
		return promotionRates, err
	}

	rows, err := dbConn.SQL.QueryContext(ctx, queryGetAllPromotionRates, st)
	if err != nil {
		return promotionRates, err
	}
	defer rows.Close()

	for rows.Next() {
		var p PromotionRate
		err := rows.Scan(
			&p.Id,
			&p.Title,
			&p.RateTypeId,
			&p.RoomTypeId,
			&p.PromotionId,
			&p.StartDate,
			&p.EndDate,
			&p.Rate,
			&p.Status,
			&p.CreatedAt,
			&p.UpdatedAt,
		)

		if err != nil {
			return promotionRates, err
		}

		promotionRates = append(promotionRates, p)
	}

	if err = rows.Err(); err != nil {
		return promotionRates, err
	}

	return promotionRates, nil

}

// Get select all rooms  data from table and return all rooms slice to request
func (s *PromotionRate) AdminGet() ([]PromotionRate, error) {
	var promotionRates []PromotionRate
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	dbConn, err := drivers.ConnectDB("pgx", drivers.PgDsn)
	if err != nil {
		return promotionRates, err
	}

	rows, err := dbConn.SQL.QueryContext(ctx, queryAdminGetAllPromotionRates)
	if err != nil {
		return promotionRates, err
	}
	defer rows.Close()

	for rows.Next() {
		var p PromotionRate
		err := rows.Scan(
			&p.Id,
			&p.Title,
			&p.RateTypeId,
			&p.RoomTypeId,
			&p.PromotionId,
			&p.StartDate,
			&p.EndDate,
			&p.Status,
			&p.Rate,
			&p.CreatedAt,
			&p.UpdatedAt,
			// &p.PromotionType.Id,
			// &p.PromotionType.Title,
		)

		if err != nil {
			return promotionRates, err
		}

		promotionRates = append(promotionRates, p)
	}

	if err = rows.Err(); err != nil {
		return promotionRates, err
	}

	return promotionRates, nil

}

// GetRoomByID select room by id and return to request
func (s *PromotionRate) GetById(id int) (PromotionRate, error) {
	var pr PromotionRate
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	dbConn, err := drivers.ConnectDB("pgx", drivers.PgDsn)
	if err != nil {
		return pr, err
	}

	err = dbConn.SQL.QueryRowContext(ctx, queryGetPromotionRateById, id).Scan(
		&pr.Id,
		&pr.Title,

		&pr.Rate,
		pr.RoomTypeId,
		pr.RateTypeId,
		pr.PromotionId,
		&pr.StartDate,
		&pr.EndDate,
		&pr.Status,
		&pr.CreatedAt,
		&pr.UpdatedAt,
		// &pr.PromotionType.Id,
		// &pr.PromotionType.Title,
	)
	if err != nil {
		return pr, err
	}
	defer dbConn.SQL.Close()

	return pr, nil

}

// Update update room data
func (s *PromotionRate) Update(pr PromotionRate) error {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	dbConn, err := drivers.ConnectDB("pgx", drivers.PgDsn)
	if err != nil {
		return err
	}
	_, err = dbConn.SQL.QueryContext(ctx, queryUpdatePromotionRateById,
		pr.Title,
		pr.RoomTypeId,
		pr.RateTypeId,
		pr.PromotionId,
		pr.StartDate,
		pr.EndDate,
		pr.Rate,
		pr.Status,
		time.Now(),
		pr.Id,
	)

	return nil
}

// Delete is delete room by id
func (s *PromotionRate) Delete(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	dbConn, err := drivers.ConnectDB("pgx", drivers.PgDsn)
	if err != nil {
		return err
	}

	_, err = dbConn.SQL.ExecContext(ctx, queryDeletePromotionRateById, id)
	if err != nil {
		return err
	}

	return nil
}
