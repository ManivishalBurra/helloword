package main

import (
	"log"
	"net/http"
	"os"

	C "github.com/ManivishalBurra/Eltrocab/controllers"
	"github.com/gorilla/mux"
)

func main() {
	port := os.Getenv("PORT")

	app := mux.NewRouter()
	app.HandleFunc("https://eltrocaps.herokuapp.com/driver", C.CreateDriver)
	app.HandleFunc("https://eltrocaps.herokuapp.com/login/driver", C.LoginDriver)
	app.HandleFunc("https://eltrocaps.herokuapp.com/cabrequests", C.FetchRequest)
	app.HandleFunc("https://eltrocaps.herokuapp.com/user", C.CreateUser)
	app.HandleFunc("https://eltrocaps.herokuapp.com/login/user", C.LoginUser)
	app.HandleFunc("https://eltrocaps.herokuapp.com/bookride", C.BookRide)
	app.HandleFunc("https://eltrocaps.herokuapp.com/driverconfirm", C.DriverConfirm)
	app.HandleFunc("https://eltrocaps.herokuapp.com/ridestatus", C.RideStatus)
	app.HandleFunc("https://eltrocaps.herokuapp.com/usercancelride", C.UserCancelRide)
	app.HandleFunc("https://eltrocaps.herokuapp.com/drivercancelride", C.DriverCancelRide)
	app.HandleFunc("https://eltrocaps.herokuapp.com/user/logout", C.UserLogout)
	app.HandleFunc("https://eltrocaps.herokuapp.com/driver/logout", C.DriverLogout)
	log.Fatal(http.ListenAndServe(":"+port, app))
}
