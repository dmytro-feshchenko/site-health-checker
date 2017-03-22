package main

import (
	"fmt"
	"net/http"
	"os"
	"site-checker/controllers"
	"site-checker/db"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var connectionString = "host=localhost user=postgres dbname=upguard sslmode=disable password=postgres"

func main() {
	var err error
	// init connection to the db
	db.DBCon, err = gorm.Open("postgres", connectionString)
	if err != nil {
		panic(fmt.Sprintf("Database connection error: %s", err.Error()))
	}
	// close connection after stopping
	defer db.DBCon.Close()

	// apply auto migrations to all required models
	// db.DBCon.AutoMigrate(&models.User{}, &models.Lesson{}, &models.Course{})

	// run the web server
	RunServer()
}

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

	e.POST("/api/v1/login", controllers.Login)

	// Route => handler
	e.GET("/api/v1/status", func(c echo.Context) error {
		return c.JSON(http.StatusOK, 200)
	})

	// Restricted group
	r := e.Group("")
	// jwt token auth middleware
	r.Use(middleware.JWT([]byte(os.Getenv("SECRET"))))

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
