package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	dbrepo "MAD_Rest_Assign/dbrepository"
	domain "MAD_Rest_Assign/domain"
	mongoutils "MAD_Rest_Assign/utils"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	mongoSession, _ := mongoutils.RegisterMongoSession(os.Getenv("MONGO_HOST"))

	dbname := "restaurant"
	repoaccess := dbrepo.NewMongoRepository(mongoSession, dbname)
	fmt.Println(repoaccess)

	file, err := os.Open("./restaurant.json")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	var res domain.Restaurant
	count := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		jsn := []byte(scanner.Text())
		json.Unmarshal(jsn, &res)
		res.DBID = domain.NewID()
		repoaccess.Store(&res)
		count = count + 1
	}
	fmt.Println(count)
}
