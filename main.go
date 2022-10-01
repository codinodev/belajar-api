package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"text/template"
)

type Employee struct {
	Id       int
	Name     string
	Age      int
	Division string
}

var employee = []Employee{
	{Id: 1, Name: "budi", Age: 22, Division: "Developer"},
	{Id: 2, Name: "eko", Age: 23, Division: "Developer"},
	{Id: 3, Name: "kiki", Age: 20, Division: "UI"},
}

var port = ":8080"

func main() {
	http.HandleFunc("/", getEmployees)
	http.HandleFunc("/employees", createEmpoyees)

	fmt.Println("Application is listening on port", port)
	http.ListenAndServe(port, nil)
}

func getEmployees(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	if request.Method == "GET" {
		tpl, err := template.ParseFiles("template.html")

		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		tpl.Execute(writer, employee)

		// json.NewEncoder(writer).Encode(employee)
		// return
	}

	http.Error(writer, "Invalid method", http.StatusBadRequest)
}

func createEmpoyees(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Contnt-Type", "application/json")

	if request.Method == "POST" {
		name := request.FormValue("name")
		age := request.FormValue("age")
		division := request.FormValue("division")

		convertAge, err := strconv.Atoi(age)

		if err != nil {
			http.Error(writer, "Invalid age", http.StatusBadRequest)
			return
		}

		newEmployees := Employee{
			Id:       len(employee) + 1,
			Name:     name,
			Age:      convertAge,
			Division: division,
		}

		employee = append(employee, newEmployees)
		json.NewEncoder(writer).Encode(newEmployees)
		return
	}

	http.Error(writer, "Invalid method", http.StatusBadRequest)
}
