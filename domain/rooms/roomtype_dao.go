package rooms

import (
	"context"
	"time"

	"github.com/psinthorn/go_smallsite/datasources/drivers"
)

const (
	queryInsertRoomType = `INSERT INTO room_types (title, description, status, created_at, updated_at) values ($1,$2,$3,$4,$5) returning id`
)

var RoomTypeService roomTypeInterface = &RoomType{}

type RoomType roomType
type roomTypeInterface interface {
	Create(RoomType) (int, error)
}

func (rs *RoomType) Create(r RoomType) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	dbConn, err := drivers.ConnectDB("pgx", drivers.PgDsn)
	if err != nil {
		panic(err)
	}

	var newRoomTypeID int
	err = dbConn.SQL.QueryRowContext(ctx, queryInsertRoomType, r.Title, r.Description, r.Status, r.CreatedAt, r.UpdatedAt).Scan(&newRoomTypeID)
	if err != nil {
		return 0, nil
	}
	defer dbConn.SQL.Close()

	return newRoomTypeID, nil
}

func (rs *RoomType) GetAll() {}
func (rs *RoomType) Get()    {}
func (rs *RoomType) Update() {}
func (rs *RoomType) Delete() {}
