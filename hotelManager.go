package hotelManager

type Room struct {
	Id          int    `json:"id"`
	RoomNumber  int    `json:"room_number"`
	DoubleBed   int    `json:"double_bed"`
	SingleBed   int    `json:"single_bed"`
	Description string `json:"description"`
	Price       int    `json:"price"`
}

type UsersList struct {
	Id     int
	UserId int
	RoomId int
}

type Booking struct {
	Id            int
	RoomId        int
	Name          string `json:"name"`
	Phone         string `json:"phone"`
	ArrivalDate   string `json:"arrival_date"`
	DepartureDate string `json:"departure_date"`
	GuestsNumber  int    `json:"guests_number"`
	IsBooking     bool   `json:"is_booking"`
	Comment       string `json:"comment"`
}
