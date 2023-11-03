package main

import (
	"log"
	"pkg/database"
)

func main() {

	var err error

	db := &database.Db{}
	err = db.LoadDatabase("db.json")
	if err != nil {
		log.Printf("Unable to load DB file: %s\n", err)
		return
	}
	defer db.Close()

	// // example 1

	// doc := map[string]interface{}{
	// 	"title":    "The Fairy Tale",
	// 	"author":   "Stephen King",
	// 	"price":    20.5,
	// 	"in_stock": true,
	// }

	// err = db.Insert(doc)
	// if err != nil {
	// 	log.Printf("Could not insert doc: %s\n", err)
	// 	return
	// }

	// // example 2

	// doc = map[string]interface{}{
	// 	"title":    "Dagon",
	// 	"author":   "Hovard Lovecraft",
	// 	"price":    12.99,
	// 	"in_stock": false,
	// }

	// err = db.Insert(doc)
	// if err != nil {
	// 	log.Printf("Could not insert doc: %s\n", err)
	// 	return
	// }

	// // example 3

	// var doc map[string]interface{}
	// for i := 1; i <= 10000; i++ {

	// 	doc = map[string]interface{}{
	// 		"title":    fmt.Sprintf("Some book %d", i),
	// 		"author":   fmt.Sprintf("Some author %d", i),
	// 		"price":    float64(i) + 0.44,
	// 		"in_stock": false,
	// 	}
	// 	err = db.Insert(doc)
	// 	if err != nil {
	// 		log.Printf("Could not insert doc: %s\n", err)
	// 		return
	// 	}
	// }

	// find

	data, err := db.Find(`{"$or": [{"title": {"$contains": "1999"}}, {"price": "9.44"}]}`)
	if err != nil {
		log.Printf("Could not find doc: %s\n", err)
		return
	}
	println(data)

}
