package routes

import (
	"encoding/json"

	"github.com/coyle/research-1/lang/storj-node-go/storage/boltdb"

	"github.com/kataras/iris"
	uuid "github.com/satori/go.uuid"
)

// Users contains the items needed to process requests to the user namespace
type Users struct {
	DB *boltdb.Client
}

// CreateUser instantiates a new user
func (u *Users) CreateUser(ctx iris.Context) {
	user := &boltdb.User{}

	if err := ctx.ReadJSON(user); err != nil {
		// TODO: Handle error
	}

	userBytes, err := json.Marshal(user)
	if err != nil {
		// TODO: Handle error
	}

	u.DB.CreateUser(uuid.NewV4().Bytes(), userBytes)
}
