package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func putDynamoItem(dynamoItem DynamoLinkItem, dynamoTableName string, dynamoClient *dynamodb.Client) {

	marshalledDynamoItem, err := attributevalue.MarshalMap(dynamoItem)

	if err != nil {
		panic(err)
	}

	out, putError := dynamoClient.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String(dynamoTableName),
		Item:      marshalledDynamoItem,
	})

	if putError != nil {
		panic(putError)
	}

	fmt.Println(out.Attributes)

}
