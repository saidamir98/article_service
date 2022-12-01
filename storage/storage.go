package storage

import (
	"uacademy/blogpost/article_service/protogen/blogpost"
)

// StorageI ...
type StorageI interface {
	AddArticle(id string, entity *blogpost.CreateArticleRequest) error
	GetArticleByID(id string) (*blogpost.GetArticleByIDResponse, error)
	GetArticleList(offset, limit int, search string) (resp *blogpost.GetArticleListResponse, err error)
	UpdateArticle(entity *blogpost.UpdateArticleRequest) error
	DeleteArticle(id string) error

	GetAuthorByID(id string) (*blogpost.Author, error)
}
