package rooms

import (
	"context"
	"fmt"
	"time"

	"github.com/psinthorn/go_smallsite/datasources/drivers"
)

const (
	queryInsertRoomType         = `INSERT INTO room_types (title, description, status, created_at, updated_at) values ($1,$2,$3,$4,$5) returning id`
	queryGetRoomTypeByStatus    = `select id, title  from room_types  where status = $1`
	queryGetAllRoomTypeByStatus = `select id, title  from room_types`
	querygetRoomtypeByID        = `select id, title from room_types where id = $1`
	queryDeleteRoomTypeById     = `delete from room_types where id = $1`
)

var RoomTypeService roomTypeInterface = &RoomType{}

type RoomType roomType
type roomTypeInterface interface {
	Create(RoomType) (int, error)
	Get(string) ([]RoomType, error)
	GetAll() ([]RoomType, error)
	GetByID(int) (RoomType, error)
	Update(RoomType) error
	Delete(int) error
}

// Get
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

// Get
func (rs *RoomType) Get(status string) ([]RoomType, error) {
	var rts []RoomType
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	dbConn, err := drivers.ConnectDB("pgx", drivers.PgDsn)
	if err != nil {
		return rts, err
	}
	rows, err := dbConn.SQL.QueryContext(ctx, queryGetRoomTypeByStatus, status)
	for rows.Next() {
		var rt RoomType
		err := rows.Scan(
			&rt.ID,
			&rt.Title,
		)

		if err != nil {
			return rts, err
		}

		rts = append(rts, rt)

	}
	if err = rows.Err(); err != nil {
		return rts, err
	}
	defer rows.Close()
	fmt.Println(rts)
	return rts, nil
}

// GetAll
func (rs *RoomType) GetAll() ([]RoomType, error) {
	var rts []RoomType
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	dbConn, err := drivers.ConnectDB("pgx", drivers.PgDsn)
	if err != nil {
		return rts, err
	}
	rows, err := dbConn.SQL.QueryContext(ctx, queryGetAllRoomTypeByStatus)
	for rows.Next() {
		var rt RoomType
		err := rows.Scan(
			&rt.ID,
			&rt.Title,
		)

		if err != nil {
			return rts, err
		}

		rts = append(rts, rt)

	}
	if err = rows.Err(); err != nil {
		return rts, err
	}
	defer rows.Close()
	fmt.Println(rts)
	return rts, nil
}

// GetRoomeTypeByID
func (rs *RoomType) GetByID(id int) (RoomType, error) {
	var roomType RoomType
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	dbConn, err := drivers.ConnectDB("pgx", drivers.PgDsn)
	if err != nil {
		panic(err)
	}

	err = dbConn.SQL.QueryRowContext(ctx, querygetRoomtypeByID, id).Scan(&roomType.ID, &roomType.Title)
	if err != nil {
		return roomType, err
	}
	defer dbConn.SQL.Close()

	return roomType, nil
}

func (rs *RoomType) Update(r RoomType) error {
	return nil
}

// Delete
func (rs *RoomType) Delete(id int) error {
	ctx, cancle := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancle()

	dbConn, err := drivers.ConnectDB("pgx", drivers.PgDsn)
	if err != nil {
		return err
	}
	result, err := dbConn.SQL.ExecContext(ctx, queryDeleteRoomTypeById, id)
	if err != nil {
		return nil
	}
	fmt.Printf("Deleted result is %s ", result)
	defer dbConn.SQL.Close()

	return nil
}
