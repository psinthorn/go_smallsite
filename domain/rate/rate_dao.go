package rates

import "fmt"

var RoomRateService roomRateInterface = &RoomRate{}

type roomRateInterface interface {
	Create(RoomRate) (int, error)
	Get(string) ([]RoomRate, error)
	GetById(int) (RoomRate, error)
	Update(RoomRate) error
	Delete(int) error

	AdminGet() ([]RoomRate, error)
}

func (r *RoomRate) Create(rr RoomRate) (int, error) {
	fmt.Println("Please implement me")
	return 0, nil
}

func (r *RoomRate) Get(status string) ([]RoomRate, error) {
	var rr []RoomRate
	fmt.Println("Please implement me")
	return rr, nil
}

func (r *RoomRate) GetById(id int) (RoomRate, error) {
	var rr RoomRate
	fmt.Println("Please implement me")
	return rr, nil
}

func (r *RoomRate) Update(RoomRate) error {
	fmt.Println("Please implement me")
	return nil
}

func (r *RoomRate) Delete(id int) error {
	fmt.Println("Please implement me")
	return nil
}

func (r *RoomRate) AdminGet() ([]RoomRate, error) {
	var rr []RoomRate
	fmt.Println("Please implement me")
	return rr, nil
}
