package roomrates

import "time"

type RoomRate roomRate
type RateType rateType

type roomRate struct {
	Id              int       `json:"id"`
	Title           string    `json:"title"`
	Description     string    `json:"description"`
	RateTypeId      int       `json:"rate_type_id"`
	RoomTypeId      int       `json:"room_type_id"`
	PromotionTypeId int       `json:"promotion_type-id"`
	Rate            float32   `json:"rate"`
	StartDate       time.Time `json:"start_date"`
	EndDate         time.Time `json:"end_date"`
	Status          string    `json:"status"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

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
	ShortName   string    `json:"short_name"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type promotionRate struct {
	id               int       `json:id`
	StartDate        time.Time `json:"start_date"`
	EndDate          time.Time `json:"end_date"`
	PromotionTypeId  int       `json:"promotion_type_id"`
	PromotionPackage int       `json:"promotion_package_id"`
	RoomTypeId       int       `json:"room_type_id"`
	rateTypeId       int       `json: rateType`
	Status           string    `json:"status"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}
