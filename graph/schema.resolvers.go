package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"alshashiguchi/quiz_gem/graph/generated"
	"alshashiguchi/quiz_gem/graph/model"
	"alshashiguchi/quiz_gem/middleware/auth"
	"alshashiguchi/quiz_gem/models/users"
	"context"
	"fmt"

	sec "alshashiguchi/quiz_gem/core/security"

	log "github.com/sirupsen/logrus"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	log.Info("Start CreateUser")
	var user users.User
	var userModel model.User

	user.Name = input.Name
	user.Username = input.Username
	user.Email = input.Email
	user.Access = input.Access
	user.Situation = input.Situation
	user.Password = input.Password

	userModel, err := user.Create()

	if err != nil {
		return nil, err
	}

	log.Info("User Created ", userModel.ID)

	log.Info("End CreateUser")
	return &userModel, nil
}

func (r *mutationResolver) Login(ctx context.Context, input model.Login) (string, error) {
	var user users.User
	user.Username = input.Username
	user.Password = input.Password
	correct := user.Authenticate()
	if !correct {
		return "", &users.WrongUsernameOrPasswordError{}
	}
	token, err := sec.GenerateToken(user.Username)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (r *mutationResolver) RefreshToken(ctx context.Context, input model.RefreshTokenInput) (string, error) {
	username, err := sec.ParseToken(input.Token)
	if err != nil {
		return "", fmt.Errorf("access denied")
	}
	token, err := sec.GenerateToken(username)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	log.Info("Start GetAll Users")

	var access []model.Access
	access = append(access, model.AccessInstructor)
	userCtx := auth.ForContext(ctx, access)
	if userCtx == nil {
		return nil, fmt.Errorf("access denied")
	}

	var result []*model.User

	dbUsers, err := users.GetAll()
	if err != nil {
		return nil, err
	}

	for _, user := range dbUsers {
		result = append(result, &model.User{ID: user.ID, Username: user.Username, Name: user.Name, Email: user.Email, Access: user.Access, Situation: user.Situation})
	}
	log.Info("End GetAll Users")
	return result, nil
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
