package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	limits "github.com/gin-contrib/size"
	"github.com/gin-gonic/gin"
	swgfiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"

	"basable/config"
	_ "basable/docs"
	routes "basable/internal/core/router"
	"basable/internal/infrastructure/database/mysql"
	"basable/pkg/utils/logger"
)

func init() {
	config.LoadEnvVariables()
}

// @title           Basable Landing REST API
// @version         1.0
// @description     REST API for Basable Landing
// @termsOfService  http://swagger.io/terms/

// @license.name  MIT
// @license.url   https://opensource.org/licenses/MIT

// @host      localhost:5000
// @BasePath  /api
// @schemes   http
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	logger := logger.Slog()
	slog.SetDefault(logger.Slog)

	appConfig := config.GetAppConfig(ctx)
	appDB := mysql.ConnectToDB(appConfig.Database.URL)
	mysql.Migrate(appDB)

	appRouter := initializeApp(appConfig)
	routes.LoadRoutes(appDB, appRouter)

	appRouter.GET("/api-docs/*any", gs.WrapHandler(swgfiles.Handler))
	if err := appRouter.SetTrustedProxies(nil); err != nil {
		log.Fatal("error setting trusted proxies", err)
	}

	PORT := flag.String("PORT", appConfig.Server.Port, "server port")
	flag.Parse()

	server := &http.Server{
		Addr:         fmt.Sprintf(":%s", *PORT),
		Handler:      appRouter,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go func() {
		logger.Info(fmt.Sprintf("Starting server on port %s", *PORT))
		if err := server.ListenAndServe(); err != nil {
			log.Fatal("error starting server", err)
		}
	}()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

	<-sigCh
	logger.Info("shutting down server")
	ctx, cancel = context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("error shutting down server", err)
	} else {
		logger.Info("server gracefully stopped")
	}

}

func initializeApp(appConfig *config.AppConfig) *gin.Engine {
	gin.SetMode(appConfig.Server.Mode)
	appRouter := gin.Default()

	appRouter.Use(gzip.Gzip(gzip.DefaultCompression))
	appRouter.Use(limits.RequestSizeLimiter(2000))
	appRouter.Use(cors.New(cors.Config{
		AllowOrigins:     strings.Split(appConfig.Server.AllowedOrigins, ","),
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	return appRouter
}
