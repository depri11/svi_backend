package interfaces

import (
	"context"
	"sharing_vasion_indonesia/api_gateway/internal/models"
	article_proto "sharing_vasion_indonesia/pkg/proto"
)

type ArticleUseCase interface {
	CreateUser(ctx context.Context, payload models.CreateArticle) (*article_proto.Post, error)
}
