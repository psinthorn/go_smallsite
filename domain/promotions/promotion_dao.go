package promotions

import (
	"context"
	"time"

	"github.com/psinthorn/go_smallsite/datasources/drivers"
	"github.com/psinthorn/go_smallsite/domain/rates"
	"github.com/psinthorn/go_smallsite/domain/rooms"
)

const (
	queryInsertPromotion = `insert into promotions (title, description, price, start_date, end_date, promotion_type_id, status, created_at, updated_at) values ($1,$2,$3,$4,$5,$6,$7,$8,$9) returning id`

	queryInsertPromotionRate = `insert into promotions_room_rate (title, room_type_id, promotion_id, start_date, end_date, rate, status, created_at, updated_at) values ($1,$2,$3,$4,$5,$6,$7,$8,$9) returning id`

	queryGetAllPromotions = `select pms.id, pms.title, pms.description, pms.price, pms.promotion_type_id, pms.start_date, pms.end_date, pms.status, pms.created_at, pms.updated_at, pt.id, pt.title
							from promotions pms
							left join promotion_types pt 
							on (pms.promotion_type_id = pt.id) 
							where pms.status = $1   
							order by pms.id asc`

	queryAdminGetAllPromotions = `select pms.id, pms.title, pms.description, pms.price, pms.promotion_type_id, pms.start_date, pms.end_date, pms.status, pms.created_at, pms.updated_at, pt.id, pt.title
							from promotions pms
							left join promotion_types pt 
							on (pms.promotion_type_id = pt.id)   
							order by pms.id desc`

	queryGetPromotionById = `SELECT pm.id, pm.title, pm.description, pm.price, pm.promotion_type_id, pm.start_date, pm.end_date, pm.status, pm.created_at, pm.updated_at, pt.id, pt.title
							from promotions pm 
							left join promotion_types pt 
							on (pm.promotion_type_id = pt.id) 
							where pm.id = $1`

	queryUpdateById = `update promotions set title= $1, description = $2, promotion_type_id = $3, start_date = $4, end_date = $5, price = $6, status = $7, updated_at = $8 where id = $9`

	queryDeletePromotionById = `delete from promotions where id = $1`
)

var PromotionService promotionDomainInterface = &Promotion{}

type Promotion promotion
type promotionDomainInterface interface {
	Create(Promotion) (int, error)
	Get(string) ([]Promotion, error)
	GetById(int) (Promotion, error)
	CreatePromotionRate(int) (int, error)
	Update(Promotion) error
	Delete(int) error

	AdminGet() ([]Promotion, error)
}

// Create insert and return room data
// เพิ่มข้อมมูลห้องพักเก็บในดาต้าเบสและคืนข้อมูลที่เพิ่มสำเร็จแล้วกลับให้ผู้ใช้งาน
func (s *Promotion) Create(p Promotion) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	dbConn, err := drivers.ConnectDB("pgx", drivers.PgDsn)
	if err != nil {
		panic(err)
	}

	var newProId int
	err = dbConn.SQL.QueryRowContext(ctx, queryInsertPromotion, p.Title, p.Description, p.Price, p.StartDate, p.EndDate, p.PromotionTypeId, p.Status, p.CreatedAt, p.UpdatedAt).Scan(&newProId)
	if err != nil {
		return 0, err
	}

	// Auto generate by room type
	// get promotion type id from returnin id
	// get all room type
	// loop and generate promotion
	var newProRateId int
	roomTypes, err := rooms.RoomTypeService.GetAll()
	for _, x := range roomTypes {
		var pr rates.PromotionRate
		pr.Title = x.Title
		pr.RoomTypeId = x.ID
		pr.PromotionId = newProId
		pr.Rate = 0
		pr.Status = p.Status
		pr.StartDate = p.StartDate
		pr.EndDate = p.EndDate
		pr.CreatedAt = time.Now()
		pr.UpdatedAt = time.Now()

		err = dbConn.SQL.QueryRowContext(ctx, queryInsertPromotionRate, pr.Title, pr.RoomTypeId, pr.PromotionId, pr.StartDate, pr.EndDate, pr.Rate, pr.Status, pr.CreatedAt, pr.UpdatedAt).Scan(&newProRateId)
		if err != nil {
			return 0, err
		}
	}

	defer dbConn.SQL.Close()
	return newProId, err
}

// CreatePromotionRate
// เพิ่มข้อมมูลห้องพักเก็บในดาต้าเบสและคืนข้อมูลที่เพิ่มสำเร็จแล้วกลับให้ผู้ใช้งาน
func (s *Promotion) CreatePromotionRate(id int) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	dbConn, err := drivers.ConnectDB("pgx", drivers.PgDsn)
	if err != nil {
		return 0, err
	}

	// get promotion by id
	pm, err := s.GetById(id)

	// Get promotion rate by promotion id
	pmr, err := rates.PromotionRateService.GetRatesByPromotionId(id)
	if err != nil {
		return 0, err
	}

	var newProRateId int
	if len(pmr) <= 0 {
		// Auto generate by room type
		// get promotion type id from returnin id
		// get all room type
		// loop and generate promotion
		roomTypes, err := rooms.RoomTypeService.GetAll()
		for _, x := range roomTypes {
			var pr rates.PromotionRate
			pr.Title = x.Title
			pr.RoomTypeId = x.ID
			pr.PromotionId = id
			pr.Rate = 0
			pr.Status = pm.Status
			pr.StartDate = pm.StartDate
			pr.EndDate = pm.EndDate
			pr.CreatedAt = time.Now()
			pr.UpdatedAt = time.Now()

			err = dbConn.SQL.QueryRowContext(ctx, queryInsertPromotionRate, pr.Title, pr.RoomTypeId, pr.PromotionId, pr.StartDate, pr.EndDate, pr.Rate, pr.Status, pr.CreatedAt, pr.UpdatedAt).Scan(&newProRateId)
			if err != nil {
				return 0, err
			}
		}
	}
	defer dbConn.SQL.Close()

	return newProRateId, err
}

// Get select all rooms  data from table and return all rooms slice to request
func (s *Promotion) Get(st string) ([]Promotion, error) {
	var promotions []Promotion
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	dbConn, err := drivers.ConnectDB("pgx", drivers.PgDsn)
	if err != nil {
		return promotions, err
	}

	rows, err := dbConn.SQL.QueryContext(ctx, queryGetAllPromotions, st)
	if err != nil {
		return promotions, err
	}
	defer rows.Close()

	for rows.Next() {
		var p Promotion
		err := rows.Scan(
			&p.Id,
			&p.Title,
			&p.Description,
			&p.Price,
			&p.PromotionTypeId,
			&p.StartDate,
			&p.EndDate,
			&p.Status,
			&p.CreatedAt,
			&p.UpdatedAt,
			&p.PromotionType.Id,
			&p.PromotionType.Title,
		)

		if err != nil {
			return promotions, err
		}

		promotions = append(promotions, p)
	}

	if err = rows.Err(); err != nil {
		return promotions, err
	}

	return promotions, nil

}

// Get select all rooms  data from table and return all rooms slice to request
func (s *Promotion) AdminGet() ([]Promotion, error) {
	var promotions []Promotion
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	dbConn, err := drivers.ConnectDB("pgx", drivers.PgDsn)
	if err != nil {
		return promotions, err
	}

	rows, err := dbConn.SQL.QueryContext(ctx, queryAdminGetAllPromotions)
	if err != nil {
		return promotions, err
	}
	defer rows.Close()

	for rows.Next() {
		var p Promotion
		err := rows.Scan(
			&p.Id,
			&p.Title,
			&p.Description,
			&p.Price,
			&p.PromotionTypeId,
			&p.StartDate,
			&p.EndDate,
			&p.Status,
			&p.CreatedAt,
			&p.UpdatedAt,
			&p.PromotionType.Id,
			&p.PromotionType.Title,
		)

		if err != nil {
			return promotions, err
		}

		promotions = append(promotions, p)
	}

	if err = rows.Err(); err != nil {
		return promotions, err
	}

	return promotions, nil

}

// GetByID select room by id and return to request
func (s *Promotion) GetById(id int) (Promotion, error) {
	var pm Promotion
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	dbConn, err := drivers.ConnectDB("pgx", drivers.PgDsn)
	if err != nil {
		return pm, err
	}

	err = dbConn.SQL.QueryRowContext(ctx, queryGetPromotionById, id).Scan(
		&pm.Id,
		&pm.Title,
		&pm.Description,
		&pm.Price,
		&pm.PromotionTypeId,
		&pm.StartDate,
		&pm.EndDate,
		&pm.Status,
		&pm.CreatedAt,
		&pm.UpdatedAt,
		&pm.PromotionType.Id,
		&pm.PromotionType.Title,
	)
	if err != nil {
		return pm, err
	}
	defer dbConn.SQL.Close()

	return pm, nil

}

// Update update room data
func (s *Promotion) Update(pm Promotion) error {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	dbConn, err := drivers.ConnectDB("pgx", drivers.PgDsn)
	if err != nil {
		return err
	}
	_, err = dbConn.SQL.QueryContext(ctx, queryUpdateById,
		pm.Title,
		pm.Description,
		pm.PromotionTypeId,
		pm.StartDate,
		pm.EndDate,
		pm.Price,
		pm.Status,
		time.Now(),
		pm.Id,
	)

	return nil
}

// Delete is delete room by id
func (s *Promotion) Delete(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	dbConn, err := drivers.ConnectDB("pgx", drivers.PgDsn)
	if err != nil {
		return err
	}

	_, err = dbConn.SQL.ExecContext(ctx, queryDeletePromotionById, id)
	if err != nil {
		return err
	}

	return nil
}
