package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/najamsk/glofox/src/api/data"
)

type Classes struct {
	l *log.Logger
}

func NewClasses(l *log.Logger) *Classes {
	return &Classes{l}
}

func (c *Classes) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("classes request method is = ", r.Method)
	if r.Method == http.MethodGet {
		c.getClasses(rw, r)
		return
	}

	if r.Method == http.MethodPost {
		c.addClass(rw, r)
		return
	}
	//catch all
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (c *Classes) addClass(rw http.ResponseWriter, r *http.Request) {
	c.l.Println("handle POST Class")

	class := &data.Class{}
	err := class.FromJSON(r.Body)

	if err != nil {
		http.Error(rw, "post request body is not valid", http.StatusBadRequest)
		return
	}

	c.l.Println("new class : %#v", class)

	data.AddClass(class)

	//marshal json
	j, err := json.Marshal(class)
	if err != nil {
		http.Error(rw, "unable to send data in json format", http.StatusInternalServerError)
		return

	}
	rw.Write(j)

	// fmt.Fprintf(rw, "post request data was=   %v", d)

}

func (c *Classes) getClasses(rw http.ResponseWriter, r *http.Request) {
	c.l.Println("handle GET Classes")
	lp := data.GetClases()

	//1. convertiing classes type objects to json using json.marshal
	// d, err := json.Marshal(lp)
	// if err != nil {
	// 	http.Error(rw, "unable to send data in json format", http.StatusInternalServerError)
	// 	return
	// }
	// rw.Write(d)

	//2. doing same json convesino using encoder that is much better
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "unable to send data in json format", http.StatusInternalServerError)
		return
	}
}
