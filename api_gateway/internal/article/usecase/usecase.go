package usecase

import (
	"context"
	"log"
	"sharing_vasion_indonesia/api_gateway/internal/models"
	article_proto "sharing_vasion_indonesia/pkg/proto"
)

type articleUseCase struct {
	articleClient article_proto.ArticleServiceClient
}

func NewArticleUseCase(articleClient article_proto.ArticleServiceClient) *articleUseCase {
	return &articleUseCase{articleClient}
}

func (u *articleUseCase) CreateUser(ctx context.Context, payload models.CreateArticle) (*article_proto.Post, error) {
	req := &article_proto.CreateArticleRequest{
		Title:    payload.Title,
		Content:  payload.Content,
		Category: payload.Category,
		Status:   payload.Status,
	}

	result, err := u.articleClient.CreateArticle(ctx, req)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return result, nil
}
