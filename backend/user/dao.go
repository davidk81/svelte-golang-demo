package user

import (
	"github.com/davidk81/svelte-golang-demo/backend/patientdb"
	"github.com/davidk81/svelte-golang-demo/backend/patientdb/models"
	"github.com/valyala/fasthttp"
	"github.com/volatiletech/sqlboiler/boil"
)

// handles database operations for user table

// GetUser checks username & password, and returns User data if successful
func GetUser(username string, ctx *fasthttp.RequestCtx) (*models.User, error) {
	return models.Users(models.UserWhere.Userid.EQ(username)).One(ctx, patientdb.DB())
}

// CreateUser create new user
func CreateUser(user *models.User, ctx *fasthttp.RequestCtx) error {
	return (*user).Insert(ctx, patientdb.DB(), boil.Infer())
}
