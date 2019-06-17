package mongoutils

import (
	"log"
	uuid "github.com/satori/go.uuid"
	"gopkg.in/mgo.v2"
)

type MongoAuthObject struct {
	DBname string
}

var MongoSession *mgo.Session

func RegisterMongoSession(mongoURI string) (*mgo.Session, error) {
	var err error
	MongoSession, err = mgo.Dial(mongoURI)
	if err != nil {
		log.Fatalf("Mongo Connection Error")
		panic(err)
	}
	return MongoSession, nil
}

func NewUUID() string {
	id, _ := uuid.NewV4()
	return id.String()
}
