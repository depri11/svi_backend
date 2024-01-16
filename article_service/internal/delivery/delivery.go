package delivery

import (
	"context"
	"log"
	"sharing_vasion_indonesia/article_service/internal/interfaces"
	article_proto "sharing_vasion_indonesia/pkg/proto"
)

type delivery struct {
	ucArticle interfaces.ArticleUsecase
}

func NewDelivery(ucArticle interfaces.ArticleUsecase) *delivery {
	return &delivery{ucArticle}
}

func (d *delivery) CreateArticle(ctx context.Context, req *article_proto.CreateArticleRequest) (*article_proto.Post, error) {

	result, err := d.ucArticle.CreateArticle(ctx, req)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return result, nil

}
