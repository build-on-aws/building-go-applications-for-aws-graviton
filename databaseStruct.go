package main

type DynamoLinkItem struct {
	ID          string `dynamodbav:"shortURL" json:"shortURL"`
	OriginalURL string `dynamodbav:"originalURL" json:"originalURL"`
}
