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
	if port == "" {
		log.Fatal("port must be set")
		port = "8080"
	}
	app := mux.NewRouter()
	app.HandleFunc(port+"/driver", C.CreateDriver)
	app.HandleFunc(port+"/login/driver", C.LoginDriver)
	app.HandleFunc(port+"/cabrequests", C.FetchRequest)
	app.HandleFunc(port+"/user", C.CreateUser)
	app.HandleFunc(port+"/login/user", C.LoginUser)
	app.HandleFunc(port+"/bookride", C.BookRide)
	app.HandleFunc(port+"/driverconfirm", C.DriverConfirm)
	app.HandleFunc(port+"/ridestatus", C.RideStatus)
	app.HandleFunc(port+"/usercancelride", C.UserCancelRide)
	app.HandleFunc(port+"/drivercancelride", C.DriverCancelRide)
	app.HandleFunc(port+"/user/logout", C.UserLogout)
	app.HandleFunc(port+"/driver/logout", C.DriverLogout)
	log.Fatal(http.ListenAndServe(":"+port, app))
}
