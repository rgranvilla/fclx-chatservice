package gateway

import (
	"context"

	"github.com/rgranvilla/fclx/chatservice/internal/domain/entity"
)

type ChatGateway interface {
	CreateChat(ctx context.Context, chat *entity.Chat) error
	FindChatById(ctx context.Context, chatId string) (*entity.Chat, error)
	SaveChat(ctx context.Context, chat *entity.Chat) error
}
