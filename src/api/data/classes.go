package data

import (
	"encoding/json"
	"io"
	"time"

	uuid "github.com/satori/go.uuid"
)

const (
	layoutISO = "2006-01-02"
	layoutUS  = "January 2, 2006"
)

// startDate := "2021-02-14"
// tStart, _ := time.Parse(layoutISO, startDate)
// endDate := "2021-02-22"
type Class struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	StartDate string    `json:"startDate"`
	EndDate   string    `json:"endDate"`
	Capacity  int       `json:"capacity"`
}

type Classes []*Class

func (c *Class) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(c)
}

func (c *Classes) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(c)
}

func AddClass(c *Class) {
	c.ID = uuid.NewV4()
	classList = append(classList, c)
}

func GetClases() Classes {
	return classList
}

//store
var classList = []*Class{
	&Class{
		ID:        uuid.NewV4(),
		Name:      "Easy home cooking",
		StartDate: time.Now().AddDate(0, 0, -30).String(),
		EndDate:   time.Now().AddDate(0, 0, -15).String(),
		Capacity:  20,
	},
	&Class{
		ID:        uuid.NewV4(),
		Name:      "Home workout",
		StartDate: time.Now().AddDate(0, 0, 1).String(),
		EndDate:   time.Now().AddDate(0, 0, 14).String(),
		Capacity:  10,
	},
}

// cases to handle while creating new class

// startdate should not be bigger then enddate

// you have one class in a single day so while registering a class it should not clash with any other class
