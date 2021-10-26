package contents

import (
	"context"
	"time"

	"github.com/psinthorn/go_smallsite/datasources/drivers"
)

// import (
// 	mysql_db "github.com/psinthorn/gostack_users-api/datasources/mysql/users_db"
// 	"github.com/psinthorn/gostack_users-api/domains/errors"
// 	date_utils "github.com/psinthorn/gostack_users-api/utils/date"
// 	mysql_utils "github.com/psinthorn/gostack_users-api/utils/mysql"
// )

const (
	queryInsertContent     = "INSERT INTO contents(title, sub_title, content, content_type, category, image, tags, author, status, date_created) VALUES(?,?,?,?,?,?,?,?,?,?);"
	queryGetContentById    = "SELECT * FROM contents WHERE id = ?"
	queryDeleteContentById = "DELETE FROM contents where id = ?"
	queryGetAllContents    = "SELECT * FROM contents ORDER BY id DESC"
)

var (
	ContentService contentDomainInterface = &Content{}
)

type Content content
type contentDomainInterface interface {
	Create(Content) (int, error)
	GetAll() ([]Content, error)
}

var (
	contentDB = make(map[int64]*Content)
)

// ------------------------------------
// Create new content
func (c *Content) Create(ct Content) (int, error) {
	// Create new variable
	// create context with time out and defer cancle
	// create database connection and close database connection
	// return result

	var newContentId int
	ctx, cancle := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancle()

	dbConn, err := drivers.ConnectDB("pgx", drivers.PgDsn)
	if err != nil {
		return 0, err
	}

	err = dbConn.SQL.QueryRowContext(ctx, queryInsertContent,
		ct.Id,
		ct.Title,
		ct.SubTitle,
		ct.Content,
		ct.ContentType,
		ct.Section,
		ct.Category,
		ct.Image,
		ct.Tags,
		ct.Author,
		ct.Status,
		ct.CreatedAt,
		ct.UpdatedAt).Scan(&newContentId)

	if err != nil {
		return 0, err
	}
	defer dbConn.SQL.Close()

	return newContentId, nil
}

// ------------------------------------
// GetAll contents
func (c *Content) GetAll() ([]Content, error) {
	// prepare contents array
	var Contents []Content

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	dbConn, err := drivers.ConnectDB("pgx", drivers.PgDsn)
	if err != nil {
		return Contents, err
	}
	results, err := dbConn.SQL.QueryContext(ctx, queryGetAllContents)
	if err != nil {
		return Contents, err
	}
	defer results.Close()

	// Read each row of data from resutls and append content to contents array
	for results.Next() {
		var ct Content
		err := results.Scan(
			&ct.Id,
			&ct.Title,
			&ct.SubTitle,
			&ct.Content,
			&ct.ContentType,
			&ct.Category,
			&ct.Image,
			&ct.Tags,
			&ct.Author,
			&ct.Status,
			&ct.CreatedAt,
			&ct.UpdatedAt,
		)

		if err != nil {
			return Contents, err
		}
		Contents = append(Contents, ct)
	}

	if len(Contents) == 0 {
		return nil, err
	}
	return Contents, nil

}
