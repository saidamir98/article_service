package article

import (
	"context"
	"log"
	blogpost "uacademy/blogpost/article_service/protogen/blogpost"
)

type ArticleService struct {
	blogpost.UnimplementedArticleServiceServer
}

// Ping ...
func (s *ArticleService) Ping(ctx context.Context, req *blogpost.Empty) (*blogpost.Pong, error) {
	log.Println("Ping")

	return &blogpost.Pong{
		Message: "OK",
	}, nil
}
