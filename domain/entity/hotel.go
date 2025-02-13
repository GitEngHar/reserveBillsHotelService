package entity

type Hotel struct {
	ID             int
	Name           string
	PricePerNight  int
	RoomsAvailable int
	//Address        Address
}

func NewHotel(id int, name string, pricePerNight int, roomsAvailable int /*, address Address*/) *Hotel {
	return &Hotel{
		ID: id, Name: name,
		PricePerNight:  pricePerNight,
		RoomsAvailable: roomsAvailable,
	}
}
