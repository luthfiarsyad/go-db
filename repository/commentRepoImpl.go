package repository

import (
	"belajar-golang-database/entity"
	"context"
	"database/sql"
)

type commentRepositoryImpl struct {
	db *sql.DB
}

func NewCommentRepository(db *sql.DB) *commentRepositoryImpl {
	return &commentRepositoryImpl{db}
}

func (repo *commentRepositoryImpl) Insert(ctx context.Context, comment entity.Comment) (entity.Comment, error) {
	query := "INSERT INTO comment(email,comment) VALUES(?,?)"
	res, err := repo.db.ExecContext(ctx, query, comment.Email, comment.Comment)
	if err != nil {
		return comment, err
	}
	lastInsertID, _ := res.LastInsertId()
	comment.Id = int(lastInsertID)
	return comment, nil
}

func (repo *commentRepositoryImpl) FindById(ctx context.Context, id int) (entity.Comment, error) {
	var comment entity.Comment
	query := "SELECT id,email,comment FROM comment WHERE id = ?"
	row := repo.db.QueryRowContext(ctx, query, id)
	err := row.Scan(&comment.Id, &comment.Email, &comment.Comment)
	if err != nil {
		return comment, err
	}
	return comment, nil
}
func (repo *commentRepositoryImpl) FindAll(ctx context.Context) ([]entity.Comment, error) {
	var comments []entity.Comment
	query := "SELECT id,email,comment FROM comment"
	rows, err := repo.db.QueryContext(ctx, query)
	if err != nil {
		return comments, err
	}
	defer rows.Close()
	for rows.Next() {
		var comment entity.Comment
		rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		comments = append(comments, comment)
	}

	return comments, nil
}
