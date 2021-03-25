package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

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
		c.l.Println("about to process put request")
		c.l.Println("PUT", r.URL.Path)

		rp := strings.ReplaceAll(r.URL.Path, "/", "")

		c.l.Println("rp = ", rp)
		// expect the id in the URI
		// reg := regexp.MustCompile(`/^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$`)

		//parse string into uuid
		// Parsing UUID from string input
		u2, err := uuid.FromString(rp)
		if err != nil {
			http.Error(rw, "Invalid id", http.StatusBadRequest)
			return
		}
		fmt.Printf("Successfully parsed: %s \n", u2)

		c.updateClass(u2, rw, r)
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
		c.l.Println("cant parse json coming as request body")
		http.Error(rw, "unable to unmarshal json", http.StatusBadRequest)
		return
	}
	c.l.Println("call data.updateClass")
	err = data.UpdateClass(id, class)
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

func (c *Classes) addClass(rw http.ResponseWriter, r *http.Request) {
	c.l.Println("handle POST Class")

	class := &data.Class{}
	err := class.FromJSON(r.Body)

	if err != nil {
		http.Error(rw, "post request body is not valid", http.StatusBadRequest)
		return
	}

	c.l.Println("Class to add : %#v", class)

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
