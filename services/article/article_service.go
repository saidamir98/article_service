package article

import (
	"context"
	"fmt"
	"log"

	"uacademy/blogpost/article_service/models"
	blogpost "uacademy/blogpost/article_service/protogen/blogpost"
	"uacademy/blogpost/article_service/storage"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type articleService struct {
	stg storage.StorageI
	blogpost.UnimplementedArticleServiceServer
}

// NewArticleService ...
func NewArticleService(stg storage.StorageI) *articleService {
	return &articleService{
		stg: stg,
	}
}

// Ping ...
func (s *articleService) Ping(ctx context.Context, req *blogpost.Empty) (*blogpost.Pong, error) {
	log.Println("Ping")

	return &blogpost.Pong{
		Message: "OK",
	}, nil
}

// CreateArticle ...
func (s *articleService) CreateArticle(ctx context.Context, req *blogpost.CreateArticleRequest) (*blogpost.Article, error) {
	fmt.Println(req)

	id := uuid.New()

	err := s.stg.AddArticle(id.String(), models.CreateArticleModel{
		Content: models.Content{
			Title: req.Content.Title,
			Body:  req.Content.Body,
		},
		AuthorID: req.AuthorId,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.AddArticle: %s", err.Error())
	}

	article, err := s.stg.GetArticleByID(id.String())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.Stg.GetArticleByID: %s", err.Error())
	}

	var updatedAt string
	if article.UpdatedAt != nil {
		updatedAt = article.UpdatedAt.String()
	}

	return &blogpost.Article{
		Id: article.ID,
		Content: &blogpost.Content{
			Title: article.Title,
			Body:  article.Body,
		},
		AuthorId:  article.Author.ID,
		CreatedAt: article.CreatedAt.String(),
		UpdatedAt: updatedAt,
	}, nil
}

// UpdateArticle ....
func (s *articleService) UpdateArticle(ctx context.Context, req *blogpost.UpdateArticleRequest) (*blogpost.Article, error) {
	err := s.stg.UpdateArticle(models.UpdateArticleModel{
		ID: req.Id,
		Content: models.Content{
			Title: req.Content.Title,
			Body:  req.Content.Body,
		},
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.UpdateArticle: %s", err.Error())
	}

	article, err := s.stg.GetArticleByID(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.GetArticleByID: %s", err.Error())
	}

	var updatedAt string
	if article.UpdatedAt != nil {
		updatedAt = article.UpdatedAt.String()
	}

	return &blogpost.Article{
		Id: article.ID,
		Content: &blogpost.Content{
			Title: article.Title,
			Body:  article.Body,
		},
		AuthorId:  article.Author.ID,
		CreatedAt: article.CreatedAt.String(),
		UpdatedAt: updatedAt,
	}, nil
}

// DeleteArticle ....
func (s *articleService) DeleteArticle(ctx context.Context, req *blogpost.DeleteArticleRequest) (*blogpost.Article, error) {
	article, err := s.stg.GetArticleByID(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.GetArticleByID: %s", err.Error())
	}

	var updatedAt string
	if article.UpdatedAt != nil {
		updatedAt = article.UpdatedAt.String()
	}

	err = s.stg.DeleteArticle(article.ID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.DeleteArticle: %s", err.Error())
	}

	return &blogpost.Article{
		Id: article.ID,
		Content: &blogpost.Content{
			Title: article.Title,
			Body:  article.Body,
		},
		AuthorId:  article.Author.ID,
		CreatedAt: article.CreatedAt.String(),
		UpdatedAt: updatedAt,
	}, nil
}

// GetArticleList ....
func (s *articleService) GetArticleList(ctx context.Context, req *blogpost.GetArticleListRequest) (*blogpost.GetArticleListResponse, error) {
	res := &blogpost.GetArticleListResponse{
		Articles: make([]*blogpost.Article, 0),
	}

	articleList, err := s.stg.GetArticleList(int(req.Offset), int(req.Limit), req.Search)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.DeleteArticle: %s", err.Error())
	}

	for _, v := range articleList {
		var updatedAt string
		if v.UpdatedAt != nil {
			updatedAt = v.UpdatedAt.String()
		}

		res.Articles = append(res.Articles, &blogpost.Article{
			Id: v.ID,
			Content: &blogpost.Content{
				Title: v.Title,
				Body:  v.Body,
			},
			AuthorId:  v.AuthorID,
			CreatedAt: v.CreatedAt.String(),
			UpdatedAt: updatedAt,
		})
	}

	return res, nil
}

// GetArticleByID ....
func (s *articleService) GetArticleByID(ctx context.Context, req *blogpost.GetArticleByIDRequest) (*blogpost.GetArticleByIDResponse, error) {
	article, err := s.stg.GetArticleByID(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.GetArticleByID: %s", err.Error())
	}

	if article.DeletedAt != nil {
		return nil, status.Errorf(codes.NotFound, "s.stg.GetArticleByID: %s", err.Error())
	}

	var updatedAt string
	if article.UpdatedAt != nil {
		updatedAt = article.UpdatedAt.String()
	}

	var authorUpdatedAt string
	if article.Author.UpdatedAt != nil {
		authorUpdatedAt = article.Author.UpdatedAt.String()
	}

	return &blogpost.GetArticleByIDResponse{
		Id: article.ID,
		Content: &blogpost.Content{
			Title: article.Title,
			Body:  article.Body,
		},
		Author: &blogpost.GetArticleByIDResponse_Author{
			Id:        article.Author.ID,
			Fullname:  article.Author.Fullname,
			CreatedAt: article.CreatedAt.String(),
			UpdatedAt: authorUpdatedAt,
		},
		CreatedAt: article.CreatedAt.String(),
		UpdatedAt: updatedAt,
	}, nil
}
