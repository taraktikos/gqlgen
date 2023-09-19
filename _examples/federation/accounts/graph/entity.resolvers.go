package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.38

import (
	"context"

	"github.com/99designs/gqlgen/_examples/federation/accounts/graph/model"
)

// FindEmailHostByID is the resolver for the findEmailHostByID field.
func (r *entityResolver) FindEmailHostByID(ctx context.Context, id string) (*model.EmailHost, error) {
	return &model.EmailHost{
		ID:   id,
		Name: "Email Host " + id,
	}, nil
}

// FindUserByID is the resolver for the findUserByID field.
func (r *entityResolver) FindUserByID(ctx context.Context, id string) (*model.User, error) {
	name := "User " + id
	if id == "1234" {
		name = "Me"
	}

	return &model.User{
		ID:       id,
		Username: name,
	}, nil
}

// Entity returns EntityResolver implementation.
func (r *Resolver) Entity() EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }
