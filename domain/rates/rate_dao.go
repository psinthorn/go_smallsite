package rates

import (
	"context"
	"fmt"
	"time"

	"github.com/psinthorn/go_smallsite/datasources/drivers"
)

const (
	queryInsertRate           = `insert into rates (title, description, start_date, end_date, status, created_at, updated_at) values ($1,$2,$3,$4,$5,$6) returning id`
	queryGetAllRate           = `select * from rates order by id asc`
	queryGetAllRateWithStatus = `select * from rates where status = $1 order by id asc`
	queryGetRateById          = `select * from rates where id = $1`
	queryUpdateRateById       = `update rates set title= $1, acronym = $2, description = $3, status = $4, updated_at = $5 where id = $6`

	queryDeleteRateById = `delete from rates where id = $1`

	querySelectRoomTypeIdAndTitle = `select id, title frome room_type`
)

var RoomRateService roomRateInterface = &RoomRate{}

type roomRateInterface interface {
	Create(RoomRate) (int, error)
	Get(string) ([]RoomRate, error)
	GetById(int) (RoomRate, error)
	Update(RoomRate) error
	Delete(int) error

	AdminGet() ([]RoomRate, error)
}

func (r *RoomRate) Create(rr RoomRate) (int, error) {
	// rooms types

	// promotion

	// rate type

	return 0, nil
}

func (r *RoomRate) Get(status string) ([]RoomRate, error) {
	var rr []RoomRate
	fmt.Println("Please implement me")
	return rr, nil
}

func (r *RoomRate) GetById(id int) (RoomRate, error) {
	var rr RoomRate
	fmt.Println("Please implement me")
	return rr, nil
}

func (r *RoomRate) Update(RoomRate) error {
	fmt.Println("Please implement me")
	return nil
}

func (r *RoomRate) Delete(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	dbConn, err := drivers.ConnectDB("pgx", drivers.PgDsn)
	if err != nil {
		panic(err)
	}

	_, err = dbConn.SQL.ExecContext(ctx, queryDeleteRateById, id)
	if err != nil {
		return err
	}
	dbConn.SQL.Close()

	return nil
}

func (r *RoomRate) AdminGet() ([]RoomRate, error) {
	var rr []RoomRate

	return rr, nil
}
