package database

import (
	"context"

	"github.com/dev-hyunsang/clone-twitter-backend/ent"
	"github.com/dev-hyunsang/clone-twitter-backend/ent/user"
	"github.com/dev-hyunsang/clone-twitter-backend/models"
)

func CreateUser(userData models.User) error {
	client, err := ConnectMySQL()
	if err != nil {
		return err
	}

	err = client.User.Create().
		SetUserUUID(userData.UserUUID).
		SetUserEmail(userData.UserEmail).
		SetUserPassword(userData.UserPassword).
		SetUserPhoneNumber(userData.UserPhoneNumber).
		SetUserBirthday(userData.UserBirthday).
		SetCreatedAt(userData.CreatedAt).
		SetUpdatedAt(userData.UpdatedAt).
		Exec(context.Background())

	return err
}

func QueryUser(email string) (*ent.User, error) {
	client, err := ConnectMySQL()
	if err != nil {
		return nil, err
	}

	userData, err := client.User.Query().
		Where(user.UserEmail(email)).
		Only(context.Background())
	if err != nil {
		return nil, err
	}

	return userData, nil
}
