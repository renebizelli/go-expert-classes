package main

import (
	"fmt"
	"net/http"
	"renebizelli/api/configs"
	"renebizelli/api/internal/entities"
	"renebizelli/api/internal/infra/database"
	"renebizelli/api/internal/infra/middlewares"
	"renebizelli/api/internal/infra/webserver/handlers"
	pkg_utils "renebizelli/api/pkg/utils"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	_ "renebizelli/api/docs"

	httpSwagger "github.com/swaggo/http-swagger"
)

// @title GO Expert API Example
// @version 1
// @description This is a sample server Petstore server.
// @host localhost:8080
// @BasePath /
// @schemes http
// @contact.email rene.bizelli@gmail.com
// @contact.name René Bizelli
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {

	config := configs.LoadConfig(".")

	db, err := initDb()

	pkg_utils.PanicIfError(err, "database error")

	productDb := database.NewProduct(db)
	userDb := database.NewUser(db)

	productHandler := handlers.NewProductHandler(productDb)
	userHandler := handlers.NewUserHandler(userDb, config.JWT.Token, config.JWT.ExpiresIn)

	r := chi.NewRouter()

	r.Use(middleware.Recoverer)     // não deixa a app cair, mesmo com panic
	r.Use(middlewares.CustomLogger) // loga as requisições

	r.Route("/products", func(c chi.Router) {

		c.Use(jwtauth.Verifier(config.JWT.Token))
		c.Use(jwtauth.Authenticator)

		c.Post("/", productHandler.CreateProduct)
		c.Get("/{page}/{limit}/{sort}", productHandler.GetAllProducts)
		c.Get("/{id}", productHandler.GetProduct)
		c.Put("/{id}", productHandler.UpdateProduct)
		c.Delete("/{id}", productHandler.DeleteProduct)
	})

	r.Post("/users", userHandler.CreateUser)
	r.Post("/login", userHandler.Login)

	r.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL(fmt.Sprintf("http://localhost:%s/docs/doc.json", config.WebServer.Port))))

	http.ListenAndServe(fmt.Sprintf(":%s", config.WebServer.Port), r)
}

func initDb() (*gorm.DB, error) {

	db, e := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if e != nil {
		return nil, e
	}

	db.AutoMigrate(&entities.User{}, &entities.Product{})

	return db, nil
}
