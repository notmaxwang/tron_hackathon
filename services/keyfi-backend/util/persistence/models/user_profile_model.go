package models

import (
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type UserProfileModel struct {
	WalletAddress    string
	SignedMessage    string
	SignedMessageTTL int64
	FirstName        string
	LastName         string
	MiddleName       string
	Role             string
	ChatS3Object     string
	CreationTime     int64
}

func (model *UserProfileModel) ToDaoItem() *map[string]types.AttributeValue {
	if model.CreationTime == 0 {
		model.CreationTime = time.Now().UTC().Unix()
	}

	return &map[string]types.AttributeValue{
		"wallet_address":     &types.AttributeValueMemberS{Value: model.WalletAddress},
		"signed_message":     &types.AttributeValueMemberS{Value: model.SignedMessage},
		"signed_message_ttl": &types.AttributeValueMemberN{Value: strconv.FormatInt(model.SignedMessageTTL, 10)},
		"first_name":         &types.AttributeValueMemberS{Value: model.FirstName},
		"last_name":          &types.AttributeValueMemberS{Value: model.LastName},
		"middle_name":        &types.AttributeValueMemberS{Value: model.MiddleName},
		"role":               &types.AttributeValueMemberS{Value: model.MiddleName},
		"chat_s3_object":     &types.AttributeValueMemberS{Value: model.ChatS3Object},
		"creation_time":      &types.AttributeValueMemberN{Value: strconv.FormatInt(model.CreationTime, 10)},
	}
}

func (model *UserProfileModel) Key() *map[string]types.AttributeValue {
	return &map[string]types.AttributeValue{
		"wallet_address": &types.AttributeValueMemberS{Value: model.WalletAddress},
	}
}

func (model *UserProfileModel) Populate(item *map[string]types.AttributeValue) {
	// Check if item is nil
	if item == nil {
		return
	}

	// Dereference the pointer to the map and extract values
	attributes := *item

	if val, ok := attributes["wallet_address"].(*types.AttributeValueMemberS); ok {
		model.WalletAddress = val.Value
	}
	if val, ok := attributes["signed_message"].(*types.AttributeValueMemberS); ok {
		model.SignedMessage = val.Value
	}
	if val, ok := attributes["signed_message_ttl"].(*types.AttributeValueMemberN); ok {
		if ttl, err := strconv.ParseInt(val.Value, 10, 64); err == nil {
			model.SignedMessageTTL = ttl
		}
	}
	if val, ok := attributes["first_name"].(*types.AttributeValueMemberS); ok {
		model.FirstName = val.Value
	}
	if val, ok := attributes["last_name"].(*types.AttributeValueMemberS); ok {
		model.LastName = val.Value
	}
	if val, ok := attributes["middle_name"].(*types.AttributeValueMemberS); ok {
		model.MiddleName = val.Value
	}
	if val, ok := attributes["role"].(*types.AttributeValueMemberS); ok {
		model.Role = val.Value
	}
	if val, ok := attributes["chat_s3_object"].(*types.AttributeValueMemberS); ok {
		model.ChatS3Object = val.Value
	}
	if val, ok := attributes["creation_time"].(*types.AttributeValueMemberN); ok {
		if creationTime, err := strconv.ParseInt(val.Value, 10, 64); err == nil {
			model.CreationTime = creationTime
		}
	}
}
