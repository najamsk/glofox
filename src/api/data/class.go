package data

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Class struct {
	ID        uuid.UUID
	Name      string
	StartDate time.Time
	EndDate   time.Time
	Capacity  int
}

// cases to handle while creating new class

// startdate should not be bigger then enddate

// you have one class in a single day so while registering a class it should not clash with any other class
