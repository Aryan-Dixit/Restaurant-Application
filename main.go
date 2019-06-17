package main

import (
	"fmt"
	"os"
	"strings"

	dbrepo "MAD_Rest_Assign/dbrepository"
	mongoutils "MAD_Rest_Assign/utils"
)

func main() {
	mongoSession, _ := mongoutils.RegisterMongoSession(os.Getenv("MONGO_HOST"))

	dbname := "restaurant"
	repoaccess := dbrepo.NewMongoRepository(mongoSession, dbname)
	fmt.Println(repoaccess)
	firstCommand := os.Args[1]
	secondCommand := os.Args[2]
	s := strings.Split(secondCommand, "=")
	s[0] = s[0][2:]
	var count int
	if firstCommand == "find" {
		if s[0] == "type_of_food" {
			result, _ := repoaccess.FindByTypeOfFood(s[1])
			fmt.Println(result)
		} else if s[0] == "postcode" {
			result, _ := repoaccess.FindByTypeOfPostCode(s[1])
			fmt.Println(result)
		} else {
			fmt.Println("Invalid arguments")
		}
	} else if firstCommand == "count" {
		if s[0] == "type_of_food" {
			count, _ = repoaccess.CountRestaurantByFood(s[1])
			fmt.Println("count Restaurant ByFood ", count)
		} else {
			fmt.Println("Invalid arguments")
		}
	} else {
		fmt.Println("Invalid arguments")
	}
}
