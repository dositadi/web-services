package main

import (
	"db"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

// Todo: Implement the Push function into database

func InitializeServer() (map[string]db.EmployeeDetails, error) {
	jsonData, err := db.ReadFromDB()
	if err != nil {
		return nil, err
	}

	if len(jsonData) == 0 {
		records, err := json.MarshalIndent(db.Employees, "", " ")
		if err != nil {
			return nil, fmt.Errorf("Error initializing database: %w", err)
		}
		err1 := db.WriteToDB(records)
		if err1 != nil {
			return nil, fmt.Errorf("Error initializing database: %w", err1)
		}

		// Reading from the DB after writing to it
		jsonData, err := db.ReadFromDB()
		if err != nil {
			return nil, err
		}

		storeRecords := make(map[string]db.EmployeeDetails)

		err2 := json.Unmarshal(jsonData, &storeRecords)
		if err2 != nil {
			return nil, fmt.Errorf("Unmarshalling Error: %s", err2)
		}
		return storeRecords, nil
	}

	records := make(map[string]db.EmployeeDetails)

	err2 := json.Unmarshal(jsonData, &records)
	if err2 != nil {
		return nil, fmt.Errorf("Unmarshalling Error: %s", err2)
	}
	return records, nil
}

func validateTable(name string) bool {
	return name == "Employees"
}

func validateKey(key string) bool {
	_, ok := db.Employees[key]
	return ok
}

func GetEmployeeDetails(records map[string]db.EmployeeDetails, table string, key string) ([]byte, error) {
	if !validateTable(table) {
		return nil, fmt.Errorf("Resource Fetch Error: Table does not exist!")
	}
	if !validateKey(key) {
		return nil, fmt.Errorf("Resource Fetch Error: Incorrect key!")
	}

	value := records[key]
	employeeBuffer, err := json.Marshal(value)
	if err != nil {
		return nil, fmt.Errorf("Error marshalling result. Err: %w", err)
	}
	return employeeBuffer, nil
}

func main() {
	records, err := InitializeServer()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var DBKeys []string

	for key := range records {
		DBKeys = append(DBKeys, key)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		table := r.FormValue("table")
		id := r.FormValue("id")

		result, err := GetEmployeeDetails(records, table, id)
		if err != nil {
			fmt.Fprintf(w, "Error Fetching Results: %v", err)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(result)
	})

	http.HandleFunc("/Employees", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Fatalln("Error reading post body: ", err)
		}

		var employee db.EmployeeDetails

		err1 := json.Unmarshal(body, &employee)
		if err1 != nil {
			log.Fatalln("Unable to unmarshal employee json: ", err1)
		}

		key := "add" + strconv.Itoa(rand.Intn(500))
		DBKeys = append(DBKeys, key)
		EmployeeExists := false

		for _, key := range DBKeys {
			if records[key].ID == employee.ID {
				EmployeeExists = true
				break
			}
		}

		if !EmployeeExists {
			records[key] = employee

			recordsJson, err2 := json.MarshalIndent(records, "", " ")
			if err2 != nil {
				log.Fatalln("Error marshalling updated records to json: ", err2)
			}

			db.WriteToDB(recordsJson)

			w.WriteHeader(http.StatusCreated)
			fmt.Fprintf(w, "Employee Post request successful for %s\n", employee.ID)
		} else {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Employee already exists for %s\n", employee.ID)
		}
	})

	log.Fatalln(http.ListenAndServe(":8080", nil))
}
