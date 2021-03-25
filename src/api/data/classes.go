package data

import (
	"encoding/json"
	"fmt"
	"io"
	"time"

	"github.com/go-playground/validator/v10"
	uuid "github.com/satori/go.uuid"
)

const (
	layoutISO = "2006-01-02"
	layoutUS  = "January 2, 2006"
)

type Class struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name" validate:"required"`
	StartDate string    `json:"startDate" validate:"required,beforeEndDate"`
	EndDate   string    `json:"endDate" validate:"required"`
	Capacity  int       `json:"capacity"`
}

type Classes []*Class

var ErrClassNotFound = fmt.Errorf("Class not found")

func (c Class) validateStartDate(fl validator.FieldLevel) bool {

	sDate := fl.Field().String()

	fmt.Println("SDate =", sDate)

	tStart, err := time.Parse(layoutISO, sDate)
	if err != nil {
		fmt.Println("validattion tStart fails = ", tStart)
		return false
	}
	tEnd, err := time.Parse(layoutISO, c.EndDate)
	if err != nil {
		fmt.Println("validattion tEnd fails = ", tEnd)
		return false
	}

	if tStart.After(tEnd) {
		return false
	}

	return true

}

func (c *Class) Validate() error {
	validate := validator.New()

	validate.RegisterValidation("beforeEndDate", c.validateStartDate)
	return validate.Struct(c)
}

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

func UpdateClass(id uuid.UUID, c *Class) error {
	_, pos, err := findClass(id)

	if err != nil {
		return err
	}
	c.ID = id
	classList[pos] = c
	return nil
}
func findClass(id uuid.UUID) (*Class, int, error) {
	for k, v := range classList {

		if uuid.Equal(v.ID, id) {
			return v, k, nil
		}
	}
	return nil, -1, ErrClassNotFound
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

// startDate := "2021-02-14"
// tStart, _ := time.Parse(layoutISO, startDate)
// endDate := "2021-02-22"
