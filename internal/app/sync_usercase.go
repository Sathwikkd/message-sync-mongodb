package app

import (
	"chat-sync-service/internal/domain"
)

type SyncUsecase struct {
	repo domain.MessageRepository
}

func NewSyncUsecase(r domain.MessageRepository) *SyncUsecase {
	return &SyncUsecase{
		repo: r,
	}
}

func (u *SyncUsecase) SyncMessages(userID, currentMsgID string) ([]domain.Message, error) {
	return u.repo.GetMessagesAfter(userID, currentMsgID)
}
