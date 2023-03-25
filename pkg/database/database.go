package database

import (
	"bufio"
	"encoding/json"
	"os"
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

func (db *Db) Insert(doc map[string]interface{}) error {

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
