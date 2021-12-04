package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	C "github.com/ManivishalBurra/Eltrocab/controllers"
	"github.com/gorilla/mux"
)

func main() {
	port := os.Getenv("PORT")

	app := mux.NewRouter()
	app.HandleFunc("/", index_handle)
	app.HandleFunc("https://eltrocaps.herokuapp.com/driver", CreateDriver).Methods("POST")
	app.HandleFunc("https://eltrocaps.herokuapp.com/login/driver", C.LoginDriver).Methods("POST")
	app.HandleFunc("https://eltrocaps.herokuapp.com/cabrequests", C.FetchRequest).Methods("POST")
	app.HandleFunc("https://eltrocaps.herokuapp.com/user", C.CreateUser).Methods("POST")
	app.HandleFunc("https://eltrocaps.herokuapp.com/login/user", C.LoginUser).Methods("POST")
	app.HandleFunc("https://eltrocaps.herokuapp.com/bookride", C.BookRide).Methods("POST")
	app.HandleFunc("https://eltrocaps.herokuapp.com/driverconfirm", C.DriverConfirm).Methods("POST")
	app.HandleFunc("https://eltrocaps.herokuapp.com/ridestatus", C.RideStatus).Methods("POST")
	app.HandleFunc("https://eltrocaps.herokuapp.com/usercancelride", C.UserCancelRide).Methods("POST")
	app.HandleFunc("https://eltrocaps.herokuapp.com/drivercancelride", C.DriverCancelRide).Methods("POST")
	app.HandleFunc("https://eltrocaps.herokuapp.com/user/logout", C.UserLogout).Methods("POST")
	app.HandleFunc("https://eltrocaps.herokuapp.com/driver/logout", C.DriverLogout).Methods("POST")
	log.Fatal(http.ListenAndServe(":"+port, app))
}
func index_handle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "whoah, nice nice")
}
func CreateDriver(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "whoah, nice ludo")
}
