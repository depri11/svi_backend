package interfaces

import (
	"context"
	article_proto "sharing_vasion_indonesia/pkg/proto"
)

type ArticleRepository interface {
	GetArticles(ctx context.Context, page int, limit int) (*article_proto.GetArticlesResponse, error)
	CreateArticle(ctx context.Context, post *article_proto.CreateArticleRequest) (*article_proto.Post, error)
	UpdateArticleById(ctx context.Context, post *article_proto.UpdateArticleRequest) (*article_proto.Post, error)
	GetArticleById(ctx context.Context, id string) (*article_proto.GetArticleResponse, error)
	DeleteArticleById(ctx context.Context, id string) error
}
