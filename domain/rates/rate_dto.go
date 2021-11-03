package rates

import "time"

type RoomRate roomRate

type roomRate struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	RateTypeId  int       `json:"rate_type_id"`
	RoomTypeId  int       `json:"room_type_id"`
	PromotionId int       `json:"promotion_id"`
	Rate        float32   `json:"rate"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type promotionRate struct {
	Id          int       `json:id`
	Title       string    `json:"title"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	PromotionId int       `json:"promotion_id"`
	RoomTypeId  int       `json:"room_type_id"`
	RateTypeId  int       `json: rate_type_id`
	Rate        float32   `json:"rate"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
