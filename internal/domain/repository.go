// repository.go
package domain

type MessageRepository interface {
	GetMessagesAfter(userID, currentMsgID string) ([]Message, error)
}
