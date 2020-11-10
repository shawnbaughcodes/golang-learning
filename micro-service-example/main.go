package main

import (
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"errors"
	"fmt"
)

func main()  {
	tableName := "fastify-practice-users"
	name := "name"

	serviceSession := session.Must(session.NewSessionWithOptions(session.Options {
    	SharedConfigState: session.SharedConfigEnable,
	}))
	service := dynamodb.New(serviceSession)

	result, err := service.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"name": {
				S: aws.String(name),
			},
		},
	})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if result.Item == nil {
		msg := "Could not find '" + name + "'"
		fmt.Println(errors.New(msg))
		// return nil, errors.New(msg)
	}
	var currentUser = User{}
	
	err = dynamodbattribute.UnmarshalMap(result.Item, &currentUser)
	if err != nil {
		panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
	}

	fmt.Println("Found item:")
	fmt.Println("Name:  ", currentUser.name)
}