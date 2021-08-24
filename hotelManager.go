package hotelManager

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
	Name          string `json:"name"`
	Phone         string `json:"phone"`
	ArrivalDate   string `json:"arrival_date"`
	DepartureDate string `json:"departure_date"`
	GuestsNumber  int    `json:"guests_number"`
	IsBooking     bool   `json:"is_booking"`
	Comment       string `json:"comment"`
	Status        int    `json:"status"`
}

type UsersRoom struct {
	Id     int
	UserId int
	RoomId int
}

type RoomsBooking struct {
	Id        int
	RoomId    int
	BookingId int
}
