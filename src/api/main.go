package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
	"github.com/najamsk/glofox/src/api/data"
	"github.com/najamsk/glofox/src/api/handlers"
	uuid "github.com/satori/go.uuid"
)

// var bindAddress = env.String("BIND_ADDRESS", false, ":9090", "Bind address for the server")

func debug() {
	const (
		layoutISO = "2006-01-02"
		layoutUS  = "January 2, 2006"
	)

	date := "1999-12-31"
	t, _ := time.Parse(layoutISO, date)
	fmt.Println(t)                  // 1999-12-31 00:00:00 +0000 UTC
	fmt.Println(t.Format(layoutUS)) // December 31, 1999

	fmt.Println("now = ", time.Now())
	today := time.Now().Format("2006-01-02")
	tomorrow := time.Now().AddDate(0, 0, 1).Format("2006-01-02")

	fmt.Println("start", today)
	fmt.Println("tomorrow", tomorrow)

	u2 := uuid.NewV4()
	startDate := "2021-02-14"
	tStart, _ := time.Parse(layoutISO, startDate)
	endDate := "2021-02-22"
	tEnd, _ := time.Parse(layoutISO, endDate)

	fmt.Printf("UUIDv4: %s\n", u2)
	class := &data.Class{Name: "najam awan", ID: u2, Capacity: 20, StartDate: tStart.String(), EndDate: tEnd.String()}

	fmt.Printf("class = %+v \n", class)
	//http work

}
func main() {

	l := log.New(os.Stdout, "api", log.LstdFlags)
	// hh := handlers.NewHello(l)
	// bh := handlers.NewGoodbye(l)
	ch := handlers.NewClasses(l)
	bh := handlers.NewBookings(l)

	sm := mux.NewRouter()
	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/classes", ch.GetClasses)
	getRouter.HandleFunc("/bookings", bh.GetBookings)

	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/classes", ch.AddClass)
	postRouter.Use(ch.MiddlewareClassValidation)

	postRouterBooking := sm.Methods(http.MethodPost).Subrouter()
	postRouterBooking.HandleFunc("/bookings", bh.AddBooking)
	postRouterBooking.Use(bh.MiddlewareBookingValidation)

	putRouter := sm.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/classes/{id}", ch.UpdateClass)
	putRouter.Use(ch.MiddlewareClassValidation)

	opts := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.Redoc(opts, nil)

	getRouter.Handle("/docs", sh)
	getRouter.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

	//http server launching with graceful shutdown support
	s := &http.Server{
		Addr:    ":9090",
		Handler: sm,
	}

	//use go go routine
	go func() {
		//since s.listenandserve will block we wrap it inside goroutine
		l.Println("starting server  :9090")
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	killChan := make(chan os.Signal)
	signal.Notify(killChan, os.Interrupt)
	signal.Notify(killChan, os.Kill)

	// reading from channel will block and will be unbloced if any kill interrrupt will be received
	sig := <-killChan
	l.Println("signal to shutdown, will be doing graceful shutdown", sig)

	// cleanup resources like database connections etc
	l.Println("cleaning up resources")

	//timeout context requrire for server.shutdown
	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}
