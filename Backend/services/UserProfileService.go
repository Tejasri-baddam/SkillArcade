package services

import (
	"BACKEND/models"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetUserProfileService(ctx context.Context, collection *mongo.Collection, userID string) (models.UserProfile, error) {

	userDetailsCollection := collection

	// Convert userID to ObjectID
	userObjectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return models.UserProfile{}, errors.New("invalid user ID format")
	}

	// Find the user document by user_id
	var userProfile models.UserProfile
	err = userDetailsCollection.FindOne(ctx, bson.M{"_id": userObjectID}).Decode(&userProfile)
	if err != nil {
		return models.UserProfile{}, errors.New("user not found")
	}

	return userProfile, nil
}
