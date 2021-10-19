package domain_promotions

import "time"

type promotionRoomRate struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	RoomTypeId  int       `json:"room_type_id"`
	PromotionId int       `json:"promotion_id"`
	Rate        float32   `json:"rate"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
