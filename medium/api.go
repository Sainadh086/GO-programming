package main

import (
	"bytes"         // converting json data to bytes
	"encoding/json" // converting data to json and json to data
	"fmt"           // printing the responses in the terminal
	"io/ioutil"     // reading the response of the api
	"log"           // logging the errors
	"net/http"      // to make restapi requests
	"os"            // to exist after the error
)

type Response struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

type JsonData struct {
	Title     string `json:"title"`
	Body      string `json:"body"`
	Completed bool   `json:"completed"`
}

// get data
func get_data(url string) int {
	resp, err := http.Get(url) //"https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("Error occured in get")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	respdata, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var respObject Response
	json.Unmarshal(respdata, &respObject)

	fmt.Println("Title of the note is ", respObject.Title)
	return resp.StatusCode
}

// post data to the api
func post_data(url string, data JsonData) int {
	jsondata, _ := json.Marshal(data)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsondata))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Data from post response")
	//fmt.Println(out)
	return resp.StatusCode
}

func main() {
	url := "https://jsonplaceholder.typicode.com/todos/1"
	status := get_data(url)
	fmt.Println("API status ", status)
	url = "https://jsonplaceholder.typicode.com/posts"
	var data JsonData
	data.Title = "Just implementing GO"
	data.Body = "Learning and implementing GO, helping the community to grow"
	data.Completed = true
	resp_code := post_data(url, data)
	if resp_code == 201 {
		fmt.Println("Data added successfully")
	} else {
		fmt.Println("Data posted")
	}
}
