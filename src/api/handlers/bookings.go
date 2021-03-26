package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/najamsk/glofox/src/api/data"
)

type Bookings struct {
	l *log.Logger
}

//Keyvalue to set contextwithvalue
type KeyBooking struct{}

func NewBookings(l *log.Logger) *Bookings {
	return &Bookings{l}
}

func (b *Bookings) GetBookings(rw http.ResponseWriter, r *http.Request) {
	b.l.Println("hanle GET Bookings")

	lb := data.GetBookings()

	err := lb.ToJSON(rw)
	if err != nil {
		http.Error(rw, "unable to serialize bookings in json format", http.StatusInternalServerError)
		return
	}
}

func (b *Bookings) AddBooking(rw http.ResponseWriter, r *http.Request) {
	b.l.Println("handle POST Bookings")
	//get booking from the request context
	booking := r.Context().Value(KeyBooking{}).(data.Booking)
	output, err := data.AddBooking(&booking)

	if err == data.ErrClassNotFound {
		b.l.Println("adding booking failed with error", err.Error())
		http.Error(rw, err.Error(), http.StatusNotFound)
		return
	}

	if err != nil {
		b.l.Println("adding booking failed with error", err.Error())
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	//marshal json
	j, err := json.Marshal(output)
	if err != nil {
		b.l.Println("created booking cant be serialzied")
		http.Error(rw, "unable to send bookings in json format", http.StatusInternalServerError)
		return

	}

	rw.WriteHeader(http.StatusCreated)
	rw.Write(j)
}

func (b Bookings) MiddlewareBookingValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		b.l.Println("middlewareBookingValidatoin kicks in")
		booking := data.Booking{}

		err := booking.FromJSON(r.Body)
		if err != nil {
			b.l.Println("[Error] deserializing booking from request", err.Error())
			http.Error(rw, "Error reading booking", http.StatusBadRequest)
			return
		}

		err = booking.Validate()
		if err != nil {
			b.l.Println("booking validation fails with ", err.Error())
			http.Error(rw, fmt.Sprintf("Error validating booking %s", err), http.StatusBadRequest)
			return
		}

		//add the class to the context
		ctx := context.WithValue(r.Context(), KeyBooking{}, booking)
		r = r.WithContext(ctx)

		next.ServeHTTP(rw, r)
	})

}
