package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"strings"

	"gopkg.in/yaml.v3"
)

func sample() {
	fmt.Println("trying out json file reading and operating")

	content, err := ioutil.ReadFile("/home/varshan/Documents/exp/config.json")
	if err != nil {
		log.Fatal("Error when opening the file: ", err)
	}
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the name to be changed : ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)
	fmt.Println("Name to be updates : ")
	newName, err1 := reader.ReadString('\n')
	newName = strings.TrimSpace(newName)
	if err1 != nil {
		log.Fatal("Error when reading the name : ", err1)
	}
	var payload []map[string]interface{}

	err = json.Unmarshal(content, &payload)

	fmt.Println("type of name is ", reflect.TypeOf(name))
	//fmt.Println("type is ", reflect.TypeOf(test))
	//var test string
	for i := range payload {
		test := payload[i]["name"]
		if test == name {
			payload[i]["name"] = newName
			fmt.Println("Updated name ", payload[i]["name"])
		}
	}

	if err != nil {
		log.Fatal("Error reading the file : ", err)
	}

	data, err2 := json.Marshal(payload)
	if err2 != nil {
		log.Fatal("Error when marshalling the data : ", err2)
	}
	ioutil.WriteFile("/home/varshan/Documents/exp/config.json", data, os.ModePerm)

	fmt.Println("Sample data ", payload)
	//fmt.Println("Name of the person is ", payload[0]["name"])
}

func sample_yaml() {

	yaml_file, err := ioutil.ReadFile("/home/varshan/Documents/learn/crd.yaml")

	if err != nil {
		log.Fatal("Error with reading yaml : ", err)
	}
	var yaml_payload map[string]interface{}

	err = yaml.Unmarshal(yaml_file, &yaml_payload)

	fmt.Println("Printing apiversion ", yaml_payload["apiVersion"])

}

func main() {
	sample()
	sample_yaml()
}
