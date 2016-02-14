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

type Workout struct {
	ID          int    `json:"id"`
	Equipment   string `json:"equipment"`
	User        string `json:"user"`
	Name        string `json:"name"`
	Description string `json:"description"`
	//	Gym         string `json:"gym"`
}

func main() {

	svc := dynamodb.New(session.New(), &aws.Config{Endpoint: aws.String("http://localhost:8000"), Region: aws.String("us-east-1")})

	file, e := ioutil.ReadFile("./mockdata/MOCK_DATA-workout.json")
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}

	var u []Workout
	json.Unmarshal(file, &u)

	for i := 0; i < len(u); i++ {
		fmt.Printf("Name: %v\n", u[i].Name)

		item, err := dynamodbattribute.ConvertToMap(u[i])
		if err != nil {
			fmt.Println("Failed to convert", err)
			return
		}
		fmt.Printf("Item %v\n", item)

		params := &dynamodb.PutItemInput{
			Item:      item,
			TableName: aws.String("workout"),
		}
		resp, err := svc.PutItem(params)

		if err != nil {
			fmt.Println(err.Error())
			return
		}
		// Pretty-print the response data.
		fmt.Println(resp)

	}
}
