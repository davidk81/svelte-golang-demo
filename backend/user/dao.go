package user

import (
	"github.com/davidk81/svelte-golang-demo/backend/patientdb"
	"github.com/davidk81/svelte-golang-demo/backend/patientdb/models"
	"github.com/valyala/fasthttp"
)

// handles database operations for user table

// GetUser checks username & password, and returns User data if successful
func GetUser(username string, ctx *fasthttp.RequestCtx) (*models.User, error) {
	return models.Users(models.UserWhere.Userid.EQ(username)).One(ctx, patientdb.DB())
}
