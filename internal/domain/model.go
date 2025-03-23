package domain

type Message struct {
	MessageID string `bson:"message_id" json:"messageid"`
	//UserID    string `bson:"user_id" json:"user_id"`
	MessageContent string `bson:"message_content" json:"messagecontent"`
	PrevMessageId  string `bson:"prev_message_id" json:"prevmessageid"`
	Timestamp      int64  `bson:"timestamp" json:"timestamp"`
	MessageFrom    string `bson:"message_from" json:"messageform"`
	MessageTo      string `bson:"message_to" json:"messageto"`
	MessageStatus  string `bson:"message_status" json:"messagestatus"`
	CreatedAt      string `bson:"created_at" json:"createdat"`
	DeliveredAt    string `bson:"delivered_at" json:"deliveredat"`
	ReadAt         string `bson:"read_at" json:"readat"`
	
}
