package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

//[{"id":1,"gender":"Male","first_name":"Richard","last_name":"Lee","email":"rlee0@pinterest.com","weight":199,"height":69},
type Gym struct {
	ID        int      `json:"id"`
	Name      string   `json:"name"`
	Location  string   `json:"location"`
	Equipment []string `json:"equipment"`
}

func main() {

	svc := dynamodb.New(session.New(), &aws.Config{Endpoint: aws.String("http://localhost:8000"), Region: aws.String("us-east-1")})

	//var data map[string]interface{}

	file, e := ioutil.ReadFile("./mockdata/MOCK_DATA-gym.json")
	//fmt.Print(string(file))
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}

	var u []Gym
	json.Unmarshal(file, &u)

	for i := 0; i < len(u); i++ {
		fmt.Printf(" Name: %v\n", u[i].Name)

		item, err := dynamodbattribute.ConvertToMap(u[i])
		if err != nil {
			fmt.Println("Failed to convert", err)
			return
		}
		fmt.Printf("Item %v\n", item)

		//fmt.Printf("Data %v\n", data[user])
		params := &dynamodb.PutItemInput{
			Item:      item,
			TableName: aws.String("gym"),
		}
		resp, err := svc.PutItem(params)

		if err != nil {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
			return
		}
		// Pretty-print the response data.
		fmt.Println(resp)

	}
	//fmt.Printf("AL: Data %v", data)
}
