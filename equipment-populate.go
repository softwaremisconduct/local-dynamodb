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
type Equipment struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	VideoURL    string `json:"videoUrl"`
}

func main() {

	svc := dynamodb.New(session.New(), &aws.Config{Endpoint: aws.String("http://localhost:8000"), Region: aws.String("us-east-1")})

	file, err := ioutil.ReadFile("./mockdata/MOCK_DATA-equipment.json")
	//fmt.Print(string(file))
	if err != nil {
		fmt.Printf("File error: %v\n", err)
		os.Exit(1)
	}

	var e []Equipment
	json.Unmarshal(file, &e)

	for i := 0; i < len(e); i++ {
		fmt.Printf("Data Item: %v\n", e[i].Name)

		item, conErr := dynamodbattribute.ConvertToMap(e[i])
		if conErr != nil {
			fmt.Println("Failed to convert", conErr)
			return
		}
		fmt.Printf("Item %v\n", item)

		//fmt.Printf("Data %v\n", data[user])
		params := &dynamodb.PutItemInput{
			Item:      item,
			TableName: aws.String("equipment"),
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
