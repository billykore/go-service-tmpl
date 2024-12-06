package greet

import (
	"context"
	"errors"

	"github.com/billykore/go-service-tmpl/pkg/codes"
	"github.com/billykore/go-service-tmpl/pkg/logger"
	"github.com/billykore/go-service-tmpl/pkg/status"
)

// ErrEmptyGreets indicates no messages was stored in the repository.
var ErrEmptyGreets = errors.New("empty messages")

// Repository defines the methods to interacting with persistence storage used by greet domain.
type Repository interface {
	// GetAll gets all messages.
	GetAll(ctx context.Context) ([]Greet, error)

	// Save saves message.
	Save(ctx context.Context, message Greet) error
}

type Service struct {
	log  *logger.Logger
	repo Repository
}

func NewService(log *logger.Logger, repo Repository) *Service {
	return &Service{
		log:  log,
		repo: repo,
	}
}

func (s *Service) History(ctx context.Context) ([]*Response, error) {
	messages, err := s.repo.GetAll(ctx)
	if err != nil {
		s.log.Usecase("History").Errorf("Failed to get all messages: %v", err)
		return nil, status.Error(codes.NotFound, messageGetHistoryFailed)
	}
	resp := make([]*Response, len(messages))
	for i, message := range messages {
		resp[i] = &Response{
			Name: message.Name,
		}
	}
	s.log.Usecase("History").Info("Get all messages successful")
	return resp, nil
}

func (s *Service) SayHello(ctx context.Context, req HelloRequest) (*HelloResponse, error) {
	err := s.repo.Save(ctx, Greet{Name: req.Name})
	if err != nil {
		s.log.Usecase("SayHello").Errorf("Save message failed: %v", err)
		return nil, status.Error(codes.Internal, messageSayHelloFailed)
	}
	s.log.Usecase("SayHello").Infof("Say hello to %s", req.Name)
	return &HelloResponse{Message: "Hello " + req.Name}, nil
}
