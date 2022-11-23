package controller

import (
	"belajar-golang-database/entity"
	"belajar-golang-database/repository"
	"context"
)

func CreateComment(repo repository.CommentRepository, email, comment string) (entity.Comment, error) {
	ctx := context.Background()
	commentStruct := entity.Comment{
		Email:   email,
		Comment: comment,
	}
	commentStruct, err := repo.Insert(ctx, commentStruct)
	return commentStruct, err
}

func GetComment(repo repository.CommentRepository) ([]entity.Comment, error) {
	ctx := context.Background()
	comments, err := repo.FindAll(ctx)
	return comments, err
}

func GetCommentById(repo repository.CommentRepository, id int) (entity.Comment, error) {
	ctx := context.Background()
	return repo.FindById(ctx, id)
}
