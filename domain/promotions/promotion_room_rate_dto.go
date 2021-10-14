package domain

import "time"

type PromotionRoomRate struct {
	Id           int       `json:"id"`
	RoomId       int       `json:"room_id"`
	Promotion_id int       `json:"promotion_id"`
	RateId       float32   `json:"rate_id"`
	Status       string    `json:"status"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updatedAt"`
}
