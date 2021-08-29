package hotelManager

import "errors"

type Room struct {
	Id          int    `json:"id" db:"id"`
	RoomNumber  int    `json:"room_number" db:"room_number"`
	DoubleBed   int    `json:"double_bed" db:"double_bed"`
	SingleBed   int    `json:"single_bed" db:"single_bed"`
	Description string `json:"description" db:"description"`
	Price       int    `json:"price" db:"price"`
}

type Booking struct {
	Id            int    `json:"id"`
	Name          string `json:"name" db:"name"`
	Phone         string `json:"phone" db:"phone"`
	ArrivalDate   string `json:"arrival_date" db:"arrival_date"`
	DepartureDate string `json:"departure_date" db:"departure_date"`
	GuestsNumber  int    `json:"guests_number" db:"guests_number"`
	IsBooking     bool   `json:"is_booking" db:"is_booking"`
	Comment       string `json:"comment" db:"comment"`
	Status        int    `json:"status" db:"status"`
}

type UsersRoom struct {
	Id     int `json:"id"`
	UserId int `json:"user_id" db:"user_id"`
	RoomId int `json:"room_id" db:"room_id"`
}

type RoomsBooking struct {
	Id        int `json:"id"`
	RoomId    int `json:"room_id" db:"room_id"`
	BookingId int `json:"booking_id" db:"booking_id"`
}

type UpdateRoomInput struct {
	RoomNumber  *int    `json:"room_number"`
	DoubleBed   *int    `json:"double_bed"`
	SingleBed   *int    `json:"single_bed"`
	Description *string `json:"description"`
	Price       *int    `json:"price"`
}

func (i UpdateRoomInput) Validate() error {
	if i.RoomNumber == nil && i.DoubleBed == nil && i.SingleBed == nil && i.Description == nil && i.Price == nil {
		return errors.New("update structure has no values")
	}

	return nil
}
