package go_mysql

import (
	"context"
	"database/sql"
	"errors"
	"go_mysql/entity"
	"go_mysql/repository"
	"strconv"
)

type CommentImpl struct {
	DB *sql.DB
}

func (repository *CommentImpl) InsertComment(ctx context.Context, comment entity.Comment) (entity.Comment, error) {
	stmt, _ := repository.DB.PrepareContext(ctx, "INSERT comment values('', ?, ?)")
	res, err := stmt.ExecContext(ctx, comment.Email, comment.Comment)
	stmt.Close()
	if err != nil {
		return comment, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return comment, err
	}
	comment.Id = int64(id)
	return comment, err
}
func (repository *CommentImpl) FindById(ctx context.Context, id int32) (entity.Comment, error) {
	stmt, _ := repository.DB.PrepareContext(ctx, "select * from comment where id = ?")
	rows, err := stmt.QueryContext(ctx, id)
	comment := entity.Comment{}
	defer rows.Close()
	if err != nil {
		return comment, err
	}
	if rows.Next() {
		rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		return comment, nil

	} else {
		return comment, errors.New("id " + strconv.Itoa(int(id)) + " not found")
	}

}
func (repository *CommentImpl) FindAll(ctx context.Context) ([]entity.Comment, error) {
	stmt, _ := repository.DB.PrepareContext(ctx, "select * from comment")
	rows, err := stmt.QueryContext(ctx)
	comments := []entity.Comment{}
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		comment := entity.Comment{}
		rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		comments = append(comment)

	}
	return comments, nil
}

func NewCommentRepository(db *sql.DB) repository.CommentRepository {
	return &repository.CommentRepository{}
}
