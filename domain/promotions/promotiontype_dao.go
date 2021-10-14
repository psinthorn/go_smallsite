package domain

import (
	"context"
	"fmt"
	"time"

	"github.com/psinthorn/go_smallsite/datasources/drivers"
)

const (
	queryGetAllPromotionTypes = `select id, title, description, start_date, end_date, status,created_at, updated_at from promotion_types where status = $1 order by id asc`
	queryInsertPromotionType  = `INSERT INTO promotion_types (title, description, status, created_at, updated_at) values ($1,$2,$3,$4,$5) returning id`
	querygetPromotionTypeByID = `select id, title from promotion_types where id = $1`
)

var PromotionTypeService promotionTypeInterface = &PromotionType{}

type PromotionType promotionType
type promotionTypeInterface interface {
	Create(PromotionType) (int, error)
	Get(string) ([]PromotionType, error)
	// GetRoomTypeByID(id int) (PromotionType, error)
}

func (pm *PromotionType) Create(p PromotionType) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	dbConn, err := drivers.ConnectDB("pgx", drivers.PgDsn)
	if err != nil {
		panic(err)
	}

	var newPromotionTypeID int
	err = dbConn.SQL.QueryRowContext(ctx, queryInsertPromotionType, p.Title, p.Description, p.Status, p.CreatedAt, p.UpdatedAt).Scan(&newPromotionTypeID)
	if err != nil {
		return 0, nil
	}
	defer dbConn.SQL.Close()

	return newPromotionTypeID, nil
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
			&pt.ID,
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

	fmt.Println(pts)

	return pts, nil
}

// // GetRoomeTypeByID
// func (pm *PromotionType) GetRoomTypeByID(id int) (PromotionType, error) {

// 	var promotionType PromotionType
// 	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
// 	defer cancel()

// 	dbConn, err := drivers.ConnectDB("pgx", drivers.PgDsn)
// 	if err != nil {
// 		panic(err)
// 	}

// 	err = dbConn.SQL.QueryRowContext(ctx, querygetRoomtypeByID, id).Scan(&promotionType.ID, &promotionType.Title)
// 	if err != nil {
// 		return promotionType, err
// 	}
// 	defer dbConn.SQL.Close()

// 	return promotionType, nil
// }

// func (pm *PromotionType) Update() {}

// func (pm *PromotionType) Delete() {}
