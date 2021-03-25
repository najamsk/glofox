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
	data.AddClass(&class)

	//marshal json
	j, err := json.Marshal(class)
	if err != nil {
		http.Error(rw, "unable to send data in json format", http.StatusInternalServerError)
		return

	}
	rw.Write(j)
}

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
			http.Error(rw, "Error validating class", http.StatusBadRequest)
			return
		}

		//add the class to the context
		ctx := context.WithValue(r.Context(), KeyClass{}, class)
		r = r.WithContext(ctx)

		next.ServeHTTP(rw, r)
	})

}
