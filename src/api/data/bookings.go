package data

import (
	"encoding/json"
	"io"

	"github.com/go-playground/validator/v10"
	uuid "github.com/satori/go.uuid"
)

// const (
// 	layoutISO = "2006-01-02"
// 	layoutUS  = "January 2, 2006"
// )

type Booking struct {
	ID          uuid.UUID `json:"id"`
	ClientName  string    `json:"clientName" validate:"required"`
	BookingDate string    `json:"bookingDate" validate:"required"`
	ClassID     uuid.UUID `json:"classId" validate:"required"`
}

type Bookings []*Booking

func (b *Booking) Validate() error {
	validate := validator.New()
	return validate.Struct(b)
}

func (b *Booking) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(b)
}

func (b *Bookings) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(b)
}

func AddBooking(b *Booking) (*Booking, error) {
	err := b.Validate()
	if err != nil {
		return nil, err
	}

	//locate class by id
	_, _, err = FindClass(b.ClassID)
	if err != nil {
		return nil, ErrClassNotFound
	}

	b.ID = uuid.NewV4()
	bookingList = append(bookingList, b)
	return b, nil
}

func GetBookings() Bookings {
	return bookingList
}

//strore
var bookingList = []*Booking{}
