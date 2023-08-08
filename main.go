package main

import (
	"context"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/gin-gonic/gin"
)

func getShortURL(dynamoClient *dynamodb.Client, dynamoTableName string) func(*gin.Context) {
	return func(c *gin.Context) {
		var urlToShorten DynamoLinkItem

		if err := c.BindJSON(&urlToShorten); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		urlToShorten.ID = generateShortId()
		putDynamoItem(urlToShorten, dynamoTableName, dynamoClient)

		c.IndentedJSON(http.StatusCreated, urlToShorten)
	}
}

func getFullURL(dynamoClient *dynamodb.Client, dynamoTableName string) func(*gin.Context) {
	return func(c *gin.Context) {
		shortUrlId := c.Param("id")
		fullUrlItem := getDynamoItem(shortUrlId, dynamoTableName, dynamoClient)

		c.IndentedJSON(http.StatusOK, fullUrlItem.OriginalURL)
	}
}

func getDynamodbClient(awsRegion string) *dynamodb.Client {
	//can ommit second arg if AWS_REGION env var is defined
	cfg, configError := config.LoadDefaultConfig(context.TODO(), func(o *config.LoadOptions) error {
		o.Region = awsRegion
		return nil
	})

	if configError != nil {
		panic(configError)
	}

	return dynamodb.NewFromConfig(cfg)

}

func main() {
	// only set up dynamo client once
	awsRegion := "us-west-2"
	dynamoClient := getDynamodbClient(awsRegion)
	dynamoTableName := "goUrlShortener"

	router := gin.Default()
	router.POST("/shortenURL", getShortURL(dynamoClient, dynamoTableName))
	router.GET("/getFullURL/:id", getFullURL(dynamoClient, dynamoTableName))

	router.Run("0.0.0.0:8080")
}
