package repository

import (
	"context"
	"database/sql"
	"log"
	article_proto "sharing_vasion_indonesia/pkg/proto"
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
	query := `
	INSERT INTO posts
  (title, content, category, created_date, updated_date, status)
  VALUES (?, ?, ?, ?, ?, ?);
  `
	now := time.Now()

	result, err := r.db.ExecContext(ctx, query, post.Title, post.Content, post.Category, now, now, post.Status)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	idResult, err := result.LastInsertId()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	timestampProto := timestamppb.New(now)

	return &article_proto.Post{
		Id:          int32(idResult),
		Title:       post.Title,
		Content:     post.Content,
		Category:    post.Category,
		Status:      post.Status,
		CreatedDate: timestampProto,
		UpdatedDate: timestampProto,
	}, nil
}
