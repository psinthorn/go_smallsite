package contents

import "time"

type category struct {
	ID          string
	Title       string
	Desc        string
	Section     string
	Parent      string
	AccessLevel int
	CreatedAt   time.Time `json: "created_at"`
	UpdatedAt   time.Time `json: "updated_at"`
}
