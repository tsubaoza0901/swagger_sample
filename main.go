package main

import (
	"net/http"
	"strconv"

	echoSwagger "github.com/swaggo/echo-swagger"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"

	"ozawa-swagger-sample/docs"
)

// --------
// model↓
// --------

// User ...
type User struct {
	ID   uint   `json:"id" example:"1"`
	Name string `json:"name" example:"Tom"`
	Age  int    `json:"age" example:"20"`
}

// --------
// router↓
// --------

// InitRouting ...
func InitRouting(e *echo.Echo, u *User) {
	e.POST("user", u.CreateUser)
	e.PUT("user/:id", u.UpdateUser)
	e.DELETE("user/:id", u.DeleteUser)
	e.GET("user/:id", u.GetUser)
	e.GET("users", u.GetUsers)
}

// --------
// handler↓
// --------

// CreateUser ...
func (u *User) CreateUser(c echo.Context) error {
	user := User{}

	if err := c.Bind(&user); err != nil {
		log.Error(err)
		return err
	}

	user = User{
		ID:   1,
		Name: user.Name,
		Age:  user.Age,
	}

	return c.JSON(http.StatusOK, user)
}

// UpdateUser ...
func (u *User) UpdateUser(c echo.Context) error {
	user := User{}

	if err := c.Bind(&user); err != nil {
		log.Error(err)
		return err
	}

	return c.JSON(http.StatusOK, "Updated")
}

// DeleteUser ...
func (u *User) DeleteUser(c echo.Context) error {
	if err := c.Bind(u); err != nil {
		log.Error(err)
		return err
	}
	return c.JSON(http.StatusOK, "Deleted")
}

// GetUser ...
// @Summary Show a user
// @Description get user by ID
// @ID get-user-by-int
// @Accept application/json
// @Produce application/json
// @Param id path int true "User ID"
// @Success 200 {object} User
// @Failure 400,404 "Error"
// @Failure 500 "Error"
// @Failure default "Error"
// @Resource /user
// @Router /user/{id} [get]
func (u *User) GetUser(c echo.Context) error {
	user := User{}

	if err := c.Bind(&user); err != nil {
		log.Error(err)
		return err
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Error(err)
		return err
	}

	// Getメソッドのイメージ
	if id == 1 {
		user = User{
			ID:   1,
			Name: "Tom",
			Age:  29,
		}
	} else if id == 2 {
		user = User{
			ID:   2,
			Name: "Bob",
			Age:  35,
		}

	} else {
		return c.JSON(http.StatusOK, "Not Found")
	}

	return c.JSON(http.StatusOK, user)
}

// GetUsers ...
func (u *User) GetUsers(c echo.Context) error {
	users := []*User{}

	name := c.QueryParam("name")

	// Get Allメソッドのイメージ
	if name == "" {
		users = []*User{
			{
				ID:   1,
				Name: "Tom",
				Age:  29,
			},
			{
				ID:   2,
				Name: "Bob",
				Age:  35,
			},
		}
	} else if name == "Tom" {
		users = []*User{
			{
				ID:   1,
				Name: "Tom",
				Age:  29,
			},
		}
	} else if name == "Bob" {
		users = []*User{
			{
				ID:   2,
				Name: "Bob",
				Age:  35,
			},
		}
	} else {
		return c.JSON(http.StatusOK, "Not Found")
	}

	return c.JSON(http.StatusOK, users)
}

// --------
// main.go↓
// --------

// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	e := echo.New()

	// swagger用のルーティング
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// programmaticallyにswagger infoを設定したい場合は以下のようにすることも可能
	docs.SwaggerInfo.Title = "Swagger Example API"
	docs.SwaggerInfo.Description = "This is a sample server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:9105"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	u := new(User)
	InitRouting(e, u)

	e.Start(":9105")
}
