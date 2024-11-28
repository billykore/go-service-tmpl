package repo

import (
	"context"
	"errors"

	"github.com/billykore/go-service-tmpl/domain/greet"
)

type GreetRepo struct {
	names []string
}

func NewGreetRepo() *GreetRepo {
	return &GreetRepo{}
}

func (r *GreetRepo) GetAll(_ context.Context) ([]greet.Message, error) {
	if len(r.names) == 0 {
		return nil, greet.ErrEmptyMessages
	}
	var greets []greet.Message
	for _, name := range r.names {
		greets = append(greets, greet.Message{Name: name})
	}
	return greets, nil
}

func (r *GreetRepo) Save(_ context.Context, message greet.Message) error {
	if message.Name == "" {
		return errors.New("message name is empty")
	}
	r.names = append(r.names, message.Name)
	return nil
}
