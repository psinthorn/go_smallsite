package domain

import (
	"context"
	"fmt"
	"time"

	"github.com/psinthorn/go_smallsite/datasources/drivers"
)

const (
	queryInsertRoom  = "insert into rooms (roomtype_id, room_name, room_no, description, status, created_at, updated_at) values ($1,$2,$3,$4,$5,$6,$7) returning id"
	queryGetRoomByID = `SELECT id, roomtype_id, room_name, room_no, description, status, created_at, updated_at FROM rooms WHERE id = $1`
)

var RoomService roomDomainInterface = &Room{}

type Room room
type roomDomainInterface interface {
	Create(Room) (int, error)
	GetRoomByID(int) (Room, error)
}

// Create insert and return room data
// เพิ่มข้อมมูลห้องพักเก็บในดาต้าเบสและคืนข้อมูลที่เพิ่มสำเร็จแล้วกลับให้ผู้ใช้งาน
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

func (s *Room) GetRoomByID(id int) (Room, error) {
	fmt.Println("room id for getroomby id is: ", id)
	var roombyId Room
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	dbConn, err := drivers.ConnectDB("pgx", drivers.PgDsn)
	if err != nil {
		return roombyId, err
	}

	err = dbConn.SQL.QueryRowContext(ctx, queryGetRoomByID, id).Scan(
		&roombyId.ID,
		&roombyId.RoomTypeId,
		&roombyId.RoomName,
		&roombyId.RoomNo,
		&roombyId.Description,
		&roombyId.Status,
		&roombyId.CreatedAt,
		&roombyId.UpdatedAt,
	)
	if err != nil {
		return roombyId, err
	}

	fmt.Println("print room from dao: ", roombyId)
	defer dbConn.SQL.Close()

	return roombyId, nil

}
