package user

import (
	"context"

	"github.com/davidk81/svelte-golang-demo/backend/patientdb"
	"github.com/davidk81/svelte-golang-demo/backend/patientdb/models"
	"github.com/volatiletech/sqlboiler/boil"
)

// handles database operations for user table

// getUser fetches user by userid
func getUser(ctx context.Context, userid string) (*models.User, error) {
	return models.Users(models.UserWhere.Userid.EQ(userid)).One(ctx, patientdb.DB())
}

// createUser create new user
func createUser(ctx context.Context, user *models.User) error {
	return (*user).Insert(ctx, patientdb.DB(), boil.Infer())
}
