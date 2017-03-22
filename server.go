package server

import (
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// RunServer - run the web server
func RunServer() {
	e := echo.New()

	// middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// configure CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.DELETE,
			echo.PATCH, echo.POST},
	}))

	e.Static("/static/*", "client/dist/static")
	e.Static("/public/*", "public/")

	// e.POST("/api/v1/login", controllers.Login)

	// Route => handler
	e.GET("/api/v1/status", func(c echo.Context) error {
		return c.JSON(http.StatusOK, 200)
	})

	// Restricted group
	r := e.Group("")
	// jwt token auth middleware
	r.Use(middleware.JWT([]byte("secret")))

	// Route => handler
	// r.POST("/api/v1/users", controllers.CreateUser)
	// e.GET("/api/v1/users", controllers.GetUsersList)
	// e.GET("/api/v1/users/:id", controllers.GetUser)
	//
	// r.POST("/api/v1/courses", controllers.CreateCourse)
	// e.GET("/api/v1/courses", controllers.GetCoursesList)
	// e.GET("/api/v1/courses/:id", controllers.GetUser)
	// r.PUT("/api/v1/courses/:id", controllers.UpdateCourse)
	// r.DELETE("/api/v1/courses/:id", controllers.DeleteCourse)

	// run the web server
	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
