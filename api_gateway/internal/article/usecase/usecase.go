package usecase

import (
	"context"
	"log"
	"sharing_vasion_indonesia/api_gateway/internal/models"
	article_proto "sharing_vasion_indonesia/pkg/proto"
	"strconv"
)

type articleUseCase struct {
	articleClient article_proto.ArticleServiceClient
}

func NewArticleUseCase(articleClient article_proto.ArticleServiceClient) *articleUseCase {
	return &articleUseCase{articleClient}
}

func (u *articleUseCase) GetArticles(ctx context.Context, limit string, offset string) (*article_proto.GetArticlesResponse, error) {

	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	offsetInt, err := strconv.Atoi(offset)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	result, err := u.articleClient.GetArticles(ctx, &article_proto.GetArticlesRequest{Limit: int32(limitInt), Offset: int32(offsetInt)})
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return result, nil
}

func (u *articleUseCase) CreateArticle(ctx context.Context, payload models.CreateArticleRequest) (*article_proto.Post, error) {

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

func (u *articleUseCase) UpdateArticle(ctx context.Context, payload models.UpdateArticleRequest) (*article_proto.Post, error) {
	id, err := strconv.Atoi(payload.ID)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	req := &article_proto.UpdateArticleRequest{
		Id:       int32(id),
		Title:    payload.Title,
		Content:  payload.Content,
		Category: payload.Category,
		Status:   payload.Status,
	}

	result, err := u.articleClient.UpdateArticleById(ctx, req)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return result, nil
}

func (u *articleUseCase) GetArticle(ctx context.Context, id string) (*article_proto.GetArticleResponse, error) {
	result, err := u.articleClient.GetArticleById(ctx, &article_proto.GetArticleByIdRequest{Id: id})
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return result, nil
}

func (u *articleUseCase) DeleteArticle(ctx context.Context, id string) error {
	_, err := u.articleClient.DeleteArticleById(ctx, &article_proto.DeleteArticleRequest{Id: id})
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
