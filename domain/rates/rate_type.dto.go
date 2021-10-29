package rates

import "time"

// Rate type
// description: default rate is rackrate
// need roomrate with auto generate for all room type
// Rate Types
// - Rack rate / rr
// - whole sale / ws
// - oversea travel a/ ota
// - Member / member
// - Very importance Person / VIP
// - Promotion / pmt
//	- package
// 	- discount
// - Complimentary / com
type rateType struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Acronym     string    `json:"acronym"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
