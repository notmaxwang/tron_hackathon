package persistence

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

const mainTableName string = "main_table"
const region string = "us-east-1"

var dbClient *dynamodb.Client
var MainTableDao *DataAccessObject

func GetDynamoDBClient() (*dynamodb.Client, error) {
	if dbClient == nil {
		cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
		if err != nil {
			log.Fatalf("unable to load SDK config, %v", err)
			return nil, err
		}
		// Using the Config value, create the DynamoDB client
		dbClient = dynamodb.NewFromConfig(cfg)
	}
	return dbClient, nil
}

func GetMainTableDao() (*DataAccessObject, error) {
	client, err := GetDynamoDBClient()
	if err != nil {
		return nil, err
	}

	return &DataAccessObject{
		Client:  client,
		tableName: mainTableName,
		region:    region,
	}, nil
}
