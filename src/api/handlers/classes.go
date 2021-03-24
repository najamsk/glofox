package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/najamsk/glofox/src/api/data"
	uuid "github.com/satori/go.uuid"
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

	if r.Method == http.MethodPut {
		c.l.Println("PUT", r.URL.Path)
		// expect the id in the URI
		reg := regexp.MustCompile(`/([0-9]+)`)
		g := reg.FindAllStringSubmatch(r.URL.Path, -1)

		if len(g) != 1 {
			c.l.Println("Invalid URI more than one id")
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}

		if len(g[0]) != 2 {
			c.l.Println("Invalid URI more than one capture group")
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}

		idString := g[0][1]
		id, err := strconv.Atoi(idString)
		if err != nil {
			c.l.Println("Invalid URI unable to convert to numer", idString)
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}

		fmt.Println("change this id into uuid.uuid and pass it on instead of int", id)

		c.updateClass(uuid.NewV4(), rw, r)
		return
	}
	//catch all
	rw.WriteHeader(http.StatusMethodNotAllowed)
}
func (c *Classes) updateClass(id uuid.UUID, rw http.ResponseWriter, r *http.Request) {
	c.l.Println("handle PUT Class")
	class := &data.Class{}

	err := class.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "unable to unmarshal json", http.StatusBadRequest)
		return
	}
	err = data.UpdateClass(id, class)
	if err == data.ErrClassNotFound {
		http.Error(rw, "class not found", http.StatusNotFound)
		return
	}
	if err == nil {
		http.Error(rw, "class not found", http.StatusInternalServerError)
		return
	}

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

	//2. doing same json convesino using encoder that is much better
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "unable to send data in json format", http.StatusInternalServerError)
		return
	}
}
