package usecase

import (
	"context"
	"log"
	"sharing_vasion_indonesia/article_service/internal/interfaces"
	article_proto "sharing_vasion_indonesia/pkg/proto"
)

type usecase struct {
	repoArticle interfaces.ArticleRepository
}

func NewUseCase(repoArticle interfaces.ArticleRepository) *usecase {
	return &usecase{repoArticle}
}

func (uc *usecase) CreateArticle(ctx context.Context, post *article_proto.CreateArticleRequest) (*article_proto.Post, error) {
	result, err := uc.repoArticle.CreateArticle(ctx, post)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return result, nil
}
