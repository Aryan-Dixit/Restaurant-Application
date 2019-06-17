package main

import (
	"MAD_Rest_Assign/dbrepository"
	handlerlib "MAD_Rest_Assign/restapplication/packages/httphandlers"
	"MAD_Rest_Assign/restapplication/restaurantcrudhandler"
	mongoutils "MAD_Rest_Assign/utils"
	logger "log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func init() {
	http.DefaultClient.Timeout = time.Minute * 10
}

func main() {

	pingHandler := &handlerlib.PingHandler{}
	dbname := "restaurant"
	mongosession, err := mongoutils.RegisterMongoSession(os.Getenv("MONGO_HOST"))
	if err != nil {
		logger.Fatal("Error in creating mongoSession")
	}
	dbsession := dbrepository.NewMongoRepository(mongosession, dbname)
	handler := restaurantcrudhandler.NewRestaurantCrudHandler(dbsession)
	logger.Println("Setting up resources.")
	logger.Println("Starting service")
	h := mux.NewRouter()
	h.Handle("/restaurant/ping", pingHandler)
	h.Handle("/restaurant/", handler)
	h.Handle("/restaurant/{id}", handler)
	logger.Fatal(http.ListenAndServe(":8080", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(h)))

}
