package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func main() {

	svc := dynamodb.New(session.New(), &aws.Config{Endpoint: aws.String("http://localhost:8000"), Region: aws.String("us-east-1")})

	params := &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{ // Required
			{ // Required
				AttributeName: aws.String("id"), // Required
				AttributeType: aws.String("S"),        // Required
			},
			{ // Required
				AttributeName: aws.String("user"), // Required
				AttributeType: aws.String("S"),     // Required
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{ // Required
			{ // Required
				AttributeName: aws.String("id"), // Required
				KeyType:       aws.String("HASH"),     // Required

			}, {
				AttributeName: aws.String("user"), // Required
				KeyType:       aws.String("RANGE"), // Required
			},

		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{ // Required
			ReadCapacityUnits:  aws.Int64(1), // Required
			WriteCapacityUnits: aws.Int64(1), // Required
		},
		TableName: aws.String("workout"), // Required
		GlobalSecondaryIndexes: []*dynamodb.GlobalSecondaryIndex{
			{ // Required
				IndexName: aws.String("user"), // Required
				KeySchema: []*dynamodb.KeySchemaElement{ // Required
					{ // Required
						AttributeName: aws.String("user"), // Required
						KeyType:       aws.String("HASH"),  // Required
					},
				},
				Projection: &dynamodb.Projection{ // Required
					NonKeyAttributes: []*string{
						aws.String("user"),
                                                aws.String("id"),
                                                aws.String("equipment"),
                                                aws.String("gym"),

					},
					ProjectionType: aws.String("INCLUDE"),
				},
				ProvisionedThroughput: &dynamodb.ProvisionedThroughput{ // Required
					ReadCapacityUnits:  aws.Int64(1), // Required
					WriteCapacityUnits: aws.Int64(1), // Required
				},
			},
			// More values...
		},
	}
	resp, err := svc.CreateTable(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}
