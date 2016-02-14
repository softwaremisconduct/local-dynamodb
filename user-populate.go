package user

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"golang.org/x/crypto/bcrypt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type User struct {
	ID          int    `json:"id"`
	Gender      string `json:"gender"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Weight      int    `json:"weight"`
	Height      int    `json:"height"`
	Username    string `json:"username"`
	Password    string `json:"password,omitempty"`
	EncPassword string `json:"enc_password"`
}

func clear(b []byte) {
	for i := 0; i < len(b); i++ {
		b[i] = 0
	}
}

func Crypt(password []byte) ([]byte, error) {
	defer clear(password)
	return bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
}

func Verify(hash []byte, password []byte) error {
	return bcrypt.CompareHashAndPassword(hash, password)
}

func main() {

	svc := dynamodb.New(session.New(), &aws.Config{Endpoint: aws.String("http://localhost:8000"), Region: aws.String("us-east-1")})

	file, e := ioutil.ReadFile("./mockdata/MOCK_DATA-user-2.json")
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}

	var u []User
	json.Unmarshal(file, &u)

	for i := 0; i < len(u); i++ {
		fmt.Printf(" First Name: %v\n", u[i].FirstName)

		encPass, cryptErr := Crypt([]byte(u[i].Password))
		if cryptErr != nil {
			fmt.Println("Failed to hash password", cryptErr)
			return
		}
		u[i].EncPassword = string(encPass)
		u[i].Password = ""

		item, err := dynamodbattribute.ConvertToMap(u[i])
		if err != nil {
			fmt.Println("Failed to convert", err)
			return
		}
		fmt.Printf("Item %v\n", item)

		params := &dynamodb.PutItemInput{
			Item:      item,
			TableName: aws.String("user"),
		}
		resp, err := svc.PutItem(params)

		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println(resp)

	}
}
