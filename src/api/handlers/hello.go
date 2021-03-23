package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("hello endpoint has been called")
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("err reading body %v", err.Error())
		//set error long menthod
		// rw.WriteHeader(http.StatusBadRequest)
		// rw.Write([]byte("oops"))
		// return
		//or shortcut method
		http.Error(rw, "oops", http.StatusBadRequest)
		return
	}

	h.l.Printf("request body is %s \n", d)
	fmt.Fprintf(rw, "handler says body was  %s", d)
}
