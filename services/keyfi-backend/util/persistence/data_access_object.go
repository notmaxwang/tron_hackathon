package persistence

import (
	"context"
	"keyfi-backend/util/persistence/models"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type DataAccessObject struct {
	Client    *dynamodb.Client
	tableName string
	region    string
}

func (dao *DataAccessObject) GetItem(walletAddress string) (*models.UserProfileModel, error) {
	dynamoDBKey := map[string]types.AttributeValue{
		"wallet_address": &types.AttributeValueMemberS{Value: walletAddress},
	}
	res, err := dao.Client.GetItem(context.TODO(), &dynamodb.GetItemInput{
		ConsistentRead: aws.Bool(true),
		TableName:      aws.String(dao.tableName),
		Key:            dynamoDBKey,
	})
	if err != nil {
		return nil, err
	}

	model := models.UserProfileModel{}
	model.Populate(&res.Item)
	return &model, nil
}

func (dao *DataAccessObject) PutItem(item *models.UserProfileModel) error {
	itemAttr := item.ToDaoItem()
	input := &dynamodb.PutItemInput{
		TableName: &dao.tableName,
		Item:      *itemAttr,
	}

	_, err := dao.Client.PutItem(context.TODO(), input)

	if err != nil {
		log.Printf("Failed to put item for %s in DynamoDB - %s\n", item.WalletAddress, err.Error())
	}
	return err
}
