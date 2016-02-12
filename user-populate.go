package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

//[{"id":1,"gender":"Male","first_name":"Richard","last_name":"Lee","email":"rlee0@pinterest.com","weight":199,"height":69},
type User struct {
	ID        int    `json:"id"`
	Gender    string `json:"gender"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Weight    int    `json:"weight"`
	Height    int    `json:"height"`
	//password  string `json:"password"`
}

func main() {

	svc := dynamodb.New(session.New(), &aws.Config{Endpoint: aws.String("http://localhost:8000"), Region: aws.String("us-east-1")})

	file, e := ioutil.ReadFile("./mockdata/MOCK_DATA-user.json")
	fmt.Print(string(file))
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}

	//defer file.Close()
	u := make([]User, 0)
	json.Unmarshal(file, &u)
	//fmt.Printf("JSON Results: %v\n", u)

	for i := 0; i < len(u); i++ {
		fmt.Printf(" First Name: %v\n", u[i].FirstName)

		input := &dynamodb.PutItemInput

		input.Item = u[i]
		input.TableName = "user"

		svc.PutItem(input)

	}
}
