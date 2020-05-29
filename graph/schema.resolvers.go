package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"alshashiguchi/quiz_gem/graph/generated"
	"alshashiguchi/quiz_gem/graph/model"
	"log"
	"strconv"

	"alshashiguchi/quiz_gem/models/users"
	"context"
	"fmt"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	var user users.User
	var userModel model.User

	user.Name = input.Name
	user.Username = input.Username
	user.Email = input.Email
	user.Access = input.Access
	user.Situation = input.Situation

	idCreated, err := user.Create()
	log.Println("User Created")

	if err != nil {
		return nil, err
	}

	userModel = user.ConvertToGraphModelUser()
	userModel.ID = strconv.Itoa(int(idCreated))

	return &userModel, nil
}

func (r *mutationResolver) Login(ctx context.Context, input model.Login) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) RefreshToken(ctx context.Context, input model.RefreshTokenInput) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	var users []*model.User
	users = append(users, &model.User{Name: "Andre", Access: model.AccessAdmin, Situation: model.UserStatusActive})

	return users, nil
}

func (r *queryResolver) User(ctx context.Context, id *int) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
