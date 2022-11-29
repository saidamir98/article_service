package author

import (
	"context"
	"log"
	"uacademy/blogpost/article_service/protogen/blogpost"
)

type AuthorService struct {
	blogpost.UnimplementedAuthorServiceServer
}

// Ping ...
func (s *AuthorService) Ping(ctx context.Context, req *blogpost.Empty) (*blogpost.Pong, error) {
	log.Println("Ping")

	return &blogpost.Pong{
		Message: "OK",
	}, nil
}
