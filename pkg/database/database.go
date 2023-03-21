package database

import (
	"encoding/json"
	"log"
	"os"
)

type Db struct {
	PathToFile string
}

func (db Db) Insert(doc map[string]interface{}) error {

	f, err := os.OpenFile(db.PathToFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}

	defer f.Close()

	docJson, err := json.Marshal(doc)
	if err != nil {
		log.Printf("Could not marshal json: %s\n", err)
		return err
	}

	if _, err = f.WriteString(string(docJson) + "\n"); err != nil {
		log.Printf("Could not write data to file: %s\n", err)
		return err
	}

	return nil
}
