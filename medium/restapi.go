package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os/exec"

	"github.com/gorilla/mux"
)

// func getDate(w http.ResponseWriter, r *http.Request) {

// 	output := script.Exec("date")

// 	// if err != nil {
// 	// 	log.Fatal("Error when executing the command")
// 	// }

// 	var output_json map[string]interface{}
// 	output_json["status"] = "success"
// 	output_json["output"] = output

// 	//response := json.Marshal(output_json)
// 	json.NewEncoder(w).Encode(output_json)
// }

// func main() {
// 	log.Fatal("Starting the API server")
// 	router := mux.NewRouter()
// 	router.HandleFunc("/", getDate).Methods("GET")
// 	log.Fatal("Router created")
// 	//router.HandleFunc("/", getDate).Methods("GET")
// 	log.Fatal(http.ListenAndServe(":8000", router))

// }

func getDate(w http.ResponseWriter, r *http.Request) {

	output, err := exec.Command("date").Output() //script.Exec("date")

	if err != nil {
		log.Fatal("Command not executed")
	}
	log.Println(string(output))
	//ot_decode, err := base64.StdEncoding.DecodeString(string(output))

	// if err != nil {
	// 	log.Fatal("Error when decoding the output")
	// }

	var output_json = make(map[string]interface{})
	output_json["status"] = "success"
	output_json["output"] = string(output)

	json.NewEncoder(w).Encode(output_json)
}

func main() {
	log.Println("Starting the API server...")
	router := mux.NewRouter()
	router.HandleFunc("/", getDate).Methods("GET")
	log.Println("Router created.")
	log.Fatal(http.ListenAndServe(":8000", router))
}
