package database

import (
	"bufio"
	"encoding/json"
	"errors"
	"log"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
)

type Db struct {
	storage      []map[string]interface{}
	inMemoryOnly bool
	pathToFile   string
	file         *os.File
}

func (db *Db) LoadDatabase(filePath string) error {

	var err error
	db.file, err = os.OpenFile(filePath, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0600)
	if err != nil {
		return err
	}

	fileInfo, err := db.file.Stat()
	if err != nil {
		return err
	}
	if fileInfo.Size() > 0 {

		scanner := bufio.NewScanner(db.file)

		var line string
		for scanner.Scan() {

			line = scanner.Text()
			if err := scanner.Err(); err != nil {
				return err
			}

			var doc map[string]interface{}
			err := json.Unmarshal([]byte(line), &doc)
			if err != nil {
				return err
			}
			db.storage = append(db.storage, doc)

		}

	}

	db.pathToFile = filePath
	db.inMemoryOnly = false

	return nil
}

func (db *Db) Close() {

	if db.inMemoryOnly == true {
		return
	}
	db.file.Close()
}

func (db *Db) prepareDocForInsert(doc map[string]interface{}) (map[string]interface{}, error) {
	_, ok := doc["_id"]
	if ok {
		return doc, errors.New("Unacceptable key: _id")
	}
	doc["_id"] = uuid.New().String()

	_, ok = doc["_created"]
	if ok {
		return doc, errors.New("Unacceptable key: _created")
	}
	now := time.Now().UTC()
	doc["_created"] = now

	_, ok = doc["_updated"]
	if ok {
		return doc, errors.New("Unacceptable key: _updated")
	}
	return doc, nil
	doc["_updated"] = now

	return doc, nil
}

func (db *Db) Insert(doc map[string]interface{}) error {

	var err error
	doc, err = db.prepareDocForInsert(doc)
	if err != nil {
		return err
	}
	db.storage = append(db.storage, doc)

	docJson, err := json.Marshal(doc)
	if err != nil {
		return err
	}

	if _, err = db.file.WriteString(string(docJson) + "\n"); err != nil {
		return err
	}

	return nil
}

func (db *Db) Find(queryString string) (string, error) {

	var err error
	var queryMap map[string]interface{}
	err = json.Unmarshal([]byte(queryString), &queryMap)
	if err != nil {
		return "", errors.New("Could not unmarshal JSON")
	}

	for key := range queryMap {
		if strings.HasPrefix(key, "$") {
			log.Println(key)
		}
	}

	return "", nil
}

func (db *Db) matchQuery(query map[string]interface{}) (bool, error) {
	return true, nil
}
