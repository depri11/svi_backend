package delivery

import (
	"context"
	"log"
	"sharing_vasion_indonesia/article_service/internal/interfaces"
	article_proto "sharing_vasion_indonesia/pkg/proto"

	"github.com/golang/protobuf/ptypes/empty"
)

type delivery struct {
	ucArticle interfaces.ArticleUsecase
}

func NewDelivery(ucArticle interfaces.ArticleUsecase) *delivery {
	return &delivery{ucArticle}
}

func (d *delivery) GetArticles(ctx context.Context, req *article_proto.GetArticlesRequest) (*article_proto.GetArticlesResponse, error) {
	result, err := d.ucArticle.GetArticles(ctx, req)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return result, nil
}

func (d *delivery) CreateArticle(ctx context.Context, req *article_proto.CreateArticleRequest) (*article_proto.Post, error) {
	result, err := d.ucArticle.CreateArticle(ctx, req)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return result, nil
}

func (d *delivery) GetArticleById(ctx context.Context, req *article_proto.GetArticleByIdRequest) (*article_proto.GetArticleResponse, error) {
	result, err := d.ucArticle.GetArticleById(ctx, req.Id)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return result, nil
}

func (d *delivery) UpdateArticleById(ctx context.Context, req *article_proto.UpdateArticleRequest) (*article_proto.Post, error) {
	result, err := d.ucArticle.UpdateArticleById(ctx, req)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return result, nil
}

func (d *delivery) DeleteArticleById(ctx context.Context, req *article_proto.DeleteArticleRequest) (*empty.Empty, error) {
	err := d.ucArticle.DeleteArticleById(ctx, req.Id)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &empty.Empty{}, nil
}
