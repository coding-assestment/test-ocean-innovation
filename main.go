package main

import (
	"log"
	"net/http"
	"os"
	"test-ocean-innovation/handlers"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		err = godotenv.Load("local.env")
		if err != nil {
			log.Fatalf("Some error occured. Err: %s", err)
		}
	}

	PORT := os.Getenv("PORT")

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	g := e.Group("/api/v1")
	g.GET("/teams", handlers.GetAllTeams)
	g.POST("/teams", handlers.CreateTeam)
	g.GET("/teams/:id", handlers.GetTeam)
	g.PUT("/teams/:id", handlers.UpdateTeam)
	g.DELETE("/teams/:id", handlers.DeleteTeam)
	g.POST("/teams/:id/player", handlers.AddTeamPlayer)
	g.GET("/player/:id", handlers.GetPlayer)

	g.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.Logger.Fatal(e.Start(":" + PORT))
}
