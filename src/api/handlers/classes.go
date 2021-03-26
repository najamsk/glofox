// Package classification of Class API
//
// Documentation for Class API
//
//	Schemes: http
//	BasePath: /
//	Version: 1.0.0
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
// swagger:meta
package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/najamsk/glofox/src/api/data"
	uuid "github.com/satori/go.uuid"
)

// A list of classes returns int the response
// swagger:response classesResponse
type classesResponse struct {
	// All classes in the store
	// in: body
	Body []data.Class
}

type Classes struct {
	l *log.Logger
}

//Keyvalue to set contextwithvalue
type KeyClass struct{}

func NewClasses(l *log.Logger) *Classes {
	return &Classes{l}
}

func (c *Classes) UpdateClass(rw http.ResponseWriter, r *http.Request) {
	c.l.Println("handle PUT Class")

	vars := mux.Vars(r)
	id := vars["id"]
	// Parsing UUID from string input
	u2, err := uuid.FromString(id)
	if err != nil {
		c.l.Println("cant parse id of class")
		http.Error(rw, "class id is not valid please provide valid uuid", http.StatusBadRequest)
		return
	}
	fmt.Printf("Successfully parsed: %s", u2)

	//get class from request contxt since middle ware set it
	class := r.Context().Value(KeyClass{}).(data.Class)

	c.l.Println("call data.updateClass")
	err = data.UpdateClass(u2, &class)

	if err == data.ErrClassNotFound {
		c.l.Println("class to update not found")
		http.Error(rw, "class not found", http.StatusNotFound)
		return
	}
	if err != nil {
		c.l.Println("some error on server =", err.Error())
		http.Error(rw, "class not found", http.StatusInternalServerError)
		return
	}

}

func (c *Classes) AddClass(rw http.ResponseWriter, r *http.Request) {
	c.l.Println("handle POST Class")
	class := r.Context().Value(KeyClass{}).(data.Class)
	err := data.AddClass(&class)
	if err != nil {
		c.l.Printf("eror adding class %s \n", err.Error())
		http.Error(rw, err.Error(), http.StatusNotAcceptable)
		return
	}

	//marshal json
	j, err := json.Marshal(class)
	if err != nil {
		http.Error(rw, "unable to send classes in json format", http.StatusInternalServerError)
		return

	}
	rw.Write(j)
}

// swagger:route GET / classes listClasses
// Return a list of classes from the data store
// responses:
//	200: classesResponse

// GetClasses returns the classes list from the data store
func (c *Classes) GetClasses(rw http.ResponseWriter, r *http.Request) {
	c.l.Println("handle GET Classes")
	lp := data.GetClases()

	//2. doing same json convesino using encoder that is much better
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "unable to send data in json format", http.StatusInternalServerError)
		return
	}
}

func (c Classes) MiddlewareClassValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		c.l.Println("middlewareClassValidatoin kicks in")
		class := data.Class{}

		err := class.FromJSON(r.Body)
		if err != nil {
			c.l.Println("[Error] deserializing class from request", err.Error())
			http.Error(rw, "Error reading class", http.StatusBadRequest)
			return
		}

		err = class.Validate()
		if err != nil {
			c.l.Println("class validation fails with ", err.Error())
			http.Error(rw, fmt.Sprintf("Error validating class %s", err), http.StatusBadRequest)
			return
		}

		//add the class to the context
		ctx := context.WithValue(r.Context(), KeyClass{}, class)
		r = r.WithContext(ctx)

		next.ServeHTTP(rw, r)
	})

}
