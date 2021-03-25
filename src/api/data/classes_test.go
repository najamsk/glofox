package data

import "testing"

func TestCheckValidation(t *testing.T) {
	c := &Class{Name: "java", StartDate: " ", EndDate: "2021-04-25"}

	err := c.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
