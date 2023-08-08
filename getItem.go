package main

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func getDynamoItem(shortUrlId string, dynamoTableName string, dynamoClient *dynamodb.Client) DynamoLinkItem {

	dynamoKey := struct {
		ID string `dynamodbav:"shortURL" json:"shortURL"`
	}{ID: shortUrlId}

	marshalledDynamoKey, err := attributevalue.MarshalMap(dynamoKey)

	if err != nil {
		panic(err)
	}

	result, err := dynamoClient.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String(dynamoTableName),
		Key:       marshalledDynamoKey,
	})

	if err != nil {
		panic(err)
	}
	var fullDynamoItem DynamoLinkItem
	err = attributevalue.UnmarshalMap(result.Item, &fullDynamoItem)

	if err != nil {
		panic(err)
	}

	return fullDynamoItem
}
