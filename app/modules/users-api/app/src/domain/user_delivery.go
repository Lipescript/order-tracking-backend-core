package domain

import "time"

type UserDelivery struct {
	UserID           string    `bson:"user_id" json:"user_id"`
	DeliveryID       string    `bson:"delivery_id" json:"delivery_id"`
	DeliveryStatus   string    `bson:"delivery_status" json:"delivery_status"`
	MessageTimestamp time.Time `bson:"message_timestamp" json:"message_timestamp"`
}

func NewUserDelivery(userID, deliveryID, deliveryStatus string, messageTimestamp time.Time) *UserDelivery {
	return &UserDelivery{
		UserID:           userID,
		DeliveryID:       deliveryID,
		DeliveryStatus:   deliveryStatus,
		MessageTimestamp: messageTimestamp,
	}
}
