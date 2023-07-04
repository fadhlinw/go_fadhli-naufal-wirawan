package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type User struct {
	Id int `json:"id" form:"id"`

	Name string `json:"name" form:"name"`

	Email string `json:"email" form:"email"`

	Password string `json:"password" form:"password"`
}

var users []User

// -------------------- controller --------------------

// get all users
func GetUsersController(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get all users",
		"users":    users,
	})
}

// get user by id

func GetUserController(c echo.Context) error {

	// your solution here
	id := c.Param("id")
	for _, user := range users {
		if strconv.Itoa(user.Id) == id {
			return c.JSON(http.StatusOK, user)
		}
	}
	return c.JSON(http.StatusNotFound, "User Tidak Ada")

}

// delete user by id

func DeleteUserController(c echo.Context) error {

	// your solution here
	id := c.Param("id")
	for i, user := range users {
		if strconv.Itoa(user.Id) == id {
			users = append(users[:i], users[i+1:]...)
			return c.JSON(http.StatusOK, users)
		}
	}
	return c.JSON(http.StatusNotFound, "User Tidak Ada")

}

// update user by id

func UpdateUserController(c echo.Context) error {

	// your solution here

	id := c.Param("id")
	for i, user := range users {
		if strconv.Itoa(user.Id) == id {
			newUser := User{}
			c.Bind(&newUser)
			users[i] = newUser

			return c.JSON(http.StatusOK, map[string]interface{}{
				"message": "User updated",
				"user":    newUser,
			})
		}
	}

	// if the user is not found, return a not found response
	return c.JSON(http.StatusNotFound, "User Tidak Ada")

}

// create new user
func CreateUserController(c echo.Context) error {
	// binding data
	user := User{}
	c.Bind(&user)

	if len(users) == 0 {
		user.Id = 1
	} else {
		newId := users[len(users)-1].Id + 1
		user.Id = newId
	}
	users = append(users, user)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success create user",
		"user":     user,
	})
}

// ---------------------------------------------------
func main() {
	e := echo.New()
	// routing with query parameter
	e.GET("/users", GetUsersController)
	e.GET("/users/:id", GetUserController)
	e.POST("/create-user", CreateUserController)
	e.DELETE("/users/:id", DeleteUserController)
	e.PUT("/users/:id", UpdateUserController)

	// start the server, and log if it fails

	e.Logger.Fatal(e.Start(":8000"))
}
