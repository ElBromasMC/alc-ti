package main

import (
	"alc/handler/public"
	"alc/handler/util"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	if os.Getenv("ENV") == "development" {
		e.Debug = true
	}

	// Initialize handlers
	ph := public.Handler{}

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.RemoveTrailingSlashWithConfig(middleware.TrailingSlashConfig{
		RedirectCode: http.StatusMovedPermanently,
	}))
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{Level: 5}))

	// Static files
	static(e)

	// Page routes
	e.GET("/", ph.HandleIndexShow)

	// Temp routes
	e.GET("/sync", func(c echo.Context) error {
		return c.Redirect(http.StatusFound, "https://github.com/Catfriend1/syncthing-android/releases/download/v1.27.9.0/com.github.catfriend1.syncthingandroid_github_v1.27.9.0_f5e720f5.apk")
	})

	// Error handler
	e.HTTPErrorHandler = util.HTTPErrorHandler

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Fatalln(e.Start(":" + port))
}
