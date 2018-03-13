package main

import (
	"github.com/coyle/research-1/lang/storj-node-go/routes"
	"github.com/coyle/research-1/lang/storj-node-go/storage/boltdb"

	"github.com/kataras/iris"
)

// type Contact struct {
// 	Id int64 `json:"id"`
// }

func main() {

	bdb, err := boltdb.New()
	if err != nil {

	}

	defer bdb.DB.Close()

	users := routes.Users{DB: bdb}
	app := iris.Default()

	SetRoutes(app, users)

	// app.Get("/", func(ctx iris.Context) {
	// 	user := make(map[string]string)
	// 	user["username"] = "admin"
	// 	ctx.JSON(user)
	// })

	// app.Get("/users/{id:long min(1)}", func(ctx iris.Context) {
	// 	var user User

	// 	id, _ := ctx.Params().GetInt64("id")

	// 	user.Id = id
	// 	user.Username = "admin"

	// 	ctx.JSON(user)
	// })

	// app.Delete("/users/{id:long min(1)}", func(ctx iris.Context) {
	// 	ctx.JSON(user)
	// })

	// app.Get("/contacts", func(ctx iris.Context) {
	// 	users := make([]string, 0)
	// 	users = append(users, "user1")
	// 	ctx.JSON(users)
	// })

	// app.Get("/contacts/{id:long min(1)}", func(ctx iris.Context) {

	// 	details := make(map[string]string)
	// 	details["id"] = ctx.Params().Get("id")
	// 	details["username"] = "admin"

	// 	ctx.JSON(details)
	// })

	app.Run(iris.Addr(":8080"))
}

// SetRoutes defines all restful routes on the service
func SetRoutes(app *iris.Application, users routes.Users) {
	app.Post("/users", users.CreateUser)
	// app.Get("/users", db.ListUsers)
	// app.Delete("/users/:id", db.DeleteUser)
}
