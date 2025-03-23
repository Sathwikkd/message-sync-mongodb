package mongodb

import (
	"chat-sync-service/internal/domain"
	"context"
	"errors"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type messageRepo struct {
	db *mongo.Database
}

// Constructor: db is injected, collection will be chosen per user dynamically
func NewMessageRepository(db *mongo.Database) domain.MessageRepository {
	return &messageRepo{
		db: db,
	}
}

func (r *messageRepo) GetMessagesAfter(userID, currentMsgID string) ([]domain.Message, error) {
	// Use userID as collection name
	collection := r.db.Collection(userID)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	//  Fetch all messages for the user
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var allMessages []domain.Message
	if err := cursor.All(ctx, &allMessages); err != nil {
		return nil, err
	}

	//  Build map[prevMessageID]Message
	messageMap := make(map[string]domain.Message)
	for _, msg := range allMessages {
		messageMap[msg.PrevMessageId] = msg
	}

	// Order messages after currentMsgID
	var orderedMessages []domain.Message
	currentID := currentMsgID

	for {
		nextMsg, ok := messageMap[currentID]
		if !ok {
			break // No more messages in chain
		}
		orderedMessages = append(orderedMessages, nextMsg)
		currentID = nextMsg.MessageID
	}

	if len(orderedMessages) == 0 {
		return nil, errors.New("no messages found after the given message_id")
	}

	log.Printf("Synced %d messages for user: %s", len(orderedMessages), userID)
	return orderedMessages, nil
}
