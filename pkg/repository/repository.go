package repository

type Authorization interface {
}

type Room interface {
}

type Booking interface {
}

type Repository struct {
	Authorization
	Booking
	Room
}

func NewRepository() *Repository {
	return &Repository{}
}
