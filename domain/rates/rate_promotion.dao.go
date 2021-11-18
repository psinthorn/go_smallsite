package rates

import (
	"context"
	"fmt"
	"time"

	"github.com/psinthorn/go_smallsite/datasources/drivers"
)

const (
	queryInsertPromotionRate = `insert into promotions_room_rate (title, image, room_type_id, promotion_type_id, rate, start_date, end_date, status, created_at, updated_at) values ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11) returning id`

	queryGetAllPromotionRates = `select pmr.id, pmr.title, pmr.image, pmr.room_type_id, pmr.promotion_id, pmr.start_date, pmr.end_date, pmr.rate, pmr.status, pmr.created_at, pmr.updated_at, rt.id, rt.title, rt.description, pm.title
							from promotions_room_rate pmr
							left join room_types rt 
							on (pmr.room_type_id = rt.id) 
							left join promotions pm 
							on (pmr.promotion_id = pm.id) 
							where pmr.status = $1
							order by pmr.rate asc`

	queryAdminGetAllPromotionRates = `select pmr.id, pmr.title, pmr.image, pmr.room_type_id, pmr.promotion_id, pmr.start_date, pmr.end_date, pmr.rate, pmr.status, pmr.created_at, pmr.updated_at
							from promotions_room_rate pmr
							order by pmr.id desc`

	queryAdminGetAllPromotionRatesById = `select pmr.id, pmr.title, pmr.image, pmr.room_type_id, pmr.promotion_id, pmr.rate, pmr.start_date, pmr.end_date, pmr.status, pmr.created_at, pmr.updated_at
							from promotions_room_rate pmr
							where promotion_id = $1
							order by pmr.id desc`

	queryGetPromotionRateById = `SELECT pm.id, pm.title, pmr.image, room_type_id, pm.promotion_id, pm.start_date, pm.end_date, pm.rate, pm.status, pm.created_at, pm.updated_at from promotions_room_rate pm where pm.id = $1`

	queryGetAllPromotionRatesByPromotionId = `select pmr.id, pmr.title, pmr.image, pmr.room_type_id, pmr.promotion_id, pmr.start_date, pmr.end_date, pmr.rate, pmr.status, pmr.created_at, pmr.updated_at, rt.id, rt.title, rt.description, pm.title
							from promotions_room_rate pmr
							left join room_types rt 
							on (pmr.room_type_id = rt.id) 
							left join promotions pm 
							on (pmr.promotion_id = pm.id) 
							where pmr.status = $1 && pmr.promotion_id = $2
							order by pmr.rate asc`

	queryUpdatePromotionRateById = `update promotions_room_rate set rate = $1, status = $2, updated_at = $3 where id = $4`

	queryDeletePromotionRateById = `delete from promotions_room_rate where id = $1`
)

var PromotionRateService promotionRateInterface = &PromotionRate{}

type promotionRateInterface interface {
	Create(PromotionRate) (int, error)
	Get(string) ([]PromotionRate, error)
	GetById(int) (PromotionRate, error)
	Update(PromotionRate) error
	Delete(int) error

	AdminGet() ([]PromotionRate, error)
	GetRatesByPromotionId(int) ([]PromotionRate, error)
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
	err = dbConn.SQL.QueryRowContext(ctx, queryInsertPromotionRate, p.Title, p.RoomTypeId, p.PromotionId, p.Rate, p.StartDate, p.EndDate, p.Status, p.CreatedAt, p.UpdatedAt).Scan(&newProId)
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
			&p.Image,
			&p.RoomTypeId,
			&p.PromotionId,
			&p.StartDate,
			&p.EndDate,
			&p.Rate,
			&p.Status,
			&p.CreatedAt,
			&p.UpdatedAt,
			&p.RoomType.ID,
			&p.RoomType.Title,
			&p.RoomType.Description,
			&p.PromotionTitle,
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
			&p.Image,
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

// GetPromotionByID select room by id and return to request
func (s *PromotionRate) GetById(id int) (PromotionRate, error) {
	var pr PromotionRate
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	dbConn, err := drivers.ConnectDB("pgx", drivers.PgDsn)
	if err != nil {
		return pr, err
	}

	fmt.Println(id)

	err = dbConn.SQL.QueryRowContext(ctx, queryGetPromotionRateById, id).Scan(
		&pr.Id,
		&pr.Title,
		&pr.Image,
		&pr.RoomTypeId,
		&pr.PromotionId,
		&pr.StartDate,
		&pr.EndDate,
		&pr.Rate,
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
		// pr.Title,
		// pr.RoomTypeId,
		// pr.PromotionId,
		// pr.StartDate,
		// pr.EndDate,
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

// GetRatesByPromotionId return all rates that belongs to promotion id
func (s *PromotionRate) GetRatesByPromotionId(id int) ([]PromotionRate, error) {
	var pmrs []PromotionRate
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	dbConn, err := drivers.ConnectDB("pgx", drivers.PgDsn)
	if err != nil {
		return pmrs, err
	}

	rows, err := dbConn.SQL.QueryContext(ctx, queryAdminGetAllPromotionRatesById, id)
	if err != nil {
		return pmrs, err
	}
	defer rows.Close()

	for rows.Next() {
		var p PromotionRate
		err := rows.Scan(
			&p.Id,
			&p.Title,
			&p.Image,
			&p.RoomTypeId,
			&p.PromotionId,
			&p.Rate,
			&p.StartDate,
			&p.EndDate,
			&p.Status,
			&p.CreatedAt,
			&p.UpdatedAt,
			// &p.PromotionType.Id,
			// &p.PromotionType.Title,
		)

		if err != nil {
			return pmrs, err
		}

		pmrs = append(pmrs, p)
	}

	if err = rows.Err(); err != nil {
		return pmrs, err
	}

	return pmrs, nil
}
