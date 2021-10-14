package domain

import (
	"context"
	"fmt"
	"time"

	"github.com/psinthorn/go_smallsite/datasources/drivers"
)

const (
	queryInsertPromotion  = "insert into rooms (title, description, price, start_date, end_date, status, created_at, updated_at) values ($1,$2,$3,$4,$5,$6,$7,$8,$9) returning id"
	queryGetAllPromotions = `select id, title, description, price, promotion_type_id, start_date, end_date, status, created_at, updated_at from promotions where status = $1  order by id asc`
	queryGetPromotionByID = `SELECT id, roomtype_id, room_name, room_no, description, status, created_at, updated_at FROM rooms WHERE id = $1`
)

var PromotionService promotionDomainInterface = &Promotion{}

type Promotion promotion
type promotionDomainInterface interface {
	Create(Promotion) (int, error)
	Get(string) ([]Promotion, error)
	// GetPromotionByID(int) (Promotion, error)
	// Update(Promotion) (Promotion, error)
	// Delete(int) error
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
	defer dbConn.SQL.Close()
	return newProId, err
}

// Get Return all rooms slice
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
			&p.ID,
			&p.Title,
			&p.Description,
			&p.Price,
			&p.PromotionTypeId,
			&p.StartDate,
			&p.EndDate,
			&p.Status,
			&p.CreatedAt,
			&p.UpdatedAt,
		)

		if err != nil {
			return promotions, err
		}

		promotions = append(promotions, p)
	}

	if err = rows.Err(); err != nil {
		return promotions, err
	}

	fmt.Println(promotions)

	return promotions, nil

}

// // GetRoomByID return room details
// func (s *Room) GetRoomByID(id int) (Room, error) {

// 	var roombyId Room
// 	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
// 	defer cancel()

// 	dbConn, err := drivers.ConnectDB("pgx", drivers.PgDsn)
// 	if err != nil {
// 		return roombyId, err
// 	}

// 	err = dbConn.SQL.QueryRowContext(ctx, queryGetRoomByID, id).Scan(
// 		&roombyId.ID,
// 		&roombyId.RoomTypeId,
// 		&roombyId.RoomName,
// 		&roombyId.RoomNo,
// 		&roombyId.Description,
// 		&roombyId.Status,
// 		&roombyId.CreatedAt,
// 		&roombyId.UpdatedAt,
// 	)
// 	if err != nil {
// 		return roombyId, err
// 	}
// 	defer dbConn.SQL.Close()

// 	return roombyId, nil

// }

// func (s *Room) Update(r Room) (Room, error) {
// 	var room Room

// 	return room, nil
// }

// func (s *Room) Delete(id int) error {

// 	return nil
// }
