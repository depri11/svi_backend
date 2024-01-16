package repository

import (
	"context"
	"database/sql"
	"log"
	article_proto "sharing_vasion_indonesia/pkg/proto"
	"strconv"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateArticle(ctx context.Context, post *article_proto.CreateArticleRequest) (*article_proto.Post, error) {
	var id string

	query := `
  INSERT INTO article.posts
  (Title, Content, Category, Created_date, Updated_date, Status)
  VALUES($1, $2, $3, $4, $5, $6, $7) returning id
  `
	now := time.Now()

	err := r.db.QueryRowContext(ctx, query, post.Title, post.Content, post.Category, now, now, post.Status).Scan(&id)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	timestampProto := timestamppb.New(now)

	return &article_proto.Post{
		Id:          int32(idInt),
		Title:       post.Title,
		Content:     post.Content,
		Category:    post.Category,
		Status:      post.Status,
		CreatedDate: timestampProto,
		UpdatedDate: timestampProto,
	}, nil
}
