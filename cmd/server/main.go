package main

import (
	"idk/config"
	"idk/store/sql"
	"idk/transport/rest"
	"idk/user"
	"log"
)

func main() {
	c, err := config.InitConfig()
	if err != nil {
		log.Fatalf("failed to initialize config: %v", err)
	}

	// Store
	dbUrl := "postgres://" + c.DBUser + ":" + c.DBPassword + "@" + c.DBHost + ":" + c.DBPort + "/" + c.DBName + "?sslmode=disable"

	store, err := sql.New(dbUrl)
	if err != nil {
		log.Fatalf("failed to create store: %v", err)
	}
	defer store.Close()

	// User
	userRepo := sql.NewUserRepository(store)
	userService := user.NewService(userRepo)

	// REST
	router := rest.NewRouter(userService, c)
	router.Logger.Fatal(router.Start(":" + c.AppPort))
}
