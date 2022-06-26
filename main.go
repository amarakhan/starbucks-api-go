package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"errors"
    "strconv"
	"example/starbucks-api-go/model"
)

// TODO
// 1. Setup DB at Heroku
// 2. Add auth
// 3. Setup github vars for db conneciton
// 4. Move out user and menu funcs to seperate package?

// Temp data until database is setup
var menu = []model.Food{
	{ID: 1, Name: "Scone", Price: 4, DateAdded: "2022-01-01 01:01:01"},
	{ID: 2, Name: "Cookie", Price: 4, DateAdded: "2022-04-20 01:01:01"},
	{ID: 2, Name: "Decaf Coffee", Price: 3, DateAdded: "2022-07-11 01:01:01"},
}

var users = []model.User{
	{ID: 1, FirstName: "Spongebob", LastName: "Squarepants", Email:"spongebob@example.com", Staff:1, Address1: "1 Bikini Bottom", Address2: "", Zip: "", State: "", Country: "CA", AddDate: "2022-04-20 01:01:01", ModDate: "2022-06-20 01:01:01"},
	{ID: 1, FirstName: "Squidward", LastName: "tentacles", Email:"squidward@example.com", Staff:1, Address1: "2 Bikini Bottom", Address2: "", Zip: "", State: "", Country: "CA", AddDate: "2022-04-20 01:01:01", ModDate: "2022-06-20 01:01:01"},
}

// MENU FUNCS
func getMenu(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, menu)
}

func getMenuItemById(id int64) (*model.Food, error) {
	for i, b := range menu {
		if b.ID == id {
			return &menu[i], nil
		}
	}

	return nil, errors.New("Food not found in menu.")
}

func menuItemById(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.ParseInt(id, 0, 64)
	food, err := getMenuItemById(idInt)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Menu item not found."})
		return
	}

	c.IndentedJSON(http.StatusOK, food)
}

func addFood(c *gin.Context) {
	var newFood model.Food

	if err := c.BindJSON(&newFood); err != nil {
		return
	}

	menu = append(menu, newFood)
	c.IndentedJSON(http.StatusCreated, newFood)
}

// USER FUNCS
func getUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
}

func addUser(c *gin.Context) {
	var newUser model.User

	if err := c.BindJSON(&newUser); err != nil {
		return
	}

	users = append(users, newUser)
	c.IndentedJSON(http.StatusCreated, newUser)
}

func userById(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.ParseInt(id, 0, 64)
	user, err := getUserById(idInt)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "User not found."})
		return
	}

	c.IndentedJSON(http.StatusOK, user)
}

func getUserById(id int64) (*model.User, error) {
	for i, b := range users {
		if b.ID == id {
			return &users[i], nil
		}
	}

	return nil, errors.New("User not found.")
}

func main () {
	router := gin.Default()
	// menu routes
	router.GET("/menu", getMenu)
	router.GET("/menu/:id", menuItemById)
	router.POST("/menu", addFood)
	// user routes
	router.GET("/users", getUsers)
	router.POST("/user", addUser)
	router.GET("/user/:id", userById)

	router.Run("localhost:8080")
}