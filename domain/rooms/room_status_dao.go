package domain

import (
	"context"
	"time"

	"github.com/psinthorn/go_smallsite/datasources/drivers"
)

const (
	queryInsertRoomStatus = "INSERT INTO room_status (title, description, status, created_at, updated_at) values ($1,$2,$3,$4,$5) returning id"
)

var RoomStatusService roomStatusInterface = &RoomStatus{}

type RoomStatus roomStatus
type roomStatusInterface interface {
	Create(RoomStatus) (int, error)
	GetAll() error
	Get() error
	Update() error
	Delete() error
}

func (rs *RoomStatus) Create(st RoomStatus) (int, error) {
	// Create contect with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// connect database
	dbConn, err := drivers.ConnectDB("pgx", drivers.PgDsn)
	if err != nil {
		return 0, err
	}

	// Run SQL statment with queryRowContext
	var newStatusId int
	err = dbConn.SQL.QueryRowContext(ctx, queryInsertRoomStatus, st.Title, st.Description, st.Status, st.CreatedAt, st.UpdatedAt).Scan(&newStatusId)
	if err != nil {
		return 0, err
	}

	defer dbConn.SQL.Close()
	return newStatusId, nil
}

// GetAll
func (rs *RoomStatus) GetAll() error {
	return nil
}

// Get
func (rs *RoomStatus) Get() error {
	return nil
}

// Update
func (rs *RoomStatus) Update() error {
	return nil
}

// Delete
func (rs *RoomStatus) Delete() error {
	return nil
}
