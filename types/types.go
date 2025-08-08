package types

import (
	"fmt"
	"strings"
)

type Rider struct {
	RiderID int     `json:"riderID"`
	Name    string  `json:"name"`
	Lat     float64 `json:"lat"`
	Long    float64 `json:"long"`
}

// ------------------- Hotel Domain -------------------

type HotelRating int

const (
	MinHotelRating HotelRating = 1
	MaxHotelRating HotelRating = 5
)

func NewHotelRating(value int) HotelRating {
	if value < int(MinHotelRating) {
		value = int(MinHotelRating)
	}
	if value > int(MaxHotelRating) {
		value = int(MaxHotelRating)
	}
	return HotelRating(value)
}

func (r HotelRating) Stars() string {
	full := strings.Repeat("★", int(r))
	empty := strings.Repeat("☆", int(MaxHotelRating)-int(r))
	return full + empty
}

func (r HotelRating) String() string {
	return fmt.Sprintf("%d star(s): %s", r, r.Stars())
}

type Hotel struct {
	ID       int         `json:"id"`
	Name     string      `json:"name"`
	Location string      `json:"location"`
	Rating   HotelRating `json:"rating"`
}

// ------------------- Customer Domain -------------------

type Customer struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Email string `json:"email"`
	// You could also add Address here if needed
}

// ------------------- Cart and Order Domain -------------------

type CartItem struct {
	ID       int     `json:"id"`       // Unique ID for item
	Name     string  `json:"name"`     // e.g., "Veg Burger"
	Quantity int     `json:"quantity"` // How many of this item
	Price    float64 `json:"price"`    // Price per item
}

type Order struct {
	OrderID     int        `json:"orderID"`
	CustomerID  int        `json:"customerID"`
	HotelID     int        `json:"hotelID"`
	Items       []CartItem `json:"items"`
	TotalAmount float64    `json:"totalAmount"`
	Status      string     `json:"status"` // e.g., "placed", "accepted", "delivered"
	CreatedAt   string     `json:"createdAt"`
}
