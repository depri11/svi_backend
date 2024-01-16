package interfaces

import (
	"context"
	article_proto "sharing_vasion_indonesia/pkg/proto"
)

type ArticleUsecase interface {
	CreateArticle(ctx context.Context, post *article_proto.CreateArticleRequest) (*article_proto.Post, error)
}
