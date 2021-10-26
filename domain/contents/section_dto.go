package contents

import (
	"time"
)

type section struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Desc        string    `json:"desc"`
	Parent      int       `json:"parent"`
	Status      string    `json:"status"`
	AccessLevel int       `json:"access_level"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

//
// Validate content input
//

// func (st *content) Validate() error {
// 	st.Title = strings.TrimSpace(st.Title)
// 	if st.Title == "" {
// 		return errors.New("Title can not be empty!")
// 	}
// 	return nil
// }
