package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"

	userdao "github.com/wahyudibo/try-buildkite-go-project/internal/application/user/dao"
	"github.com/wahyudibo/try-buildkite-go-project/internal/application/user/features/getuserbyid"
	"github.com/wahyudibo/try-buildkite-go-project/internal/config"
	postgrespkg "github.com/wahyudibo/try-buildkite-go-project/internal/pkg/postgres"
)

func main() {
	ctx := context.Background()

	cfg, err := config.New()
	if err != nil {
		log.Fatalf("failed to initialize config: %v", err)
	}

	dbCfg, err := postgrespkg.NewConfig()
	if err != nil {
		log.Fatalf("failed to initialize database config: %v", err)
	}

	dbConn, err := postgrespkg.NewConnection(ctx, dbCfg)
	if err != nil {
		log.Fatalf("failed to initialize database: %v", err)
	}

	userDAO := userdao.New(dbConn)
	getUserByIDHandler := getuserbyid.New(userDAO)

	app := fiber.New()

	app.Get("/api/users/:userId", getUserByIDHandler.Handler)

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})

	log.Fatal(app.Listen(fmt.Sprintf(":%d", cfg.HTTPServerPort)))
}
