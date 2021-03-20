package go_mysql

import (
	"entity"
	"fmt"
	"testing"
)

func RepositoryTest(t *testing.T) {
	CommentRepository := NewCommentRepository(OpenConnection())

	result, err := CommentRepository.InsertComment(entity.Comment{
		Id:      nil,
		Email:   "mochjuangpajri@gmail.com",
		Comment: "test repository",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(result)

}
