package user

import (
	"github.com/davidk81/svelte-golang-demo/backend/patientdb"
	"github.com/davidk81/svelte-golang-demo/backend/patientdb/models"
	"github.com/valyala/fasthttp"
	"github.com/volatiletech/sqlboiler/boil"
)

// handles database operations for user table

// getUser fetches user by userid
func getUser(userid string, ctx *fasthttp.RequestCtx) (*models.User, error) {
	return models.Users(models.UserWhere.Userid.EQ(userid)).One(ctx, patientdb.DB())
}

// createUser create new user
func createUser(user *models.User, ctx *fasthttp.RequestCtx) error {
	return (*user).Insert(ctx, patientdb.DB(), boil.Infer())
}
