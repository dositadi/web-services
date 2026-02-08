package main

import (
	"db"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

const URL = "http://localhost:8080"

func String(employee db.EmployeeDetails) string {
	result := fmt.Sprintf("---Employee Details---\nID: %s\nName: %s\nDepartment: %s\nState: %s\nCity: %s\nStreet: %s\n", employee.ID, employee.Name, employee.Department, employee.Address.State, employee.Address.City, employee.Address.Street)
	return result
}

func GetData(table string, id string) {
	query := fmt.Sprintf("?table=%s&id=%s", table, id)
	getUrl := URL + query

	response, err := http.Get(getUrl)
	if err != nil {
		log.Fatal("Get Error: ", err)
	}

	defer response.Body.Close()

	responseBuffer, err2 := io.ReadAll(response.Body)
	if err2 != nil {
		log.Fatal("Response Read Error: ", err2)
	}

	if responseBuffer[0] != '{' {
		fmt.Println(string(responseBuffer))
		os.Exit(1)
	}

	var employeeDetail db.EmployeeDetails

	err3 := json.Unmarshal(responseBuffer, &employeeDetail)
	if err3 != nil {
		log.Fatal("JSON Error: ", err3)
	}
	fmt.Println(response.Status)
	fmt.Print(String(employeeDetail))
}

func main() {
	/* table := "Employees"
	id := "add159" */

	postUrl := URL + "/Employees"
	postBody := `{"ID":"user-20","Name":"Enebe Peace","Department": "Human Resources","Address":{"Street":"1 Jude Avenue","City": "Aroko","State": "Benue"}}`

	jsondata := strings.NewReader(postBody)

	//GetData(table, id)

	response, err := http.Post(postUrl, "employee/upload", jsondata)
	if err != nil {
		log.Fatalln("Error posting to DB: ", err)
	}

	defer response.Body.Close()

	resp, err2 := io.ReadAll(response.Body)
	if err2 != nil {
		log.Fatalln("Error reading response: ", err2)
	}

	fmt.Println(string(resp))
}
