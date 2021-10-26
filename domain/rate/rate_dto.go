package rates

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
