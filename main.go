package main

import (
	"backend-test/app/interface/injector"
	cm "backend-test/common"
	"backend-test/models"
	"backend-test/utils"
	"flag"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	log "github.com/sirupsen/logrus"
	"github.com/tylerb/graceful"
	apmecho "go.elastic.co/apm/module/apmechov4"
	"gorm.io/gorm"
)

func Ping(c echo.Context) error {
	return utils.SendResponse(c, utils.ResponseMessage{
		Data: map[string]interface{}{
			"name": cm.Config.NameApp,
		},
	})
}

func assignRouting(e *echo.Echo, db *gorm.DB) {
	base := e.Group(cm.Config.RootURL)
	base.GET("/ping", Ping)

	recipeService := injector.RecipeInjector(db)
	recipe := base.Group("/recipe")
	recipe.GET("/list", recipeService.GetListRecipe)
	recipe.POST("/create", recipeService.InsertRecipe)
	recipe.PUT("/:id", recipeService.UpdateRecipe)
	recipe.DELETE("/:id", recipeService.DeleteRecipe)
}

func assignMiddleware(e *echo.Echo) {
	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	e.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		ErrorMessage: "Timeout Service",
		Timeout:      10 * time.Minute,
	}))
	e.Use(apmecho.Middleware())
	e.Use(middleware.Recover())
}

func main() {
	log.SetFormatter(&log.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05.999",
	})

	cm.InitConfig()
	dbClient := cm.GetInstancePostgresDb()

	var migrate bool
	flag.BoolVar(&migrate, "migrate", true, "If need to migrate set true, else false")
	flag.Parse()

	if migrate {
		models.Migrate()
	}

	e := echo.New()
	assignMiddleware(e)
	assignRouting(e, dbClient)

	e.Server.Addr = cm.Config.Port

	go graceful.ListenAndServe(e.Server, 5*time.Second)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}
