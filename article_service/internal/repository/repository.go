package repository

import (
	"context"
	"database/sql"
	"errors"
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

func (r *repository) GetArticles(ctx context.Context, limit int, offset int) (*article_proto.GetArticlesResponse, error) {
	var results article_proto.GetArticlesResponse
	var count int32

	countQuery := `
	SELECT count(1)
	FROM posts
	`

	query := `
	SELECT id, title, content, category, status
	FROM posts
	ORDER BY id
	LIMIT ? OFFSET ?;
  `

	rows, err := r.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	err = r.db.QueryRowContext(ctx, countQuery).Scan(&count)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	results.Total = count

	for rows.Next() {
		var result article_proto.GetArticleResponse
		err := rows.Scan(&result.Id, &result.Title, &result.Content, &result.Category, &result.Status)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		results.Data = append(results.Data, &result)
	}

	if err := rows.Err(); err != nil {
		log.Println(err)
		return nil, err
	}

	return &results, nil
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

func (r *repository) UpdateArticleById(ctx context.Context, post *article_proto.UpdateArticleRequest) (*article_proto.Post, error) {
	query := `
		UPDATE posts
		SET title = ?, content = ?, category = ?, updated_date = ?, status = ?
		WHERE id = ?;
	`

	now := time.Now()

	result, err := r.db.ExecContext(ctx, query, post.Title, post.Content, post.Category, now, post.Status, post.Id)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if rowsAffected == 0 {
		return nil, errors.New("no rows affected during update operation")
	}

	timestampProto := timestamppb.New(now)

	return &article_proto.Post{
		Id:          post.Id,
		Title:       post.Title,
		Content:     post.Content,
		Category:    post.Category,
		Status:      post.Status,
		CreatedDate: timestampProto,
		UpdatedDate: timestampProto,
	}, nil
}

func (r *repository) GetArticleById(ctx context.Context, id string) (*article_proto.GetArticleResponse, error) {
	var result article_proto.GetArticleResponse

	query := `
	SELECT title, content, category, status
	FROM posts
	WHERE id = ?;
  `

	err := r.db.QueryRowContext(ctx, query, id).Scan(&result.Title, &result.Content, &result.Category, &result.Status)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &result, nil
}

func (r *repository) DeleteArticleById(ctx context.Context, id string) error {
	query := `
	DELETE FROM posts
	WHERE id = ?;
  `

	res, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		log.Println(err)
		return err
	}

	rowsNum, err := res.RowsAffected()
	if err != nil {
		log.Println(err)
		return err
	}

	if rowsNum == 0 {
		return errors.New("no rows affected during delete operation")
	}

	return nil
}
