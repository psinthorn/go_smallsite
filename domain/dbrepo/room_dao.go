package dbrepo

import (
	"context"
	"time"

	"github.com/psinthorn/go_smallsite/datasources/drivers"
)

const (
	queryInsertRoom = "insert into rooms (roomtype_id, room_name, room_no, description, status, created_at, updated_at) values ($1,$2,$3,$4,$5,$6,$7) returning id"
)

var RoomService roomDomainInterface = &Room{}

type Room room
type roomDomainInterface interface {
	Create(Room) (int, error)
}

func (s *Room) Create(room Room) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	dbConn, err := drivers.ConnectDB("pgx", drivers.PgDsn)
	if err != nil {
		panic(err)
	}

	var newRoomId int
	err = dbConn.SQL.QueryRowContext(ctx, queryInsertRoom, room.RoomTypeId, room.RoomName, room.RoomNo, room.Description, room.Status, room.CreatedAt, room.UpdatedAt).Scan(&newRoomId)
	if err != nil {
		return 0, err
	}
	defer dbConn.SQL.Close()

	return newRoomId, err
}
