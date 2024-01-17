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

func (uc *usecase) GetArticles(ctx context.Context, payload *article_proto.GetArticlesRequest) (*article_proto.GetArticlesResponse, error) {
	result, err := uc.repoArticle.GetArticles(ctx, int(payload.Limit), int(payload.Offset))
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return result, nil
}

func (uc *usecase) CreateArticle(ctx context.Context, article *article_proto.CreateArticleRequest) (*article_proto.Post, error) {
	result, err := uc.repoArticle.CreateArticle(ctx, article)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return result, nil
}

func (uc *usecase) GetArticleById(ctx context.Context, id string) (*article_proto.GetArticleResponse, error) {
	result, err := uc.repoArticle.GetArticleById(ctx, id)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return result, nil
}

func (uc *usecase) UpdateArticleById(ctx context.Context, article *article_proto.UpdateArticleRequest) (*article_proto.Post, error) {
	result, err := uc.repoArticle.UpdateArticleById(ctx, article)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return result, nil
}

func (uc *usecase) DeleteArticleById(ctx context.Context, id string) error {
	err := uc.repoArticle.DeleteArticleById(ctx, id)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
