package data

import (
	"fmt"
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreateValidBooking(t *testing.T) {
	c := &Class{ID: uuid.NewV4(), Name: "java", StartDate: "2021-03-23", EndDate: "2021-04-25"}
	b := &Booking{ClientName: "Najam Awan", BookingDate: "2021-04-25", ClassID: c.ID}

	output, err := AddBooking(b)

	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(c.ID)
	fmt.Println(output.ClassID)
	assert.Equal(t, output.ClassID, c.ID)

}

func TestCreateInvalidValidBooking(t *testing.T) {
	c := &Class{ID: uuid.NewV4(), Name: "java", StartDate: "2021-03-23", EndDate: "2021-04-25"}
	b := &Booking{BookingDate: "2021-04-25", ClassID: c.ID}

	output, err := AddBooking(b)

	assert.NotNil(t, err)
	assert.Nil(t, output)

}

func TestGetListOfBookings(t *testing.T) {
	c := &Class{ID: uuid.NewV4(), Name: "java", StartDate: "2021-03-23", EndDate: "2021-04-25"}
	b := &Booking{ClientName: "tomm", BookingDate: "2021-04-25", ClassID: c.ID}
	b1 := &Booking{ClientName: "gerry", BookingDate: "2021-04-25", ClassID: c.ID}

	output, err := AddBooking(b)
	if err != nil {
		t.Fatal(err)
	}
	output2, err2 := AddBooking(b1)
	if err2 != nil {
		t.Fatal(err)
	}

	result := GetBookings()

	fmt.Printf("bookins are : %#v \n", result)

	assert.Len(t, result, 2)
	assert.Equal(t, b.ClientName, output.ClientName)
	assert.Equal(t, b1.ClientName, output2.ClientName)

}
